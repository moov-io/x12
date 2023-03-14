// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// 005010X222A1

package rule_5010_837p

import "github.com/moov-io/x12/pkg/rules"

var L1000ARule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUBMITTER NAME-1000A",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"41"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {AcceptValues: []string{"46"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "PER",
		Description: "SUBMITTER EDI CONTACT INFORMATION-1000A",
		RepeatCount: 2,
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"IC"}},
			"03": {AcceptValues: []string{"TE"}},
			"05": {AcceptValues: []string{"EM"}, Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var L1000BRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "RECEIVER NAME-1000B",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"40"}},
			"02": {AcceptValues: []string{"2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {AcceptValues: []string{"46"}},
		},
	},
}

var L2000ARule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "HL",
		Description: "HIERARCHICAL LEVEL",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {AcceptValues: []string{"20"}},
			"04": {AcceptValues: []string{"1"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "PRV",
		Description: "BILLING / PAY-TO PROVIDER SPECIALTY 2000A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"BI"}},
			"02": {AcceptValues: []string{"PXC"}},
		},
	},
}

var L2010AARule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "BILLING PROVIDER NAME 2010AA",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"85"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "BILLING PROVIDER ADDRESS",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "BILLING PROVIDER CITY, STATE, ZIP CODE",
		Elements:    map[string]rules.ElementRule{},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "BILLING PROVIDER SECONDARY IDENTIFICATION",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"EI", "SY"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "BILLING PROVIDER SECONDARY IDENTIFICATION",
		RepeatCount: 2,
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0B", "1G"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "PER",
		Description: "BILLING PROVIDER CONTACT INFORMATION",
		RepeatCount: 2,
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"IC"}},
			"03": {AcceptValues: []string{"EM"}},
		},
	},
}

var L2010ABRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PAY TO ADDRESS NAME 2010AB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"87"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_NOTUSED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "PAY-TO ADDRESS 2010AB",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAY-TO ADDRESS 2010AB",
		Elements:    map[string]rules.ElementRule{},
	},
}

var L2010ACRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PAY TO PLAN NAME 2010AC",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"PE"}},
			"02": {AcceptValues: []string{"2"}},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PI", "XV"}},
			"09": {Mask: rules.MASK_REQUIRED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "PAY-TO ADDRESS 2010AB",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAY-TO ADDRESS 2010AB",
		Elements:    map[string]rules.ElementRule{},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "PAY-TO ADDRESS 2010AB",
		RepeatCount: 2,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"2U", "FY", "NF", "EI"}},
		},
	},
}

var L2000BRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "HL",
		Description: "SUBSCRIBER HIERARCHICAL LEVEL 2000B",
		Elements: map[string]rules.ElementRule{
			"03": {AcceptValues: []string{"22"}},
			"04": {AcceptValues: []string{"0", "1"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "SBR",
		Description: "SUBSCRIBER INFORMATION 2000B",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"A", "B", "C", "D", "E", "F", "G", "H", "P", "S", "T", "U"}},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"18"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"11", "12", "13", "14", "15", "16", "17", "AM", "BL", "CH", "CI", "DS", "FI", "HM", "LM", "MA", "MB", "MC", "OF", "TV", "VA", "WC", "ZZ"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "PAT",
		Description: "SUBSCRIBER INFORMATION 2000B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {Mask: rules.MASK_NOTUSED},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"D8"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01"}},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"Y"}},
		},
	},
}

var L2010BARule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUBSCRIBER SECONDARY IDENTIFICATION 2010BA",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"IL"}},
			"02": {AcceptValues: []string{"1"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"II", "MI"}},
			"09": {Mask: rules.MASK_REQUIRED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "SUBSCRIBER ADDRESS 2010BA",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "SUBSCRIBER ADDRESS 2010BA",
		Mask:        rules.MASK_OPTIONAL,
		Elements:    map[string]rules.ElementRule{},
	},
	3: rules.SegmentRule{
		Name:        "DMG",
		Description: "SUBSCRIBER DEMOGRAPHIC INFORMATION 2010BA",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"D8"}},
			"03": {AcceptValues: []string{"F", "M", "U"}},
		},
	},
}

var L2010BBRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PAYER NAME",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"PR"}},
			"02": {AcceptValues: []string{"2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PI", "XV"}},
			"09": {Mask: rules.MASK_REQUIRED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "PAYER ADDRESS 2010BB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAYER CITY, STATE, ZIP CODE",
		Mask:        rules.MASK_OPTIONAL,
		Elements:    map[string]rules.ElementRule{},
	},
}

var L2300Rule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "CLM",
		Description: "CLAIM INFORMATION 2300",
		Elements: map[string]rules.ElementRule{
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {
				Mask:           rules.MASK_REQUIRED,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"02": {AcceptValues: []string{"B"}},
				},
			},
			"06": {AcceptValues: []string{"Y", "N"}},
			"07": {AcceptValues: []string{"A", "B", "C"}},
			"08": {AcceptValues: []string{"N", "W", "Y"}},
			"09": {AcceptValues: []string{"I", "Y"}},
			"10": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"P"}},
			"11": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"AA", "EM", "OA"}},
					"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AA", "EM", "OA"}},
					"03": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"AA", "EM", "OA"}},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
				},
			},
			"12": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"02", "03", "05", "09"}},
			"13": {Mask: rules.MASK_NOTUSED},
			"14": {Mask: rules.MASK_NOTUSED},
			"15": {Mask: rules.MASK_NOTUSED},
			"16": {Mask: rules.MASK_NOTUSED},
			"17": {Mask: rules.MASK_NOTUSED},
			"18": {Mask: rules.MASK_NOTUSED},
			"19": {Mask: rules.MASK_NOTUSED},
			"20": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE ONSET OF CURRENT ILLNESS OR SYMPTOM",
		RepeatCount: 8,
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"431", "454", "304", "453", "439", "484", "455", "471"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "DTP",
		Description: "DISABILITY DATES",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"314", "360", "361"}},
			"02": {AcceptValues: []string{"D8", "RD8"}},
		},
	},
	3: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE OF DISCHARGE",
		RepeatCount: 4,
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"297", "296", "435", "096"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE OF DISCHARGE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"090", "091"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE OF DISCHARGE",
		RepeatCount: 2,
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"444", "050"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	6: rules.SegmentRule{
		Name:        "PWK",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2300",
		RepeatCount: 10,
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"03", "04", "05", "06", "07", "08", "09", "10", "11", "13", "15", "21", "A3", "A4", "AM", "AS", "B2", "B3", "B4", "BR", "BS", "BT", "CB", "CK", "CT", "D2", "DA", "DB", "DG", "DJ", "DS", "EB", "HC", "HR", "I5", "R", "LA", "M1", "MT", "NN", "OB", "OC", "OD", "OE", "OX", "OZ", "P4", "P5", "PE", "PN", "PO", "PQ", "PY", "PZ", "RB", "RR", "RT", "RX", "SG", "V5", "XP"}},
			"02": {AcceptValues: []string{"AA", "BM", "EL", "EM", "FT", "FX"}},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	7: rules.SegmentRule{
		Name:        "AMT",
		Description: "PATIENT AMOUNT PAID 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"F5"}},
		},
	},
	8: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"4N"}},
			"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "6", "7"}},
		},
	},
	9: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"F5"}},
			"02": {AcceptValues: []string{"Y", "N"}},
		},
	},
	10: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		RepeatCount: 13,
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"EW", "9F", "G1", "F8", "X4", "9A", "9C", "LX", "D9", "EA", "P4", "1J"}},
		},
	},
	11: rules.SegmentRule{
		Name:        "NTE",
		Description: "CLAIM NOTE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"ADD", " CER", " DCP", " DGN", " TPO"}},
		},
	},
	12: rules.SegmentRule{
		Name:        "CR1",
		Description: "AMBULANCE TRANSPORT INFORMATION 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"LB"}, Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {AcceptValues: []string{"A", "B", "C", "D", "E"}},
			"05": {AcceptValues: []string{"DH"}},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
		},
	},
	13: rules.SegmentRule{
		Name:        "HI",
		Description: "HEALTH CARE DIAGNOSIS CODE 2300",
		Elements: map[string]rules.ElementRule{
			"01": {
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABK", "BK"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
			"02": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
			"03": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"04": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"05": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"06": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"07": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"08": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"09": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"10": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"11": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"12": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
		},
	},
}

var L2310ARule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "REFERRING PROVIDER NAME 2310A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"DN", "P3"}},
			"02": {AcceptValues: []string{"1"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "REF",
		Description: "REFERRING PROVIDER NAME 2310A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2"}},
		},
	},
}

var L2310BRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "RENDERING PROVIDER NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"82"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "PRV",
		Description: "RENDERING PROVIDER NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"PE"}},
			"02": {AcceptValues: []string{"PXC"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "REF",
		Description: "RENDERING PROVIDER NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2310CRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SERVICE FACILITY LOCATION 2310C",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"77"}},
			"02": {AcceptValues: []string{"2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "FACILITY LOCATION ADDRESS",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "FACILITY LOCATION CITY/STATE/ZIP",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "LOCATION SECONDARY IDENTIFICATION",
		RepeatCount: 3,
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0B", "G2", "LU"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "PER",
		Description: "FACILITY CONTACT INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"IC"}, Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {AcceptValues: []string{"TE"}, Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {AcceptValues: []string{"EX"}, Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2310DRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUPERVISING PROVIDER NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"DQ"}},
			"02": {AcceptValues: []string{"1"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "REF",
		Description: "SUPERVISING SECONDARY IDENTIFICATION",
		RepeatCount: 4,
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2310ERule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "AMBULANCE PICK-UP LOCATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"PW"}},
			"02": {AcceptValues: []string{"2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "AMBULANCE PICK-UP LOCATION ADDRESS",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE PICK-UP LOCATION CITY/STATE/ZIP",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2310FRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "AMBULANCE DROP-OFF LOCATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"45"}},
			"02": {AcceptValues: []string{"2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "AMBULANCE DROP-OFF LOCATION ADDRESS",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE DROP-OFF LOCATION CITY/STATE/ZIP",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2320Rule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "AMBULANCE DROP-OFF LOCATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"45"}},
			"02": {AcceptValues: []string{"2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "AMBULANCE DROP-OFF LOCATION ADDRESS",
		RepeatCount: 1,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE DROP-OFF LOCATION CITY/STATE/ZIP",
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2330ARule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "AMBULANCE DROP-OFF LOCATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"IL"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "AMBULANCE DROP-OFF LOCATION ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE DROP-OFF LOCATION CITY/STATE/ZIP",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "BILLING PROVIDER SECONDARY IDENTIFICATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"SY"}},
		},
	},
}

var L2330BRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "AMBULANCE DROP-OFF LOCATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"PR"}},
			"02": {AcceptValues: []string{"2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "AMBULANCE DROP-OFF LOCATION ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE DROP-OFF LOCATION CITY/STATE/ZIP",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "DTP",
		Description: "OTHER PAYMENT NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"573"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT NAME",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"2U", "EI", "FY", "NF"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"G1"}},
		},
	},
	6: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"9F"}},
		},
	},
	7: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"T4"}},
			"02": {AcceptValues: []string{"Y"}},
		},
	},
	8: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"F8"}},
		},
	},
}

var L2330CRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"DN", "P3"}},
			"02": {AcceptValues: []string{"1"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 3,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2"}},
		},
	},
}

var L2330DRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"82"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 3,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0B", "IG", "G2", "LU"}},
		},
	},
}

var L2330ERule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"77"}},
			"02": {AcceptValues: []string{"2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 3,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0B", "G2", "LU"}},
		},
	},
}

var L2330FRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"DQ"}},
			"02": {AcceptValues: []string{"1"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 3,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2330GRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"85"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 3,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"G2", "LU"}},
		},
	},
}

var L2400Rule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "LX",
		Description: "SERVICE LINE",
		Mask:        rules.MASK_REQUIRED,
		Elements:    map[string]rules.ElementRule{},
	},
	1: rules.SegmentRule{
		Name:        "SV1",
		Description: "SERVICE LINE",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"MJ", "UN"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_REQUIRED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"Y"}},
			"10": {Mask: rules.MASK_NOTUSED},
			"11": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"Y"}},
			"12": {Mask: rules.MASK_OPTIONAL},
			"13": {Mask: rules.MASK_NOTUSED},
			"14": {Mask: rules.MASK_NOTUSED},
			"15": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"0"}},
			"16": {Mask: rules.MASK_NOTUSED},
			"17": {Mask: rules.MASK_NOTUSED},
			"18": {Mask: rules.MASK_NOTUSED},
			"19": {Mask: rules.MASK_NOTUSED},
			"20": {Mask: rules.MASK_NOTUSED},
			"21": {Mask: rules.MASK_NOTUSED},
		},
	},
	2: rules.SegmentRule{
		Name:        "SV5",
		Description: "SERVICE LINE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"DA"}},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED},
			"06": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1", "4", "6"}},
			"07": {Mask: rules.MASK_NOTUSED},
		},
	},
	3: rules.SegmentRule{
		Name:        "DTP",
		Description: "SERVICE LINE",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"472"}},
			"02": {AcceptValues: []string{"D8", "RD8"}},
		},
	},
}

var L2000CRule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PATIENT HIERARCHICAL LEVEL",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {AcceptValues: []string{"23"}},
			"04": {AcceptValues: []string{"0"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "PAT",
		Description: "PATIENT HIERARCHICAL LEVEL",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "19", "20", "21", "39", "40", "53", "G8"}},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"D8"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01"}},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"Y"}},
		},
	},
}

var L2010CARule = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"QC"}},
			"02": {AcceptValues: []string{"1"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_NOTUSED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "DMG",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"D8"}},
			"03": {AcceptValues: []string{"F", "M", "U"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"Y4"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "REF",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"1W", "SY"}},
		},
	},
	6: rules.SegmentRule{
		Name:        "PER",
		Description: "CLAIM INFORMATION",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"IC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {AcceptValues: []string{"TE"}},
			"05": {AcceptValues: []string{"EX"}, Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
		},
	},
}

const (
	MAXCOUNT = 200
)

var TransactionSetRule = rules.TransactionRule{
	ST: rules.SegmentRule{
		Name:        "ST",
		Description: "TRANSACTION SET HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"837"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {AcceptValues: []string{"005010X222A1"}, Mask: rules.MASK_REQUIRED},
		},
	},
	BHT: rules.SegmentRule{
		Name:        "BHT",
		Description: "BEGINNING OF HIERARCHICAL TRANSACTION",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"0019"}},
			"02": {AcceptValues: []string{"00", "18"}},
			"06": {AcceptValues: []string{"31", "CH", "RP"}},
		},
	},
	SE: rules.SegmentRule{
		Name:        "SE",
		Description: "TRANSACTION SET TRAILER",
		Mask:        rules.MASK_REQUIRED,
		Elements:    map[string]rules.ElementRule{},
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
			Segments:    L2000ARule,
			Mask:        rules.MASK_REQUIRED,
			RepeatCount: MAXCOUNT,
			Name:        "2000A",
			SubLoopRule: map[int]rules.LoopRule{
				0: {
					Segments: L2010AARule,
					Mask:     rules.MASK_REQUIRED,
					Name:     "2010AA",
				},
				1: {
					Segments: L2010ABRule,
					Mask:     rules.MASK_OPTIONAL,
					Name:     "2010AB",
				},
				2: {
					Segments: L2010ACRule,
					Mask:     rules.MASK_OPTIONAL,
					Name:     "2010AC",
				},
				3: {
					Segments:    L2000BRule,
					Mask:        rules.MASK_REQUIRED,
					RepeatCount: MAXCOUNT,
					Name:        "2000B",
					SubLoopRule: map[int]rules.LoopRule{
						0: {
							Segments: L2010BARule,
							Mask:     rules.MASK_REQUIRED,
							Name:     "2010BA",
						},
						1: {
							Segments: L2010BBRule,
							Mask:     rules.MASK_REQUIRED,
							Name:     "2010BB",
						},
						2: {
							Segments:    L2300Rule,
							Mask:        rules.MASK_REQUIRED,
							RepeatCount: 100,
							Name:        "2300",
							SubLoopRule: map[int]rules.LoopRule{
								0: {
									Segments:    L2310ARule,
									Mask:        rules.MASK_OPTIONAL,
									RepeatCount: 2,
									Name:        "2310A",
								},
								1: {
									Segments: L2310BRule,
									Mask:     rules.MASK_OPTIONAL,
									Name:     "2310B",
								},
								2: {
									Segments: L2310CRule,
									Mask:     rules.MASK_OPTIONAL,
									Name:     "2310C",
								},
								3: {
									Segments: L2310DRule,
									Mask:     rules.MASK_OPTIONAL,
									Name:     "2310D",
								},
								4: {
									Segments: L2310ERule,
									Mask:     rules.MASK_OPTIONAL,
									Name:     "2310E",
								},
								5: {
									Segments: L2310FRule,
									Mask:     rules.MASK_OPTIONAL,
									Name:     "2310F",
								},
								6: {
									Segments:    L2320Rule,
									Mask:        rules.MASK_OPTIONAL,
									RepeatCount: 10,
									Name:        "2320",
									SubLoopRule: map[int]rules.LoopRule{
										0: {
											Segments: L2330ARule,
											Mask:     rules.MASK_REQUIRED,
											Name:     "2330A",
										},
										1: {
											Segments: L2330BRule,
											Mask:     rules.MASK_REQUIRED,
											Name:     "2330B",
										},
										2: {
											Segments:    L2330CRule,
											Mask:        rules.MASK_OPTIONAL,
											RepeatCount: 2,
											Name:        "2330C",
										},
										3: {
											Segments: L2330DRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2330D",
										},
										4: {
											Segments: L2330ERule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2330E",
										},
										5: {
											Segments: L2330FRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2330F",
										},
										6: {
											Segments: L2330GRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2330G",
										},
									},
								},
								7: {
									Segments:    L2400Rule,
									Mask:        rules.MASK_REQUIRED,
									RepeatCount: 50,
									Name:        "2400",
								},
							},
						},
						3: {
							Segments: L2000CRule,
							Mask:     rules.MASK_OPTIONAL,
							Name:     "2000C",
							SubLoopRule: map[int]rules.LoopRule{
								0: {
									Segments: L2010CARule,
									Mask:     rules.MASK_REQUIRED,
									Name:     "2010CA",
								},
								1: {
									Segments:    L2300Rule,
									Mask:        rules.MASK_REQUIRED,
									RepeatCount: 100,
									Name:        "2300",
									SubLoopRule: map[int]rules.LoopRule{
										0: {
											Segments:    L2310ARule,
											Mask:        rules.MASK_OPTIONAL,
											RepeatCount: 2,
											Name:        "2310A",
										},
										1: {
											Segments: L2310BRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2310B",
										},
										2: {
											Segments: L2310CRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2310C",
										},
										3: {
											Segments: L2310DRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2310D",
										},
										4: {
											Segments: L2310ERule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2310E",
										},
										5: {
											Segments: L2310FRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2310F",
										},
										6: {
											Segments:    L2320Rule,
											Mask:        rules.MASK_OPTIONAL,
											RepeatCount: 10,
											Name:        "2320",
											SubLoopRule: map[int]rules.LoopRule{
												0: {
													Segments: L2330ARule,
													Mask:     rules.MASK_REQUIRED,
													Name:     "2330A",
												},
												1: {
													Segments: L2330BRule,
													Mask:     rules.MASK_REQUIRED,
													Name:     "2330B",
												},
												2: {
													Segments:    L2330CRule,
													Mask:        rules.MASK_OPTIONAL,
													RepeatCount: 2,
													Name:        "2330C",
												},
												3: {
													Segments: L2330DRule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2330D",
												},
												4: {
													Segments: L2330ERule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2330E",
												},
												5: {
													Segments: L2330FRule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2330F",
												},
												6: {
													Segments: L2330GRule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2330G",
												},
											},
										},
										7: {
											Segments:    L2400Rule,
											Mask:        rules.MASK_REQUIRED,
											RepeatCount: 50,
											Name:        "2400",
										},
									},
								},
							},
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
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"HC"}},
			"07": {AcceptValues: []string{"X"}},
			"08": {AcceptValues: []string{"005010X222A1"}},
		},
	},
	GE: rules.SegmentRule{
		Name:        "GE",
		Description: "FUNCTIONAL GROUP TRAILER",
		Mask:        rules.MASK_REQUIRED,
		Elements:    map[string]rules.ElementRule{},
	},
	Trans: TransactionSetRule,
}

var InterchangeRule = rules.InterChangeRule{
	ISA: rules.SegmentRule{
		Name:        "ISA",
		Description: "INTERCHANGE CONTROL HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"00"}},
			"03": {AcceptValues: []string{"00"}},
			"12": {AcceptValues: []string{"00501"}},
			"14": {AcceptValues: []string{"0"}},
			"15": {AcceptValues: []string{"P", "T"}},
		},
	},
	IEA: rules.SegmentRule{
		Name:        "IEA",
		Description: "INTERCHANGE CONTROL TRAILER",
		Mask:        rules.MASK_REQUIRED,
		Elements:    map[string]rules.ElementRule{},
	},
	Group: GroupRule,
}
