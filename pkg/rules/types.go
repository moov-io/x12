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
	MASK_REQUIRED = "REQUIRED"
	MASK_OPTIONAL = "OPTIONAL"
	MASK_NOTUSED  = "NOTUSED"
	MASK_NONE     = ""
	MAXCOUNT      = 300
)

func IsMaskRequired(mask string) bool {
	if mask == MASK_REQUIRED || mask == MASK_NONE {
		return true
	}
	return false
}

func getMask(mask1, mask2 string) string {
	if mask1 == MASK_NONE {
		return mask2
	}
	return mask1
}

type ruleInfo struct {
	Mask        string
	RepeatCount int
	Name        string
	Description string
	Level       int
}

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
		if d.RepeatCount == MAXCOUNT {
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

type TransactionRule struct {
	ST    SegmentRule
	SE    SegmentRule
	Loops LoopSetRule

	Segments SegmentSetRule
}

func (t TransactionRule) dumpRuleInfo(level int, isRequiredOnly bool) []ruleInfo {
	var infos []ruleInfo

	if t.ST.isRequired(isRequiredOnly) {
		infos = append(infos, t.ST.dumpRuleInfo(level))
	}

	for index := 0; index < len(t.Segments); index++ {
		s := t.Segments[index]
		infos = append(infos, s.dumpRuleInfo(level))
	}

	for index := 0; index < len(t.Loops); index++ {
		l := t.Loops[index]
		infos = append(infos, l.dumpRuleInfo(level+1, isRequiredOnly)...)
	}

	if t.SE.isRequired(isRequiredOnly) {
		infos = append(infos, t.SE.dumpRuleInfo(level))
	}

	return infos
}

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

func (l LoopRule) isRequired(isRequiredOnly bool) bool {
	return !(isRequiredOnly && getMask(l.Mask, MASK_REQUIRED) != MASK_REQUIRED)
}

func (l LoopRule) dumpRuleInfo(level int, isRequiredOnly bool) []ruleInfo {
	var infos []ruleInfo

	if l.isRequired(isRequiredOnly) {
		info := ruleInfo{
			Name:        l.Name,
			Mask:        getMask(l.Mask, MASK_REQUIRED),
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
		infos = append(infos, c.dumpRuleInfo(level+2, isRequiredOnly)...)
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
	return !(isRequiredOnly && getMask(s.Mask, MASK_REQUIRED) != MASK_REQUIRED)
}

func (s SegmentRule) dumpRuleInfo(level int) ruleInfo {
	return ruleInfo{
		Name:        s.Name,
		Mask:        getMask(s.Mask, MASK_REQUIRED),
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

type ElementSetRule map[string]ElementRule

type ElementRule struct {
	Mask         string
	Name         string
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

	mask := element.Mask
	if mask == MASK_NONE {
		mask = defaultMask
	}

	return mask
}
