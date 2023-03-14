// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"errors"
	"github.com/moov-io/x12/pkg/rules"
)

type Element struct {
	rule        *rules.Elements
	description string
}

func (e Element) GetRule() *rules.Elements {
	if e.rule != nil {
		return e.rule
	}

	newRule := make(rules.Elements)
	return &newRule
}

func (e *Element) SetRule(s *rules.Elements) {
	if s == nil || e == nil {
		return
	}

	e.rule = s
}

func (e Element) GetDescription() string {
	return e.description
}

func (e *Element) SetDescription(description string) {
	e.description = description
}

type SegmentInterface interface {
	Name() string
	GetRule() *rules.Elements
	SetRule(s *rules.Elements)
	GetDescription() string
	SetDescription(string)
	GetFieldByIndex(index string) any
	SetFieldByIndex(index string, data any) error
	Validate(rule *rules.Elements) error
	Parse(data string, args ...string) (int, error)
	String(args ...string) string
}

var (
	_ SegmentInterface = (*AMT)(nil)
	_ SegmentInterface = (*BHT)(nil)
	_ SegmentInterface = (*CAS)(nil)
	_ SegmentInterface = (*CLM)(nil)
	_ SegmentInterface = (*CR1)(nil)
	_ SegmentInterface = (*DMG)(nil)
	_ SegmentInterface = (*DTP)(nil)
	_ SegmentInterface = (*GE)(nil)
	_ SegmentInterface = (*GS)(nil)
	_ SegmentInterface = (*HI)(nil)
	_ SegmentInterface = (*HL)(nil)
	_ SegmentInterface = (*IEA)(nil)
	_ SegmentInterface = (*ISA)(nil)
	_ SegmentInterface = (*LX)(nil)
	_ SegmentInterface = (*MOA)(nil)
	_ SegmentInterface = (*N3)(nil)
	_ SegmentInterface = (*N4)(nil)
	_ SegmentInterface = (*NM1)(nil)
	_ SegmentInterface = (*NTE)(nil)
	_ SegmentInterface = (*OI)(nil)
	_ SegmentInterface = (*PAT)(nil)
	_ SegmentInterface = (*PER)(nil)
	_ SegmentInterface = (*PRV)(nil)
	_ SegmentInterface = (*PWK)(nil)
	_ SegmentInterface = (*REF)(nil)
	_ SegmentInterface = (*SBR)(nil)
	_ SegmentInterface = (*SE)(nil)
	_ SegmentInterface = (*ST)(nil)
	_ SegmentInterface = (*SV1)(nil)
	_ SegmentInterface = (*SV5)(nil)
	_ SegmentInterface = (*SVD)(nil)
)

type constructorFunc func(rule *rules.Elements) SegmentInterface

var (
	segmentConstructor = map[string]constructorFunc{
		"AMT": func(rule *rules.Elements) SegmentInterface { return NewAMT(rule) },
		"BHT": func(rule *rules.Elements) SegmentInterface { return NewBHT(rule) },
		"CAS": func(rule *rules.Elements) SegmentInterface { return NewCAS(rule) },
		"CLM": func(rule *rules.Elements) SegmentInterface { return NewCLM(rule) },
		"CR1": func(rule *rules.Elements) SegmentInterface { return NewCR1(rule) },
		"DMG": func(rule *rules.Elements) SegmentInterface { return NewDMG(rule) },
		"DTP": func(rule *rules.Elements) SegmentInterface { return NewDTP(rule) },
		"GE":  func(rule *rules.Elements) SegmentInterface { return NewGE(rule) },
		"GS":  func(rule *rules.Elements) SegmentInterface { return NewGS(rule) },
		"HI":  func(rule *rules.Elements) SegmentInterface { return NewHI(rule) },
		"HL":  func(rule *rules.Elements) SegmentInterface { return NewHL(rule) },
		"IEA": func(rule *rules.Elements) SegmentInterface { return NewIEA(rule) },
		"ISA": func(rule *rules.Elements) SegmentInterface { return NewISA(rule) },
		"LX":  func(rule *rules.Elements) SegmentInterface { return NewLX(rule) },
		"MOA": func(rule *rules.Elements) SegmentInterface { return NewMOA(rule) },
		"N3":  func(rule *rules.Elements) SegmentInterface { return NewN3(rule) },
		"N4":  func(rule *rules.Elements) SegmentInterface { return NewN4(rule) },
		"NM1": func(rule *rules.Elements) SegmentInterface { return NewNM1(rule) },
		"NTE": func(rule *rules.Elements) SegmentInterface { return NewNTE(rule) },
		"OI":  func(rule *rules.Elements) SegmentInterface { return NewOI(rule) },
		"PAT": func(rule *rules.Elements) SegmentInterface { return NewPAT(rule) },
		"PER": func(rule *rules.Elements) SegmentInterface { return NewPER(rule) },
		"PRV": func(rule *rules.Elements) SegmentInterface { return NewPRV(rule) },
		"PWK": func(rule *rules.Elements) SegmentInterface { return NewPWK(rule) },
		"REF": func(rule *rules.Elements) SegmentInterface { return NewREF(rule) },
		"SBR": func(rule *rules.Elements) SegmentInterface { return NewSBR(rule) },
		"SE":  func(rule *rules.Elements) SegmentInterface { return NewSE(rule) },
		"ST":  func(rule *rules.Elements) SegmentInterface { return NewST(rule) },
		"SV1": func(rule *rules.Elements) SegmentInterface { return NewSV1(rule) },
		"SV5": func(rule *rules.Elements) SegmentInterface { return NewSV5(rule) },
		"SVD": func(rule *rules.Elements) SegmentInterface { return NewSVD(rule) },
	}
)

func CreateSegment(name string, rule *rules.SegmentRule) (SegmentInterface, error) {

	constructor := segmentConstructor[name]
	if constructor == nil {
		return nil, errors.New("unsupported segment name")
	}

	if rule == nil {
		rule = &rules.SegmentRule{}
	}

	newSegment := constructor(&rule.Elements)
	newSegment.SetDescription(rule.Description)

	return newSegment, nil
}
