// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package rules

const (
	MASK_REQUIRED = "REQUIRED"
	MASK_OPTIONAL = "OPTIONAL"
	MASK_NOTUSED  = "NOTUSED"
	MASK_NONE     = ""
)

func IsMaskRequired(mask string) bool {
	if mask == MASK_REQUIRED || mask == MASK_NONE {
		return true
	}
	return false
}

type InterchangeRule struct {
	Name  string
	ISA   SegmentRule
	IEA   SegmentRule
	Group GroupRule
}

type GroupRule struct {
	GS    SegmentRule
	GE    SegmentRule
	Trans TransactionRule
}

type TransactionRule struct {
	ST    SegmentRule
	BHT   SegmentRule
	SE    SegmentRule
	Loops LoopSetRule
}

type LoopSetRule map[int]LoopRule

type LoopRule struct {
	Segments    SegmentSetRule
	Mask        string
	RepeatCount int
	Name        string
	Composite   LoopSetRule
}

type SegmentSetRule map[int]SegmentRule

type SegmentRule struct {
	Elements    ElementSetRule
	Mask        string
	RepeatCount int
	Name        string
	Description string
}

type ElementSetRule map[string]ElementRule

type ElementRule struct {
	Mask         string
	Name         string
	AcceptValues []string
	Composite    ElementSetRule
}

func (s LoopRule) Repeat() int {
	if s.RepeatCount > 1 {
		return s.RepeatCount
	}

	return 1
}

func (s SegmentSetRule) Get(index int) *SegmentRule {

	segment, ok := s[index]
	if ok {
		return &segment
	}

	return nil
}

func (s SegmentRule) Repeat() int {
	if s.RepeatCount > 1 {
		return s.RepeatCount
	}

	return 1
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
