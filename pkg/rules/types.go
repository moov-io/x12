// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package rules

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

const (
	MASK_REQUIRED    = "REQUIRED"
	MASK_OPTIONAL    = "OPTIONAL"
	MASK_NOTUSED     = "NOTUSED"
	MASK_NONE        = ""
	GREATER_THAN_ONE = 0xFFFF
)

func GetMask(mask, defaultMask string) string {
	if mask != MASK_NONE {
		return mask
	}

	if defaultMask != MASK_NONE {
		return defaultMask
	}

	return MASK_OPTIONAL
}

func IsMaskRequired(mask string) bool {
	if mask == MASK_NONE {
		// default mask is required
		return true
	}
	return mask == MASK_REQUIRED
}

type ruleInfo struct {
	Mask        string
	RepeatCount int
	Name        string
	Description string
	Level       int
}

// Rule for interchange
type InterchangeRule struct {
	Name  string
	ISA   SegmentRule
	IEA   SegmentRule
	Group GroupRule
}

func (r InterchangeRule) dumpRuleInfo(level int, isRequiredOnly bool) []ruleInfo {
	var infos []ruleInfo

	if r.ISA.isRequired(isRequiredOnly) {
		infos = append(infos, r.ISA.dumpRuleInfo(level))
	}
	infos = append(infos, r.Group.dumpRuleInfo(level+1, isRequiredOnly)...)
	if r.IEA.isRequired(isRequiredOnly) {
		infos = append(infos, r.IEA.dumpRuleInfo(level))
	}

	return infos
}

func (r InterchangeRule) Print(w io.Writer, isRequiredOnly bool) {
	w.Write([]byte("\n  DUMP RULE " + r.Name + "\n\n"))

	padChar := byte(' ')
	padding := 1
	tw := tabwriter.NewWriter(w, 0, 0, padding, padChar, tabwriter.Debug)

	// header line
	{
		printStr := "Segments & Rules Structure:\t"

		columns := []string{"Usage", "Repeat Count", "Description"}
		for i := 0; i < len(columns); i++ {
			printStr = printStr + " " + columns[i] + "\t"
		}
		fmt.Fprintln(tw, printStr)
	}

	selfDumps := r.dumpRuleInfo(0, isRequiredOnly)
	for _, d := range selfDumps {
		printStr := d.Name + "\t"
		printStr = printStr + d.Mask + "\t"
		if d.RepeatCount == GREATER_THAN_ONE {
			printStr = printStr + ">1\t"
		} else {
			printStr = printStr + fmt.Sprintf("%v", d.RepeatCount) + "\t"
		}
		printStr = printStr + d.Description + "\t"

		if d.Level > 0 {
			printStr = strings.Repeat(string(padChar), d.Level) + printStr
		}

		fmt.Fprintln(tw, printStr)
	}

	tw.Flush()
}

// Rule for group
type GroupRule struct {
	GS    SegmentRule
	GE    SegmentRule
	Trans TransactionRule
}

func (g GroupRule) dumpRuleInfo(level int, isRequiredOnly bool) []ruleInfo {
	var infos []ruleInfo

	if g.GS.isRequired(isRequiredOnly) {
		infos = append(infos, g.GS.dumpRuleInfo(level))
	}
	infos = append(infos, g.Trans.dumpRuleInfo(level+1, isRequiredOnly)...)
	if g.GE.isRequired(isRequiredOnly) {
		infos = append(infos, g.GE.dumpRuleInfo(level))
	}

	return infos
}

// Rule for transaction set
type TransactionRule struct {
	ST        SegmentRule
	SE        SegmentRule
	Composite LoopRule
}

func (t TransactionRule) dumpRuleInfo(level int, isRequiredOnly bool) []ruleInfo {
	var infos []ruleInfo

	if t.ST.isRequired(isRequiredOnly) {
		infos = append(infos, t.ST.dumpRuleInfo(level))
	}

	infos = append(infos, t.Composite.dumpRuleInfo(level, isRequiredOnly)...)

	if t.SE.isRequired(isRequiredOnly) {
		infos = append(infos, t.SE.dumpRuleInfo(level))
	}

	return infos
}

// Rule for loop
type LoopSetRule map[int]LoopRule

type LoopRule struct {
	Segments    SegmentSetRule
	Mask        string
	RepeatCount int
	Name        string
	Composite   LoopSetRule
}

func (l LoopRule) Repeat() int {
	if l.RepeatCount > 1 {
		return l.RepeatCount
	}

	return 1
}

func (l LoopRule) MinRepeat() int {
	if l.Mask != MASK_REQUIRED {
		return 0
	}

	if l.RepeatCount == GREATER_THAN_ONE {
		return 1
	}

	return 1
}

func (l LoopRule) isRequired(isRequiredOnly bool) bool {
	return !(isRequiredOnly && GetMask(l.Mask, MASK_REQUIRED) != MASK_REQUIRED)
}

func (l LoopRule) dumpRuleInfo(level int, isRequiredOnly bool) []ruleInfo {
	var infos []ruleInfo

	if l.isRequired(isRequiredOnly) {
		info := ruleInfo{
			Name:        l.Name,
			Mask:        GetMask(l.Mask, MASK_REQUIRED),
			RepeatCount: l.Repeat(),
			Level:       level,
		}

		infos = append(infos, info)
	} else {
		return nil
	}

	for index := 0; index < len(l.Segments); index++ {
		s := l.Segments[index]
		if s.isRequired(isRequiredOnly) {
			infos = append(infos, s.dumpRuleInfo(level+1))
		}
	}

	for index := 0; index < len(l.Composite); index++ {
		c := l.Composite[index]
		infos = append(infos, c.dumpRuleInfo(level+1, isRequiredOnly)...)
	}

	return infos
}

type SegmentSetRule map[int]SegmentRule

type SegmentRule struct {
	Elements    ElementSetRule
	Mask        string
	RepeatCount int
	Name        string
	Description string
}

func (s SegmentRule) isRequired(isRequiredOnly bool) bool {
	return !(isRequiredOnly && GetMask(s.Mask, MASK_REQUIRED) != MASK_REQUIRED)
}

func (s SegmentRule) dumpRuleInfo(level int) ruleInfo {
	return ruleInfo{
		Name:        s.Name,
		Mask:        GetMask(s.Mask, MASK_REQUIRED),
		RepeatCount: s.Repeat(),
		Description: s.Description,
		Level:       level,
	}
}

func (s SegmentRule) Repeat() int {
	if s.RepeatCount > 1 {
		return s.RepeatCount
	}

	return 1
}

func (s SegmentRule) MinRepeat() int {
	if s.Mask != MASK_REQUIRED {
		return 0
	}

	if s.RepeatCount == GREATER_THAN_ONE {
		return 1
	}

	return 1
}

type ElementSetRule map[string]ElementRule

type ElementRule struct {
	Mask         string
	Name         string
	AcceptRegex  string
	AcceptValues []string
	Composite    ElementSetRule
}

func (s SegmentSetRule) Get(index int) *SegmentRule {
	segment, ok := s[index]
	if ok {
		return &segment
	}

	return nil
}

func (e ElementSetRule) Get(name string) ElementRule {
	element, ok := e[name]
	if ok {
		return element
	}

	return ElementRule{}
}

func (e ElementSetRule) GetMask(name, defaultMask string) string {
	element, ok := e[name]
	if !ok {
		element = ElementRule{}
	}

	return GetMask(element.Mask, defaultMask)
}
