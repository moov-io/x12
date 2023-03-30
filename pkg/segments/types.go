// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"errors"
	"github.com/moov-io/x12/pkg/rules"
)

type Element struct {
	rule        *rules.ElementSetRule
	description string
}

func (e Element) GetRule() *rules.ElementSetRule {
	if e.rule != nil {
		return e.rule
	}

	newRule := make(rules.ElementSetRule)
	return &newRule
}

func (e *Element) SetRule(s *rules.ElementSetRule) {
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
	GetRule() *rules.ElementSetRule
	SetRule(s *rules.ElementSetRule)
	GetDescription() string
	SetDescription(string)
	GetFieldByIndex(index string) any
	SetFieldByIndex(index string, data any) error
	Validate(rule *rules.ElementSetRule) error
	Parse(data string, args ...string) (int, error)
	String(args ...string) string
}

var (
	_ SegmentInterface = (*AMT)(nil)
	_ SegmentInterface = (*BHT)(nil)
	_ SegmentInterface = (*CAS)(nil)
	_ SegmentInterface = (*CLM)(nil)
	_ SegmentInterface = (*CN1)(nil)
	_ SegmentInterface = (*CR1)(nil)
	_ SegmentInterface = (*CUR)(nil)
	_ SegmentInterface = (*DMG)(nil)
	_ SegmentInterface = (*DN1)(nil)
	_ SegmentInterface = (*DN2)(nil)
	_ SegmentInterface = (*DTP)(nil)
	_ SegmentInterface = (*GE)(nil)
	_ SegmentInterface = (*GS)(nil)
	_ SegmentInterface = (*HCP)(nil)
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
	_ SegmentInterface = (*SV3)(nil)
	_ SegmentInterface = (*SV5)(nil)
	_ SegmentInterface = (*SVD)(nil)
	_ SegmentInterface = (*TOO)(nil)
)

type constructorFunc func(rule *rules.ElementSetRule) SegmentInterface

var (
	segmentConstructor = map[string]constructorFunc{
		"AMT": func(rule *rules.ElementSetRule) SegmentInterface { return NewAMT(rule) },
		"BHT": func(rule *rules.ElementSetRule) SegmentInterface { return NewBHT(rule) },
		"CAS": func(rule *rules.ElementSetRule) SegmentInterface { return NewCAS(rule) },
		"CLM": func(rule *rules.ElementSetRule) SegmentInterface { return NewCLM(rule) },
		"CN1": func(rule *rules.ElementSetRule) SegmentInterface { return NewCN1(rule) },
		"CR1": func(rule *rules.ElementSetRule) SegmentInterface { return NewCR1(rule) },
		"CUR": func(rule *rules.ElementSetRule) SegmentInterface { return NewCUR(rule) },
		"DMG": func(rule *rules.ElementSetRule) SegmentInterface { return NewDMG(rule) },
		"DN1": func(rule *rules.ElementSetRule) SegmentInterface { return NewDN1(rule) },
		"DN2": func(rule *rules.ElementSetRule) SegmentInterface { return NewDN2(rule) },
		"DTP": func(rule *rules.ElementSetRule) SegmentInterface { return NewDTP(rule) },
		"GE":  func(rule *rules.ElementSetRule) SegmentInterface { return NewGE(rule) },
		"GS":  func(rule *rules.ElementSetRule) SegmentInterface { return NewGS(rule) },
		"HCP": func(rule *rules.ElementSetRule) SegmentInterface { return NewHCP(rule) },
		"HI":  func(rule *rules.ElementSetRule) SegmentInterface { return NewHI(rule) },
		"HL":  func(rule *rules.ElementSetRule) SegmentInterface { return NewHL(rule) },
		"IEA": func(rule *rules.ElementSetRule) SegmentInterface { return NewIEA(rule) },
		"ISA": func(rule *rules.ElementSetRule) SegmentInterface { return NewISA(rule) },
		"LX":  func(rule *rules.ElementSetRule) SegmentInterface { return NewLX(rule) },
		"MOA": func(rule *rules.ElementSetRule) SegmentInterface { return NewMOA(rule) },
		"N3":  func(rule *rules.ElementSetRule) SegmentInterface { return NewN3(rule) },
		"N4":  func(rule *rules.ElementSetRule) SegmentInterface { return NewN4(rule) },
		"NM1": func(rule *rules.ElementSetRule) SegmentInterface { return NewNM1(rule) },
		"NTE": func(rule *rules.ElementSetRule) SegmentInterface { return NewNTE(rule) },
		"OI":  func(rule *rules.ElementSetRule) SegmentInterface { return NewOI(rule) },
		"PAT": func(rule *rules.ElementSetRule) SegmentInterface { return NewPAT(rule) },
		"PER": func(rule *rules.ElementSetRule) SegmentInterface { return NewPER(rule) },
		"PRV": func(rule *rules.ElementSetRule) SegmentInterface { return NewPRV(rule) },
		"PWK": func(rule *rules.ElementSetRule) SegmentInterface { return NewPWK(rule) },
		"REF": func(rule *rules.ElementSetRule) SegmentInterface { return NewREF(rule) },
		"SBR": func(rule *rules.ElementSetRule) SegmentInterface { return NewSBR(rule) },
		"SE":  func(rule *rules.ElementSetRule) SegmentInterface { return NewSE(rule) },
		"ST":  func(rule *rules.ElementSetRule) SegmentInterface { return NewST(rule) },
		"SV1": func(rule *rules.ElementSetRule) SegmentInterface { return NewSV1(rule) },
		"SV3": func(rule *rules.ElementSetRule) SegmentInterface { return NewSV3(rule) },
		"SV5": func(rule *rules.ElementSetRule) SegmentInterface { return NewSV5(rule) },
		"SVD": func(rule *rules.ElementSetRule) SegmentInterface { return NewSVD(rule) },
		"TOO": func(rule *rules.ElementSetRule) SegmentInterface { return NewTOO(rule) },
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
