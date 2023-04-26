// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// https://www.treasurysoftware.com/GF-CTX-Extract.pdf

package rule_stp_820

import "github.com/moov-io/x12/pkg/rules"

var TransactionSetRule = rules.TransactionRule{
	ST: rules.SegmentRule{
		Name:        "ST",
		Description: "TRANSACTION SET HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"820"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	Segments: rules.SegmentSetRule{
		0: rules.SegmentRule{
			Name:        "BPR",
			Description: "Beginning Segment for Payment Order/Remittance Advice",
			Mask:        rules.MASK_REQUIRED,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"C"}},
				"02": {Mask: rules.MASK_REQUIRED},
				"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"C"}},
				"04": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ACH"}},
				"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CTX"}},
				"06": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01"}},
				"07": {Mask: rules.MASK_OPTIONAL},
				"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"DA"}},
				"09": {Mask: rules.MASK_OPTIONAL},
				"10": {Mask: rules.MASK_OPTIONAL},
				"11": {Mask: rules.MASK_OPTIONAL},
				"12": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01"}},
				"13": {Mask: rules.MASK_OPTIONAL},
				"14": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"DA"}},
				"15": {Mask: rules.MASK_OPTIONAL},
				"16": {Mask: rules.MASK_REQUIRED},
				"17": {Mask: rules.MASK_NOTUSED},
				"18": {Mask: rules.MASK_NOTUSED},
				"19": {Mask: rules.MASK_NOTUSED},
				"20": {Mask: rules.MASK_NOTUSED},
				"21": {Mask: rules.MASK_NOTUSED},
			},
		},
		1: rules.SegmentRule{
			Name:        "TRN",
			Description: "Trace",
			Mask:        rules.MASK_REQUIRED,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1"}},
				"02": {Mask: rules.MASK_REQUIRED},
				"03": {Mask: rules.MASK_OPTIONAL},
				"04": {Mask: rules.MASK_OPTIONAL},
			},
		},
		2: rules.SegmentRule{
			Name:        "NM1",
			Description: "Originator Name Identification",
			Mask:        rules.MASK_REQUIRED,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PR"}},
				"02": {Mask: rules.MASK_OPTIONAL},
				"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"91"}},
				"04": {Mask: rules.MASK_OPTIONAL},
				"05": {Mask: rules.MASK_NOTUSED},
				"06": {Mask: rules.MASK_NOTUSED},
				"07": {Mask: rules.MASK_NOTUSED},
				"08": {Mask: rules.MASK_NOTUSED},
				"09": {Mask: rules.MASK_NOTUSED},
			},
		},
		3: rules.SegmentRule{
			Name:        "NM1",
			Description: "Receiver Name Identification",
			Mask:        rules.MASK_REQUIRED,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PE"}},
				"02": {Mask: rules.MASK_OPTIONAL},
				"03": {Mask: rules.MASK_NOTUSED},
				"04": {Mask: rules.MASK_NOTUSED},
				"05": {Mask: rules.MASK_NOTUSED},
				"06": {Mask: rules.MASK_NOTUSED},
				"07": {Mask: rules.MASK_NOTUSED},
				"08": {Mask: rules.MASK_NOTUSED},
				"09": {Mask: rules.MASK_NOTUSED},
			},
		},
	},
	SE: rules.SegmentRule{
		Name:        "SE",
		Description: "TRANSACTION SET TRAILER",
		Mask:        rules.MASK_REQUIRED,
		Elements:    rules.ElementSetRule{},
	},
	Loops: map[int]rules.LoopRule{
		0: {
			Segments: rules.SegmentSetRule{
				0: rules.SegmentRule{
					Name:        "ENT",
					Description: "Entity",
					Mask:        rules.MASK_OPTIONAL,
					Elements: rules.ElementSetRule{
						"01": {Mask: rules.MASK_OPTIONAL},
						"02": {Mask: rules.MASK_NOTUSED},
						"03": {Mask: rules.MASK_NOTUSED},
						"04": {Mask: rules.MASK_NOTUSED},
					},
				},
				1: rules.SegmentRule{
					Name:        "RMR",
					Description: "Remittance Advice Accounts Receivable Open Item Reference",
					Mask:        rules.MASK_REQUIRED,
					Elements: rules.ElementSetRule{
						"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"IV", "PO", "R7"}},
						"02": {Mask: rules.MASK_OPTIONAL},
						"03": {Mask: rules.MASK_OPTIONAL},
						"04": {Mask: rules.MASK_REQUIRED},
						"05": {Mask: rules.MASK_OPTIONAL},
						"06": {Mask: rules.MASK_OPTIONAL},
					},
				},
				2: rules.SegmentRule{
					Name:        "REF",
					Description: "Reference Identification",
					Mask:        rules.MASK_OPTIONAL,
					RepeatCount: rules.GREATER_THAN_ONE,
					Elements: rules.ElementSetRule{
						"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"BM", "PO", "R7", "VV"}},
						"02": {Mask: rules.MASK_REQUIRED},
					},
				},
				3: rules.SegmentRule{
					Name:        "DTM",
					Description: "INDIVIDUAL COVERAGE PERIOD",
					Mask:        rules.MASK_OPTIONAL,
					Elements: rules.ElementSetRule{
						"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"003", "004", "092"}},
						"02": {Mask: rules.MASK_OPTIONAL},
						"03": {Mask: rules.MASK_NOTUSED},
						"04": {Mask: rules.MASK_NOTUSED},
						"05": {Mask: rules.MASK_NOTUSED},
						"06": {Mask: rules.MASK_NOTUSED},
					},
				},
				4: rules.SegmentRule{
					Name:        "ADX",
					Description: "Adjustment",
					Mask:        rules.MASK_OPTIONAL,
					Elements: rules.ElementSetRule{
						"01": {Mask: rules.MASK_REQUIRED},
						"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "03", "04", "05", "06", "07", "11", "12", "59", "75", "81", "CM"}},
						"03": {Mask: rules.MASK_OPTIONAL},
						"04": {Mask: rules.MASK_OPTIONAL},
					},
				},
			},
			Mask:        rules.MASK_OPTIONAL,
			Name:        "Detail",
			RepeatCount: rules.GREATER_THAN_ONE,
		},
	},
}

var GroupRule = rules.GroupRule{
	GS: rules.SegmentRule{
		Name:        "GS",
		Description: "FUNCTIONAL GROUP HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"RA"}},
			"07": {AcceptValues: []string{"X"}},
			"08": {AcceptValues: []string{"004010STP820"}},
		},
	},
	GE: rules.SegmentRule{
		Name:        "GE",
		Description: "FUNCTIONAL GROUP TRAILER",
		Mask:        rules.MASK_REQUIRED,
		Elements:    rules.ElementSetRule{},
	},
	Trans: TransactionSetRule,
}

var InterchangeRule = rules.InterchangeRule{
	Name: "EPN STP 820",
	ISA: rules.SegmentRule{
		Name:        "ISA",
		Description: "INTERCHANGE CONTROL HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"00"}},
			"03": {AcceptValues: []string{"00"}},
			"05": {AcceptValues: []string{"01", "14", "17", "30"}},
			"07": {AcceptValues: []string{"01", "14", "17", "30"}},
			"11": {AcceptValues: []string{"U"}},
			"12": {AcceptValues: []string{"00401"}},
			"14": {AcceptValues: []string{"0"}},
			"15": {AcceptValues: []string{"P"}},
			"16": {AcceptValues: []string{"~"}},
		},
	},
	IEA: rules.SegmentRule{
		Name:        "IEA",
		Description: "INTERCHANGE CONTROL TRAILER",
		Mask:        rules.MASK_REQUIRED,
		Elements:    rules.ElementSetRule{},
	},
	Group: GroupRule,
}
