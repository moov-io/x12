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

type InterChangeRule struct {
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
	Loops map[int]LoopRule
}

type LoopRule struct {
	Segments    Segments
	Mask        string
	RepeatCount int
	Name        string
	SubLoopRule map[int]LoopRule
}

func (s LoopRule) Repeat() int {
	if s.RepeatCount > 1 {
		return s.RepeatCount
	}

	return 1
}

type Segments map[int]SegmentRule

func (s Segments) Get(index int) *SegmentRule {

	segment, ok := s[index]
	if ok {
		return &segment
	}

	return nil
}

type SegmentRule struct {
	Elements    Elements
	Mask        string
	RepeatCount int
	Description string
	Name        string
}

func (s SegmentRule) Repeat() int {
	if s.RepeatCount > 1 {
		return s.RepeatCount
	}

	return 1
}

type Elements map[string]ElementRule

func (e Elements) Get(name string) ElementRule {

	element, ok := e[name]
	if ok {
		return element
	}

	return ElementRule{}
}

func (e Elements) GetMask(name, defaultMask string) string {

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

type ElementRule struct {
	AcceptValues   []string
	Mask           string
	HasSubElements bool
	SubRule        map[string]ElementRule
}
