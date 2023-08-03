// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// 005010X224A2

package rule_5010_837d

import "github.com/moov-io/x12/pkg/rules"

var L1000ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUBMITTER NAME-1000A",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
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
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"IC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {AcceptValues: []string{"EM", "FX", "TE"}},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EX", "EM", "FX", "TE"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"EX", "EM", "FX", "TE"}},
			"08": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var L1000BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "RECEIVER NAME-1000B",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
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

var L2000ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "HL",
		Description: "HIERARCHICAL LEVEL",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {AcceptValues: []string{"20"}},
			"04": {AcceptValues: []string{"1"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "PRV",
		Description: "BILLING / PAY-TO PROVIDER SPECIALTY 2000A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"BI"}},
			"02": {AcceptValues: []string{"PXC"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "CUR",
		Description: "BILLING / PAY-TO PROVIDER SPECIALTY HIERARCHICAL LEVEL 2000A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"85"}},
		},
	},
}

var L2010AARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "BILLING PROVIDER NAME 2010AA",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"85"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "BILLING PROVIDER ADDRESS",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "BILLING PROVIDER CITY, STATE, ZIP CODE",
		Mask:        rules.MASK_REQUIRED,
		Elements:    rules.ElementSetRule{},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "BILLING PROVIDER SECONDARY IDENTIFICATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"EI", "SY"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "BILLING PROVIDER SECONDARY IDENTIFICATION",
		RepeatCount: 2,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "PER",
		Description: "BILLING PROVIDER CONTACT INFORMATION",
		RepeatCount: 2,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"IC"}},
			"03": {AcceptValues: []string{"EM", "FX", "TE"}},
			"05": {AcceptValues: []string{"EM", "FX", "TE", "EX"}, Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {AcceptValues: []string{"EM", "FX", "TE", "EX"}, Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2010ABRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PAY TO ADDRESS NAME 2010AB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAY-TO ADDRESS 2010AB",
		Mask:        rules.MASK_REQUIRED,
		Elements:    rules.ElementSetRule{},
	},
}

var L2010ACRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PAY TO PLAN NAME 2010AC",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Description: "PAY-TO ADDRESS 2010AC",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAY-TO ADDRESS 2010AC",
		Mask:        rules.MASK_REQUIRED,
		Elements:    rules.ElementSetRule{},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "PAY-TO ADDRESS 2010AC",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 1,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"2U", "FY", "NF"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "PAY-TO ADDRESS 2010AC",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 1,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"EI"}},
		},
	},
}

var L2000BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "HL",
		Description: "SUBSCRIBER HIERARCHICAL LEVEL 2000B",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"03": {AcceptValues: []string{"22"}},
			"04": {AcceptValues: []string{"0", "1"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "SBR",
		Description: "SUBSCRIBER INFORMATION 2000B",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"A", "B", "C", "D", "E", "F", "G", "H", "P", "S", "T", "U"}},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"18"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {AcceptValues: []string{"12", "13", "14", "15", "16", "41", "42", "43", "47"}, Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"11", "12", "13", "14", "15", "16", "17", "AM", "BL", "CH", "CI", "DS", "FI", "HM", "LM", "MA", "MB", "MC", "OF", "TV", "VA", "WC", "ZZ"}},
		},
	},
}

var L2010BARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUBSCRIBER SECONDARY IDENTIFICATION 2010BA",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"IL"}},
			"02": {AcceptValues: []string{"1", "2"}},
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
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "SUBSCRIBER ADDRESS 2010BA",
		Mask:        rules.MASK_OPTIONAL,
		Elements:    rules.ElementSetRule{},
	},
	3: rules.SegmentRule{
		Name:        "DMG",
		Description: "SUBSCRIBER DEMOGRAPHIC INFORMATION 2010BA",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"D8"}},
			"03": {AcceptValues: []string{"F", "M", "U"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "PAY-TO ADDRESS 2010AB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"SY"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "REF",
		Description: "PAY-TO ADDRESS 2010AB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"Y4"}},
		},
	},
}

var L2010BBRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PAYER NAME",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
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
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAYER CITY, STATE, ZIP CODE",
		Mask:        rules.MASK_OPTIONAL,
		Elements:    rules.ElementSetRule{},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "PAY-TO ADDRESS 2010AB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"2U", "EI", "FY", "NF"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "PAY-TO ADDRESS 2010AB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G2", "LU"}},
		},
	},
}

var L2300Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "CLM",
		Description: "CLAIM INFORMATION 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {
				Mask: rules.MASK_REQUIRED,
				Composite: rules.ElementSetRule{
					"02": {AcceptValues: []string{"B"}},
				},
			},
			"06": {AcceptValues: []string{"Y", "N"}},
			"07": {AcceptValues: []string{"A", "C"}},
			"08": {AcceptValues: []string{"N", "W", "Y"}},
			"09": {AcceptValues: []string{"I", "Y"}},
			"10": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"P"}},
			"11": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"AA", "EM", "OA"}},
					"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AA", "EM", "OA"}},
					"03": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"AA", "EM", "OA"}},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
				},
			},
			"12": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01", "02", "03", "05"}},
			"13": {Mask: rules.MASK_NOTUSED},
			"14": {Mask: rules.MASK_NOTUSED},
			"15": {Mask: rules.MASK_NOTUSED},
			"16": {Mask: rules.MASK_NOTUSED},
			"17": {Mask: rules.MASK_NOTUSED},
			"18": {Mask: rules.MASK_NOTUSED},
			"19": {Mask: rules.MASK_OPTIONAL},
			"20": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE ONSET OF CURRENT ILLNESS OR SYMPTOM",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"439"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE ONSET OF CURRENT ILLNESS OR SYMPTOM",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"452"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	3: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE ONSET OF CURRENT ILLNESS OR SYMPTOM",
		RepeatCount: 1,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"472"}},
			"02": {AcceptValues: []string{"D8", "RD8"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE ONSET OF CURRENT ILLNESS OR SYMPTOM",
		RepeatCount: 1,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"050"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "DN1",
		Description: "DATE OF DISCHARGE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	6: rules.SegmentRule{
		Name:        "DN2",
		RepeatCount: 35,
		Description: "DATE OF DISCHARGE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"02": {AcceptValues: []string{"E", "M"}},
		},
	},
	7: rules.SegmentRule{
		Name:        "PWK",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2300",
		RepeatCount: 10,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"B4", "DA", "DG", "EB", "OZ", "P6", "RB", "RR"}},
			"02": {AcceptValues: []string{"AA", "BM", "EL", "EM", "FT", "FX"}},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AC"}},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	8: rules.SegmentRule{
		Name:        "CN1",
		Description: "CONTACT INFORMATION 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"02", "03", "04", "05", "06", "09"}},
		},
	},
	9: rules.SegmentRule{
		Name:        "AMT",
		Description: "PATIENT AMOUNT PAID 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"F5"}},
		},
	},
	10: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G3"}},
		},
	},
	11: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"4N"}},
			"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "6", "7"}},
		},
	},
	12: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"F8"}},
		},
	},
	13: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"F9"}},
		},
	},
	14: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G1"}},
		},
	},
	15: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"9A"}},
		},
	},
	16: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"9C"}},
		},
	},
	17: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"D9"}},
		},
	},
	18: rules.SegmentRule{
		Name:        "K3",
		Description: "FILE INFORMATION 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements:    rules.ElementSetRule{},
	},
	19: rules.SegmentRule{
		Name:        "NTE",
		Description: "CLAIM NOTE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"ADD"}},
		},
	},
	20: rules.SegmentRule{
		Name:        "HI",
		Description: "HEALTH CARE DIAGNOSIS CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {
				Mask: rules.MASK_REQUIRED,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABK", "BK", "TQ"}},
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
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABF", "BF", "TQ"}},
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
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABF", "BF", "TQ"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"04": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ABF", "BF", "TQ"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"05": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"06": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"07": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"08": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"09": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"10": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"11": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"ABF", "BF"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
			"12": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
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
	21: rules.SegmentRule{
		Name:        "HCP",
		Description: "CLAIM PRICING 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements:    rules.ElementSetRule{},
	},
}

var L2310ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "REFERRING PROVIDER NAME 2310A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Name:        "RRV",
		Description: "REFERRING PROVIDER SPECIALTY INFORMATION 2310A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"RF"}},
			"02": {AcceptValues: []string{"PXC"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "REF",
		Description: "REFERRING PROVIDER SECONDARY IDENTIFICATION 2310A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2"}},
		},
	},
}

var L2310BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "RENDERING PROVIDER NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"PE"}},
			"02": {AcceptValues: []string{"PXC"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "REF",
		Description: "RENDERING PROVIDER NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2310CRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SERVICE FACILITY LOCATION 2310C",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Mask:        rules.MASK_REQUIRED,
		Description: "FACILITY LOCATION ADDRESS",
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "FACILITY LOCATION CITY/STATE/ZIP",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "LOCATION SECONDARY IDENTIFICATION",
		RepeatCount: 3,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "G2", "LU"}},
		},
	},
}

var L2310DRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "ASSISTANT SURGEON NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"DD"}},
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
		Name:        "RRV",
		Description: "ASSISTANT SURGEON SPECIALTY INFORMATION 2310D",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"AS"}},
			"02": {AcceptValues: []string{"PXC"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "REF",
		Description: "ASSISTANT SURGEON SECONDARY IDENTIFICATION 2310D",
		RepeatCount: 4,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2310ERule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUPERVISING PROVIDER NAME 2310E",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Description: "SUPERVISING PROVIDER SECONDARY IDENTIFICATION 2310E",
		RepeatCount: 4,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2320Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "SBR",
		Description: "OTHER SUBSCRIBER INFORMATION 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"A", "B", "C", "D", "E", "F", "G", "H", "P", "S", "T", "U"}},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01", "18", "19", "20", "21", "39", "40", "53", "G8"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"12", "13", "14", "15", "16", "41", "42", "43", "47"}},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"11", "12", "13", "14", "15", "16", "17", "AM", "BL", "CH", "CI", "DS", "FI", "HM", "LM", "MA", "MB", "MC", "OF", "TV", "VA", "WC", "ZZ"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "CAS",
		Description: "CLAIM LEVEL ADJUSTMENTS 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CO", "CR", "OA", "PI", "PR"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "AMT",
		Description: "COB PAYER PAID AMOUNT 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"D"}},
		},
	},
	3: rules.SegmentRule{
		Name:        "AMT",
		Description: "REMAINING PATIENT LIABILITY 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"EAF"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "AMT",
		Description: "COD TOTAL NON COVERED 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"A8"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "OI",
		Description: "OTHER INSURANCE COVERAGE INFORMATION 2320",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_NOTUSED},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"N", "W", "Y"}},
			"04": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"P"}},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"I", "Y"}},
		},
	},
	6: rules.SegmentRule{
		Name:        "MOA",
		Description: "MEDICARE OUTPAIENT ADJUDICATION INFORMATION 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"08": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var L2330ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "AMBULANCE DROP-OFF LOCATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"IL"}, Mask: rules.MASK_REQUIRED},
			"02": {AcceptValues: []string{"1", "2"}, Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {AcceptValues: []string{"II", "MI"}, Mask: rules.MASK_REQUIRED},
			"09": {Mask: rules.MASK_REQUIRED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "AMBULANCE DROP-OFF LOCATION ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE DROP-OFF LOCATION CITY/STATE/ZIP",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "BILLING PROVIDER SECONDARY IDENTIFICATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"SY"}},
		},
	},
}

var L2330BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYER NAME 2330B",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PR"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"2"}},
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
		Description: "OTHER PAYER ADDRESS 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "OTHER PAYER CITY 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM CHECK OR REMITTANCE DATE 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"573"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYER SECONDARY IDENTIFIER 2330B",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"2U", "EI", "FY", "NF"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYER PRIOR AUTHORIZATION 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G1"}},
		},
	},
	6: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYER REFERRAL NUMBER 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"9F"}},
		},
	},
	7: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"T4"}},
			"02": {AcceptValues: []string{"Y"}},
		},
	},
	8: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G3"}},
		},
	},
	9: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"F8"}},
		},
	},
}

var L2330CRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2"}},
		},
	},
}

var L2330DRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"82"}},
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
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "IG", "G2", "LU"}},
		},
	},
}

var L2330ERule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2330FRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"85"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_NOTUSED},
		},
	},
	1: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G2", "LU"}},
		},
	},
}

var L2330GRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G2", "LU"}},
		},
	},
}

var L2330HRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"DD"}},
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
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2400Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "LX",
		Description: "SERVICE LINE 2400",
		Mask:        rules.MASK_REQUIRED,
		Elements:    rules.ElementSetRule{},
	},
	1: rules.SegmentRule{
		Name:        "SV3",
		Description: "DENTAL SERVICE 2400",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, Composite: rules.ElementSetRule{}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL, Composite: rules.ElementSetRule{}},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"I", "R"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_NOTUSED},
			"10": {Mask: rules.MASK_NOTUSED},
			"11": {Mask: rules.MASK_OPTIONAL, Composite: rules.ElementSetRule{}},
		},
	},
	2: rules.SegmentRule{
		Name:        "TOO",
		Description: "TOOTH INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL, Composite: rules.ElementSetRule{}},
		},
	},
	3: rules.SegmentRule{
		Name:        "DTP",
		Description: "SERVICE DATE 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"472"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "DTP",
		Description: "PRIOR PLACEMENT 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"139", "441"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "DTP",
		Description: "PRIOR PLACEMENT 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"452"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	6: rules.SegmentRule{
		Name:        "DTP",
		Description: "PRIOR PLACEMENT 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"446"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	7: rules.SegmentRule{
		Name:        "DTP",
		Description: "PRIOR PLACEMENT 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"196"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	8: rules.SegmentRule{
		Name:        "DTP",
		Description: "PRIOR PLACEMENT 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"198"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	9: rules.SegmentRule{
		Name:        "CN1",
		Description: "CONTACT INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"02", "03", "04", "05", "06", "09"}},
		},
	},
	10: rules.SegmentRule{
		Name:        "REF",
		Description: "INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 5,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G3"}},
		},
	},
	11: rules.SegmentRule{
		Name:        "REF",
		Description: "INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 5,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G1"}},
		},
	},
	12: rules.SegmentRule{
		Name:        "REF",
		Description: "INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"6R"}},
		},
	},
	13: rules.SegmentRule{
		Name:        "REF",
		Description: "INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"9A"}},
		},
	},
	14: rules.SegmentRule{
		Name:        "REF",
		Description: "INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"9C"}},
		},
	},
	15: rules.SegmentRule{
		Name:        "REF",
		Description: "INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 5,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"9F"}},
		},
	},
	16: rules.SegmentRule{
		Name:        "AMT",
		Description: "INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"T"}},
		},
	},
	17: rules.SegmentRule{
		Name:        "K3",
		Description: "INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 5,
		Elements:    rules.ElementSetRule{},
	},
	18: rules.SegmentRule{
		Name:        "HCP",
		Description: "INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {AcceptValues: []string{"AD"}},
			"11": {AcceptValues: []string{"UN"}},
			"13": {AcceptValues: []string{"T1", "T2", "T3", "T4", "T5", "T6"}},
			"14": {AcceptValues: []string{"1", "2", "3", "4", "5"}},
			"15": {AcceptValues: []string{"1", "2", "3", "4", "5", "6"}},
		},
	},
}

var L2420ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "RENDERING PROVIDER NAME 2420A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"82"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1", "2"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "PRV",
		Description: "RENDERING PROVIDER NAME 2420A",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PE"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PXC"}},
			"03": {Mask: rules.MASK_REQUIRED},
		},
	},
	2: rules.SegmentRule{
		Name:        "PRE",
		Description: "RENDERING PROVIDER NAME 2420A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"0B", "1G", "G2", "LU"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
}

var L2420BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "RENDERING PROVIDER NAME 2420B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"DD"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "PRV",
		Description: "RENDERING PROVIDER NAME 2420B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AS"}},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"PXC"}},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "PRE",
		Description: "RENDERING PROVIDER NAME 2420B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"0B", "1G", "G2", "LU"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
}

var L2420CRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "RENDERING PROVIDER NAME 2420C",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"DQ"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "PRE",
		Description: "RENDERING PROVIDER NAME 2420C",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"0B", "1G", "G2", "LU"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
}

var L2420DRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SERVICE FACILITY LOCATION NAME 2420D",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"77"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"2"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "SERVICE FACILITY LOCATION NAME 2420D",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "SERVICE FACILITY LOCATION NAME 2420D",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "PRE",
		Description: "RENDERING PROVIDER NAME 2420D",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1G", "G2", "LU"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
}

var L2430Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "SVD",
		Description: "LINE ADJUDICATION INFORMATION 2430",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AD", "ER"}},
					"02": {Mask: rules.MASK_OPTIONAL},
					"03": {Mask: rules.MASK_OPTIONAL},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
					"06": {Mask: rules.MASK_OPTIONAL},
					"07": {Mask: rules.MASK_OPTIONAL},
				},
			},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "CAS",
		Description: "LINE ADJUDICATION INFORMATION 2430",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"CO", "CR", "OA", "PI", "PR"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
		},
	},
	2: rules.SegmentRule{
		Name:        "DTP",
		Description: "LINE ADJUDICATION INFORMATION 2430",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"573"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"D8"}},
			"03": {Mask: rules.MASK_REQUIRED},
		},
	},
	3: rules.SegmentRule{
		Name:        "AMT",
		Description: "LINE ADJUDICATION INFORMATION 2430",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"EAF"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var L2000CRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "HL",
		Description: "PATIENT HIERARCHICAL LEVEL",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"23"}},
			"04": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"0"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "PAT",
		Description: "PATIENT HIERARCHICAL LEVEL",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "19", "20", "21", "39", "40", "53", "G8"}},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"D8"}},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"01"}},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"Y"}},
		},
	},
}

var L2010CARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
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
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "DMG",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"D8"}},
			"03": {AcceptValues: []string{"F", "M", "U"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"Y4"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "REF",
		Description: "CLAIM INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"1W", "SY"}},
		},
	},
}

var L2300Composite = rules.LoopSetRule{
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
		Segments:    L2320Rule,
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Name:        "2320",
		Composite: rules.LoopSetRule{
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
			7: {
				Segments: L2330HRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "2330H",
			},
		},
	},
	6: {
		Segments:    L2400Rule,
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 50,
		Name:        "2400",
		Composite: rules.LoopSetRule{
			0: {
				Segments: L2420ARule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "2420A",
			},
			1: {
				Segments: L2420BRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "2420B",
			},
			2: {
				Segments: L2420CRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "2420C",
			},
			3: {
				Segments: L2420DRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "2420D",
			},
			4: {
				Segments:    L2430Rule,
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 15,
				Name:        "2430",
			},
		},
	},
}

var TransactionSetRule = rules.TransactionRule{
	ST: rules.SegmentRule{
		Name:        "ST",
		Description: "TRANSACTION SET HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"837"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {AcceptValues: []string{"005010X224A2"}, Mask: rules.MASK_REQUIRED},
		},
	},
	Composite: rules.LoopRule{
		Name: "Transaction Loop",
		Mask: rules.MASK_REQUIRED,
		Segments: rules.SegmentSetRule{
			0: rules.SegmentRule{
				Name:        "BHT",
				Description: "BEGINNING OF HIERARCHICAL TRANSACTION",
				Mask:        rules.MASK_REQUIRED,
				Elements: rules.ElementSetRule{
					"01": {AcceptValues: []string{"0019"}},
					"02": {AcceptValues: []string{"00", "18"}},
					"06": {AcceptValues: []string{"31", "CH", "RP"}},
				},
			},
		},
		Composite: map[int]rules.LoopRule{
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
				RepeatCount: rules.GREATER_THAN_ONE,
				Name:        "2000A",
				Composite: rules.LoopSetRule{
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
						RepeatCount: rules.GREATER_THAN_ONE,
						Name:        "2000B",
						Composite: rules.LoopSetRule{
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
								Mask:        rules.MASK_OPTIONAL,
								RepeatCount: 100,
								Name:        "2300",
								Composite:   L2300Composite,
							},
							3: {
								Segments: L2000CRule,
								Mask:     rules.MASK_OPTIONAL,
								Name:     "2000C",
								Composite: rules.LoopSetRule{
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
										Composite:   L2300Composite,
									},
								},
							},
						},
					},
				},
			},
		},
	},
	SE: rules.SegmentRule{
		Name:        "SE",
		Description: "TRANSACTION SET TRAILER",
		Mask:        rules.MASK_REQUIRED,
		Elements:    rules.ElementSetRule{},
	},
}

var GroupRule = rules.GroupRule{
	GS: rules.SegmentRule{
		Name:        "GS",
		Description: "FUNCTIONAL GROUP HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"HC"}},
			"07": {AcceptValues: []string{"X"}},
			"08": {AcceptValues: []string{"005010X224A2"}},
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
	Name: "837D(005010X224A2)",
	ISA: rules.SegmentRule{
		Name:        "ISA",
		Description: "INTERCHANGE CONTROL HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00", "03"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00", "01"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "14", "20", "27", "28", "29", "30", "33", "ZZ"}},
			"06": {Mask: rules.MASK_REQUIRED},
			"07": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "14", "20", "27", "28", "29", "30", "33", "ZZ"}},
			"08": {Mask: rules.MASK_REQUIRED},
			"09": {Mask: rules.MASK_REQUIRED},
			"10": {Mask: rules.MASK_REQUIRED},
			"11": {Mask: rules.MASK_REQUIRED},
			"12": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00501"}},
			"13": {Mask: rules.MASK_REQUIRED},
			"14": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"0", "1"}},
			"15": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"P", "T"}},
			"16": {Mask: rules.MASK_REQUIRED},
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
