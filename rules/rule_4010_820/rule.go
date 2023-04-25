// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// 004010X061A1

package rule_4010_820

import "github.com/moov-io/x12/pkg/rules"

var L1000ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PREMIUM RECEIVER'S NAME",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PE"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"1", "9", "EQ", "FI", "XV"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_NOTUSED},
		},
	},
	1: rules.SegmentRule{
		Name:        "NM2",
		Description: "PREMIUM RECEIVER ADDITIONAL NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "NM3",
		Description: "PREMIUM RECEIVER'S ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "NM4",
		Description: "PREMIUM RECEIVER'S CITY",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L1000BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PREMIUM PAYER'S NAME",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PR"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"1", "9", "24", "75", "EQ", "FI", "PI"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_NOTUSED},
		},
	},
	1: rules.SegmentRule{
		Name:        "NM2",
		Description: "PREMIUM PAYER ADDITIONAL NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "NM3",
		Description: "PREMIUM PAYER'S ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "NM4",
		Description: "PREMIUM PAYER'S CITY",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "PER",
		Description: "PREMIUM PAYER'S ADMINISTRATIVE CONTACT",
		RepeatCount: rules.MAXCOUNT,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"IC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EM", "FX", "TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EM", "EX", "FX", "TE"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EM", "EX", "FX", "TE"}},
			"08": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2000ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "ENT",
		Description: "ORGANIZATION SUMMARY REMITTANCE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"2L"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"1", "9", "FI"}},
		},
	},
}

var L2300ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "RMR",
		Description: "ORGANIZATION SUMMARY REMITTANCE DETAIL",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"11", "1L", "CT", "IK"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"PA", "PI", "PO", "PP"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2310ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "IT1",
		Description: "SUMMARY LINE ITEM",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
		},
	},
}

var L2315ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "SLN",
		Description: "MEMBER COUNT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"O"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED},
		},
	},
}

var L2320ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "ADX",
		Description: "ORGANIZATION SUMMARY REMITTANCE LEVEL ADJUSTMENT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"20", "52", "53", "AA", "H1", "H6", "IA", "J3"}},
		},
	},
}

var L2320BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "ADX",
		Description: "ORGANIZATION SUMMARY REMITTANCE LEVEL ADJUSTMENT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"20", "52", "53", "AA", "AX", "H1", "H6", "IA", "J3"}},
		},
	},
}

var L2000BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "ENT",
		Description: "INDIVIDUAL REMITTANCE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"2J"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"34", "EI", "ZZ"}},
		},
	},
}

var L2100BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "INDIVIDUAL NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"EY", "QE"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"34", "EI", "N"}},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2300BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "RMR",
		Description: "ORGANIZATION SUMMARY REMITTANCE DETAIL",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"11", "9J", "AZ", "B7", "CT", "ID", "IG", "IK", "KW"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"PI", "PP"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "DTM",
		Description: "INDIVIDUAL COVERAGE PERIOD",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"582"}},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"RD8"}},
			"06": {Mask: rules.MASK_REQUIRED},
		},
	},
}

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
			Description: "FINANCIAL INFORMATION",
			Mask:        rules.MASK_REQUIRED,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"C", "D", "I", "P", "U", "X"}},
				"02": {Mask: rules.MASK_REQUIRED},
				"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"C", "D"}},
				"04": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ACH", "BOP", "CHK", "FWT", "SWT"}},
				"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CCP", "CTX"}},
				"06": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01", "04"}},
				"07": {Mask: rules.MASK_OPTIONAL},
				"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"ALC", "DA"}},
				"09": {Mask: rules.MASK_OPTIONAL},
				"10": {Mask: rules.MASK_OPTIONAL},
				"11": {Mask: rules.MASK_OPTIONAL},
				"12": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01", "04"}},
				"13": {Mask: rules.MASK_OPTIONAL},
				"14": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"DA", "SG"}},
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
			Description: "REASSOCIATION KEY",
			Mask:        rules.MASK_REQUIRED,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1", "3"}},
				"02": {Mask: rules.MASK_REQUIRED},
				"03": {Mask: rules.MASK_OPTIONAL},
				"04": {Mask: rules.MASK_OPTIONAL},
			},
		},
		2: rules.SegmentRule{
			Name:        "CUR",
			Description: "NON-US DOLLARS CURRENCY",
			Mask:        rules.MASK_OPTIONAL,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"2B", "PR"}},
				"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"MXP", "CAD", "USD"}},
				"03": {Mask: rules.MASK_OPTIONAL},
			},
		},
		3: rules.SegmentRule{
			Name:        "REF",
			Description: "PREMIUM RECEIVERS IDENTIFICATION KEY",
			Mask:        rules.MASK_OPTIONAL,
			RepeatCount: rules.MAXCOUNT,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"14", "18", "2F", "38", "72"}},
				"02": {Mask: rules.MASK_REQUIRED},
			},
		},
		4: rules.SegmentRule{
			Name:        "DTP",
			Description: "PROCESS DAY",
			Mask:        rules.MASK_OPTIONAL,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"009"}},
				"02": {Mask: rules.MASK_REQUIRED},
			},
		},
		5: rules.SegmentRule{
			Name:        "DTP",
			Description: "DELIVERY DATE",
			Mask:        rules.MASK_OPTIONAL,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"035"}},
				"02": {Mask: rules.MASK_REQUIRED},
			},
		},
		6: rules.SegmentRule{
			Name:        "DTP",
			Description: "COVERAGE PERIOD",
			Mask:        rules.MASK_OPTIONAL,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"582"}},
				"02": {Mask: rules.MASK_REQUIRED},
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
			Segments: L1000ARule,
			Mask:     rules.MASK_REQUIRED,
			Name:     "1000A",
		},
		1: {
			Segments: L1000BRule,
			Mask:     rules.MASK_REQUIRED,
			Name:     "1000B",
		},
		2: {
			Segments: L2000ARule,
			Mask:     rules.MASK_OPTIONAL,
			Name:     "2000A",
			Composite: rules.LoopSetRule{
				0: {
					Segments: L2300ARule,
					Mask:     rules.MASK_REQUIRED,
					Name:     "2300A",
					Composite: rules.LoopSetRule{
						0: {
							Segments: L2310ARule,
							Mask:     rules.MASK_OPTIONAL,
							Name:     "2310A",
							Composite: rules.LoopSetRule{
								0: {
									Segments: L2315ARule,
									Mask:     rules.MASK_OPTIONAL,
									Name:     "2315A",
								},
							},
						},
						1: {
							Segments: L2320ARule,
							Mask:     rules.MASK_OPTIONAL,
							Name:     "2320A",
						},
					},
				},
			},
		},
		3: {
			Segments:    L2000BRule,
			Mask:        rules.MASK_OPTIONAL,
			Name:        "2000B",
			RepeatCount: rules.MAXCOUNT,
			Composite: rules.LoopSetRule{
				0: {
					Segments: L2100BRule,
					Mask:     rules.MASK_OPTIONAL,
					Name:     "2100B",
				},
				1: {
					Segments: L2300BRule,
					Mask:     rules.MASK_OPTIONAL,
					Name:     "2300B",
					Composite: rules.LoopSetRule{
						0: {
							Segments:    L2320BRule,
							Mask:        rules.MASK_OPTIONAL,
							Name:        "2320B",
							RepeatCount: rules.MAXCOUNT,
						},
					},
				},
			},
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
			"08": {AcceptValues: []string{"004010X061A1"}},
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
	Name: "820(004010X061A1)",
	ISA: rules.SegmentRule{
		Name:        "ISA",
		Description: "INTERCHANGE CONTROL HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"00", "03"}},
			"03": {AcceptValues: []string{"00", "01"}},
			"05": {AcceptValues: []string{"01", "14", "20", "27", "28", "29", "30", "33", "ZZ"}},
			"07": {AcceptValues: []string{"01", "14", "20", "27", "28", "29", "30", "33", "ZZ"}},
			"11": {AcceptValues: []string{"U"}},
			"12": {AcceptValues: []string{"00401"}},
			"14": {AcceptValues: []string{"0", "1"}},
			"15": {AcceptValues: []string{"P", "T"}},
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
