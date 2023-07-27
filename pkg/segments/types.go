// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"errors"
	"fmt"
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
	"strings"
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

func (e *Element) CompositeString(buf, mask, separator, terminate string, value any) string {
	if buf == "" {
		if mask == rules.MASK_NOTUSED {
			return buf
		}
		if mask == rules.MASK_OPTIONAL && (value == nil || fmt.Sprintf("%v", value) == "") {
			return buf
		}
	}

	if buf == "" {
		buf = fmt.Sprintf("%v%s", value, terminate)
	} else {
		buf = fmt.Sprintf("%v%s", value, separator) + buf
	}

	return buf
}

func (e *Element) TerminateString(buf, name string, args ...string) string {
	if buf == "" {
		buf = fmt.Sprintf("%s%s", name, util.GetSegmentTerminator(args...))
	} else {
		buf = fmt.Sprintf("%s%s", name, util.DataElementSeparator) + buf
	}

	return buf
}

func (e *Element) VerifyCode(data, name string, args ...string) (int, string, error) {
	length := util.GetRecordSize(data, args...)
	codeLen := len(name)
	read := codeLen + 1

	if length < codeLen {
		return 0, "", fmt.Errorf("%s segment has not enough input data", name)
	} else if strings.ToUpper(name) != strings.ToUpper(data[:codeLen]) {
		return 0, "", fmt.Errorf("%s segment contains invalid code", name)
	}

	return read, data[:length], nil
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
	// Args
	// First arg: SegmentTerminator
	// Second arg: SubElementSeparator
	Parse(data string, args ...string) (int, error)
	String(args ...string) string
}

var (
	_ SegmentInterface = (*AMT)(nil)
	_ SegmentInterface = (*BHT)(nil)
	_ SegmentInterface = (*BIG)(nil)
	_ SegmentInterface = (*BPR)(nil)
	_ SegmentInterface = (*CAS)(nil)
	_ SegmentInterface = (*CLM)(nil)
	_ SegmentInterface = (*CLP)(nil)
	_ SegmentInterface = (*CN1)(nil)
	_ SegmentInterface = (*CR1)(nil)
	_ SegmentInterface = (*CR6)(nil)
	_ SegmentInterface = (*CRC)(nil)
	_ SegmentInterface = (*CTP)(nil)
	_ SegmentInterface = (*CTT)(nil)
	_ SegmentInterface = (*CUR)(nil)
	_ SegmentInterface = (*DMG)(nil)
	_ SegmentInterface = (*DN1)(nil)
	_ SegmentInterface = (*DN2)(nil)
	_ SegmentInterface = (*DTM)(nil)
	_ SegmentInterface = (*DTP)(nil)
	_ SegmentInterface = (*ENT)(nil)
	_ SegmentInterface = (*FA1)(nil)
	_ SegmentInterface = (*FA2)(nil)
	_ SegmentInterface = (*GE)(nil)
	_ SegmentInterface = (*GS)(nil)
	_ SegmentInterface = (*HCP)(nil)
	_ SegmentInterface = (*HI)(nil)
	_ SegmentInterface = (*HL)(nil)
	_ SegmentInterface = (*IEA)(nil)
	_ SegmentInterface = (*ISA)(nil)
	_ SegmentInterface = (*IT1)(nil)
	_ SegmentInterface = (*ITD)(nil)
	_ SegmentInterface = (*LQ)(nil)
	_ SegmentInterface = (*LX)(nil)
	_ SegmentInterface = (*MIA)(nil)
	_ SegmentInterface = (*MOA)(nil)
	_ SegmentInterface = (*MSG)(nil)
	_ SegmentInterface = (*N1)(nil)
	_ SegmentInterface = (*N2)(nil)
	_ SegmentInterface = (*N3)(nil)
	_ SegmentInterface = (*N4)(nil)
	_ SegmentInterface = (*N9)(nil)
	_ SegmentInterface = (*NM1)(nil)
	_ SegmentInterface = (*NTE)(nil)
	_ SegmentInterface = (*OI)(nil)
	_ SegmentInterface = (*PAT)(nil)
	_ SegmentInterface = (*PER)(nil)
	_ SegmentInterface = (*PID)(nil)
	_ SegmentInterface = (*PLB)(nil)
	_ SegmentInterface = (*PRV)(nil)
	_ SegmentInterface = (*PWK)(nil)
	_ SegmentInterface = (*QTY)(nil)
	_ SegmentInterface = (*RDM)(nil)
	_ SegmentInterface = (*REF)(nil)
	_ SegmentInterface = (*RMR)(nil)
	_ SegmentInterface = (*SAC)(nil)
	_ SegmentInterface = (*SBR)(nil)
	_ SegmentInterface = (*SE)(nil)
	_ SegmentInterface = (*ST)(nil)
	_ SegmentInterface = (*SV1)(nil)
	_ SegmentInterface = (*SV2)(nil)
	_ SegmentInterface = (*SV3)(nil)
	_ SegmentInterface = (*SV5)(nil)
	_ SegmentInterface = (*SVC)(nil)
	_ SegmentInterface = (*SVD)(nil)
	_ SegmentInterface = (*TDS)(nil)
	_ SegmentInterface = (*TOO)(nil)
	_ SegmentInterface = (*TRN)(nil)
	_ SegmentInterface = (*TS2)(nil)
	_ SegmentInterface = (*TS3)(nil)
	_ SegmentInterface = (*TXI)(nil)
)

type constructorFunc func(rule *rules.ElementSetRule) SegmentInterface

var (
	segmentConstructor = map[string]constructorFunc{
		"AMT": func(rule *rules.ElementSetRule) SegmentInterface { return NewAMT(rule) },
		"BHT": func(rule *rules.ElementSetRule) SegmentInterface { return NewBHT(rule) },
		"BIG": func(rule *rules.ElementSetRule) SegmentInterface { return NewBIG(rule) },
		"BPR": func(rule *rules.ElementSetRule) SegmentInterface { return NewBPR(rule) },
		"CAS": func(rule *rules.ElementSetRule) SegmentInterface { return NewCAS(rule) },
		"CLM": func(rule *rules.ElementSetRule) SegmentInterface { return NewCLM(rule) },
		"CLP": func(rule *rules.ElementSetRule) SegmentInterface { return NewCLP(rule) },
		"CN1": func(rule *rules.ElementSetRule) SegmentInterface { return NewCN1(rule) },
		"CR1": func(rule *rules.ElementSetRule) SegmentInterface { return NewCR1(rule) },
		"CR6": func(rule *rules.ElementSetRule) SegmentInterface { return NewCR6(rule) },
		"CRC": func(rule *rules.ElementSetRule) SegmentInterface { return NewCRC(rule) },
		"CTP": func(rule *rules.ElementSetRule) SegmentInterface { return NewCTP(rule) },
		"CTT": func(rule *rules.ElementSetRule) SegmentInterface { return NewCTT(rule) },
		"CUR": func(rule *rules.ElementSetRule) SegmentInterface { return NewCUR(rule) },
		"DMG": func(rule *rules.ElementSetRule) SegmentInterface { return NewDMG(rule) },
		"DN1": func(rule *rules.ElementSetRule) SegmentInterface { return NewDN1(rule) },
		"DN2": func(rule *rules.ElementSetRule) SegmentInterface { return NewDN2(rule) },
		"DTM": func(rule *rules.ElementSetRule) SegmentInterface { return NewDTM(rule) },
		"DTP": func(rule *rules.ElementSetRule) SegmentInterface { return NewDTP(rule) },
		"ENT": func(rule *rules.ElementSetRule) SegmentInterface { return NewENT(rule) },
		"FA1": func(rule *rules.ElementSetRule) SegmentInterface { return NewFA1(rule) },
		"FA2": func(rule *rules.ElementSetRule) SegmentInterface { return NewFA2(rule) },
		"GE":  func(rule *rules.ElementSetRule) SegmentInterface { return NewGE(rule) },
		"GS":  func(rule *rules.ElementSetRule) SegmentInterface { return NewGS(rule) },
		"HCP": func(rule *rules.ElementSetRule) SegmentInterface { return NewHCP(rule) },
		"HI":  func(rule *rules.ElementSetRule) SegmentInterface { return NewHI(rule) },
		"HL":  func(rule *rules.ElementSetRule) SegmentInterface { return NewHL(rule) },
		"IEA": func(rule *rules.ElementSetRule) SegmentInterface { return NewIEA(rule) },
		"ISA": func(rule *rules.ElementSetRule) SegmentInterface { return NewISA(rule) },
		"IT1": func(rule *rules.ElementSetRule) SegmentInterface { return NewIT1(rule) },
		"ITD": func(rule *rules.ElementSetRule) SegmentInterface { return NewITD(rule) },
		"LQ":  func(rule *rules.ElementSetRule) SegmentInterface { return NewLQ(rule) },
		"LX":  func(rule *rules.ElementSetRule) SegmentInterface { return NewLX(rule) },
		"MIA": func(rule *rules.ElementSetRule) SegmentInterface { return NewMIA(rule) },
		"MOA": func(rule *rules.ElementSetRule) SegmentInterface { return NewMOA(rule) },
		"MSG": func(rule *rules.ElementSetRule) SegmentInterface { return NewMSG(rule) },
		"N1":  func(rule *rules.ElementSetRule) SegmentInterface { return NewN1(rule) },
		"N2":  func(rule *rules.ElementSetRule) SegmentInterface { return NewN2(rule) },
		"N3":  func(rule *rules.ElementSetRule) SegmentInterface { return NewN3(rule) },
		"N4":  func(rule *rules.ElementSetRule) SegmentInterface { return NewN4(rule) },
		"N9":  func(rule *rules.ElementSetRule) SegmentInterface { return NewN9(rule) },
		"NM1": func(rule *rules.ElementSetRule) SegmentInterface { return NewNM1(rule) },
		"NTE": func(rule *rules.ElementSetRule) SegmentInterface { return NewNTE(rule) },
		"OI":  func(rule *rules.ElementSetRule) SegmentInterface { return NewOI(rule) },
		"PAT": func(rule *rules.ElementSetRule) SegmentInterface { return NewPAT(rule) },
		"PER": func(rule *rules.ElementSetRule) SegmentInterface { return NewPER(rule) },
		"PID": func(rule *rules.ElementSetRule) SegmentInterface { return NewPID(rule) },
		"PLB": func(rule *rules.ElementSetRule) SegmentInterface { return NewPLB(rule) },
		"PRV": func(rule *rules.ElementSetRule) SegmentInterface { return NewPRV(rule) },
		"PWK": func(rule *rules.ElementSetRule) SegmentInterface { return NewPWK(rule) },
		"QTY": func(rule *rules.ElementSetRule) SegmentInterface { return NewQTY(rule) },
		"RDM": func(rule *rules.ElementSetRule) SegmentInterface { return NewRDM(rule) },
		"REF": func(rule *rules.ElementSetRule) SegmentInterface { return NewREF(rule) },
		"RMR": func(rule *rules.ElementSetRule) SegmentInterface { return NewRMR(rule) },
		"SAC": func(rule *rules.ElementSetRule) SegmentInterface { return NewSAC(rule) },
		"SBR": func(rule *rules.ElementSetRule) SegmentInterface { return NewSBR(rule) },
		"SE":  func(rule *rules.ElementSetRule) SegmentInterface { return NewSE(rule) },
		"ST":  func(rule *rules.ElementSetRule) SegmentInterface { return NewST(rule) },
		"SV1": func(rule *rules.ElementSetRule) SegmentInterface { return NewSV1(rule) },
		"SV2": func(rule *rules.ElementSetRule) SegmentInterface { return NewSV2(rule) },
		"SV3": func(rule *rules.ElementSetRule) SegmentInterface { return NewSV3(rule) },
		"SV5": func(rule *rules.ElementSetRule) SegmentInterface { return NewSV5(rule) },
		"SVC": func(rule *rules.ElementSetRule) SegmentInterface { return NewSVC(rule) },
		"SVD": func(rule *rules.ElementSetRule) SegmentInterface { return NewSVD(rule) },
		"TDS": func(rule *rules.ElementSetRule) SegmentInterface { return NewTDS(rule) },
		"TOO": func(rule *rules.ElementSetRule) SegmentInterface { return NewTOO(rule) },
		"TRN": func(rule *rules.ElementSetRule) SegmentInterface { return NewTRN(rule) },
		"TS2": func(rule *rules.ElementSetRule) SegmentInterface { return NewTS2(rule) },
		"TS3": func(rule *rules.ElementSetRule) SegmentInterface { return NewTS3(rule) },
		"TXI": func(rule *rules.ElementSetRule) SegmentInterface { return NewTXI(rule) },
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
