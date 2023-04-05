// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// 005010X222A1

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
			"03": {AcceptValues: []string{"TE"}},
			"05": {AcceptValues: []string{"EM"}, Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
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
		Description: "HIERARCHICAL LEVEL 2000A",
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
		Description: "FOREIGN CURRENCY INFORMATION 2000A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"85"}},
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
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "BILLING PROVIDER ADDRESS 2010AA",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "BILLING PROVIDER CITY, STATE, ZIP CODE 2010AA",
		Mask:        rules.MASK_REQUIRED,
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
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"IC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"EM", "FX", "TE"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EX", "EM", "FX", "TE"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EX", "EM", "FX", "TE"}},
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
		Mask:        rules.MASK_REQUIRED,
		Description: "PAY-TO ADDRESS 2010AB",
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
		Description: "PAY-TO ADDRESS 2010AB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"2U", "FY", "NF"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "PAY-TO ADDRESS 2010AB",
		Mask:        rules.MASK_OPTIONAL,
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
		Elements: rules.ElementSetRule{
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

var L2010BARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUBSCRIBER SECONDARY IDENTIFICATION 2010BA",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
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
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "SUBSCRIBER ADDRESS 2010BA",
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
		Description: "SUBSCRIBER SECONDARY IDENTIFICATION 2010BA",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"SY"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
	5: rules.SegmentRule{
		Name:        "REF",
		Description: "SUBSCRIBER SECONDARY IDENTIFICATION 2010BA",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"Y4"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
	6: rules.SegmentRule{
		Name:        "REF",
		Description: "SUBSCRIBER CONTACT INFORMATION 2010BA",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"IC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EX"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var L2010BBRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PAYER NAME 2010BB",
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
		Description: "PAYER CITY, STATE, ZIP CODE 2010BB",
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
		Description: "SECONDARY IDENTIFICATION 2010BB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"2U", "EI", "FY", "NF"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "SECONDARY IDENTIFICATION 2010BB",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"G2", "LU"}},
			"02": {Mask: rules.MASK_REQUIRED},
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
			"07": {AcceptValues: []string{"A", "B", "C"}},
			"08": {AcceptValues: []string{"N", "W", "Y"}},
			"09": {AcceptValues: []string{"I", "Y"}},
			"10": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"P"}},
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
		Description: "DATE ONSET OF CURRENT ILLNESS OR SYMPTOM 2300",
		RepeatCount: 8,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"431", "454", "304", "453", "439", "484", "455", "471"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "DTP",
		Description: "DISABILITY DATES 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"314", "360", "361"}},
			"02": {AcceptValues: []string{"D8", "RD8"}},
		},
	},
	3: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE OF DISCHARGE 2300",
		RepeatCount: 4,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"297", "296", "435", "096"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE OF DISCHARGE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"090", "091"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "DTP",
		Description: "DATE OF DISCHARGE 2300",
		RepeatCount: 2,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"444", "050"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	6: rules.SegmentRule{
		Name:        "PWK",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2300",
		RepeatCount: 10,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"03", "04", "05", "06", "07", "08", "09", "10", "11", "13", "15", "21", "A3", "A4", "AM", "AS", "B2", "B3", "B4", "BR", "BS", "BT", "CB", "CK", "CT", "D2", "DA", "DB", "DG", "DJ", "DS", "EB", "HC", "HR", "I5", "R", "LA", "M1", "MT", "NN", "OB", "OC", "OD", "OE", "OX", "OZ", "P4", "P5", "PE", "PN", "PO", "PQ", "PY", "PZ", "RB", "RR", "RT", "RX", "SG", "V5", "XP"}},
			"02": {AcceptValues: []string{"AA", "BM", "EL", "EM", "FT", "FX"}},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	7: rules.SegmentRule{
		Name:        "CN1",
		Description: "CONTACT INFORMATION 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "02", "03", "04", "05", "06", "09"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	8: rules.SegmentRule{
		Name:        "AMT",
		Description: "PATIENT AMOUNT PAID 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"F5"}},
		},
	},
	9: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"4N"}},
			"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "6", "7"}},
		},
	},
	10: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"F5"}},
			"02": {AcceptValues: []string{"Y", "N"}},
		},
	},
	11: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE AUTHORIZATION EXCEPTION CODE 2300",
		RepeatCount: 12,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"EW", "9F", "G1", "F8", "X4", "9A", "9C", "LX", "D9", "EA", "P4", "1J"}},
		},
	},
	12: rules.SegmentRule{
		Name:        "K3",
		Description: "FILE INFORMATION 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
		},
	},
	13: rules.SegmentRule{
		Name:        "NTE",
		Description: "CLAIM NOTE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"ADD", " CER", " DCP", " DGN", " TPO"}},
		},
	},
	14: rules.SegmentRule{
		Name:        "CR1",
		Description: "AMBULANCE TRANSPORT INFORMATION 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
	15: rules.SegmentRule{
		Name:        "CR2",
		Description: "SERVICE INFORMATION 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_NOTUSED},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"A", "C", "D", "E", "F", "G", "M"}},
			"09": {Mask: rules.MASK_NOTUSED},
			"10": {Mask: rules.MASK_OPTIONAL},
			"11": {Mask: rules.MASK_OPTIONAL},
			"12": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"Y", "N"}},
		},
	},
	16: rules.SegmentRule{
		Name:        "HI",
		Description: "HEALTH CARE DIAGNOSIS CODE 2300",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {
				Composite: rules.ElementSetRule{
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
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
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
					"01": {AcceptValues: []string{"ABF", "BF"}},
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
					"01": {AcceptValues: []string{"ABF", "BF"}},
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
	17: rules.SegmentRule{
		Name:        "HI",
		Description: "HEALTH CARE DIAGNOSIS CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"BP"}},
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
					"01": {AcceptValues: []string{"BO"}},
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
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
		},
	},
	18: rules.SegmentRule{
		Name:        "HI",
		Description: "HEALTH CARE DIAGNOSIS CODE 2300",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
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
					"01": {AcceptValues: []string{"BG"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
				},
			},
		},
	},
	19: rules.SegmentRule{
		Name:        "HCP",
		Description: "SERVICE INFORMATION 2300",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00", "01", "02", "03", "04", "05", "07", "08", "09", "10", "11", "12", "13", "14"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_NOTUSED},
			"10": {Mask: rules.MASK_NOTUSED},
			"11": {Mask: rules.MASK_NOTUSED},
			"12": {Mask: rules.MASK_NOTUSED},
			"13": {Mask: rules.MASK_OPTIONAL},
			"14": {Mask: rules.MASK_OPTIONAL},
			"15": {Mask: rules.MASK_OPTIONAL},
		},
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
		Name:        "REF",
		Description: "REFERRING PROVIDER NAME 2310A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2"}},
		},
	},
}

var L2310BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "RENDERING PROVIDER NAME 2310B",
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
		Description: "RENDERING PROVIDER NAME 2310B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"PE"}},
			"02": {AcceptValues: []string{"PXC"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "REF",
		Description: "RENDERING PROVIDER NAME 2310B",
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
		Description: "FACILITY LOCATION ADDRESS 2310C",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "FACILITY LOCATION CITY/STATE/ZIP 2310C",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "LOCATION SECONDARY IDENTIFICATION 2310C",
		RepeatCount: 3,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "G2", "LU"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "PER",
		Description: "FACILITY CONTACT INFORMATION 2310C",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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

var L2310DRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUPERVISING PROVIDER NAME 2310D",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
	1: rules.SegmentRule{
		Name:        "REF",
		Description: "SUPERVISING SECONDARY IDENTIFICATION 2310D",
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
		Description: "AMBULANCE PICK-UP LOCATION 2310E",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Description: "AMBULANCE PICK-UP LOCATION ADDRESS 2310E",
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE PICK-UP LOCATION CITY/STATE/ZIP 2310E",
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2310FRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "AMBULANCE DROP-OFF LOCATION 2310F",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Description: "AMBULANCE DROP-OFF LOCATION ADDRESS 2310F",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE DROP-OFF LOCATION CITY/STATE/ZIP 2310F",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
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
			"02": {AcceptValues: []string{"01", "18", "19", "20", "21", "39", "40", "53", "G8"}},
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
		Description: "OTHER SUBSCRIBER INFORMATION 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"CO", "CR", "OA", "PI", "PR"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "AMT",
		Description: "OTHER SUBSCRIBER INFORMATION 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"D"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "AMT",
		Description: "OTHER SUBSCRIBER INFORMATION 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"A8"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	4: rules.SegmentRule{
		Name:        "OI",
		Description: "OTHER SUBSCRIBER INFORMATION 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_NOTUSED},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"N", "W", "Y"}},
			"04": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"P"}},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"I", "Y"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "MOA",
		Description: "OTHER SUBSCRIBER INFORMATION 2320",
		Mask:        rules.MASK_OPTIONAL,
		Elements:    rules.ElementSetRule{},
	},
}

var L2330ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER SUBSCRIBER NAME 2330A",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
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
		Description: "OTHER SUBSCRIBER ADDRESS 2330A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "OTHER SUBSCRIBER CITY/STATE/ZIP 2330A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER SUBSCRIBER SECONDARY IDENTIFICATION 2330A",
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
		Description: "OTHER PAYER ADDRESS 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "OTHER PAYER CITY/STATE/ZIP 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "DTP",
		Description: "OTHER PAYER DATE 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"573"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYER INFORMATION 2330B",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"2U", "EI", "FY", "NF"}},
		},
	},
	5: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYER INFORMATION 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G1"}},
		},
	},
	6: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYER INFORMATION 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"9F"}},
		},
	},
	7: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYER INFORMATION 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"T4"}},
			"02": {AcceptValues: []string{"Y"}},
		},
	},
	8: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER PAYER INFORMATION 2330B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"F8"}},
		},
	},
}

var L2330CRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330C",
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
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330C",
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
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330D",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330D",
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
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330E",
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
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330E",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "G2", "LU"}},
		},
	},
}

var L2330FRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330F",
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
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330F",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2330GRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330G",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Description: "OTHER PAYMENT REFERRING PROVIDER 2330G",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G2", "LU"}},
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
		Name:        "SV1",
		Description: "SERVICE LINE 2400",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
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
		Description: "SERVICE LINE 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
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
		Name:        "PWK",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		RepeatCount: 10,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"03", "04", "05", "06", "07", "08", "09", "10", "11", "13", "15", "21", "A3", "A4", "AM", "AS", "B2", "B3", "B4", "BR", "BS", "BT", "CB", "CK", "CT", "D2", "DA", "DB", "DG", "DJ", "DS", "EB", "HC", "HR", "I5", "IR", "LA", "M1", "MT", "NN", "OB", "OC", "OD", "OE", "OX", "OZ", "P4", "P5", "PE", "PN", "PO", "PQ", "PY", "PZ", "RB", "RR", "RT", "RX", "SG", "V5", "XP"}},
			"02": {AcceptValues: []string{"AA", "BM", "EL", "FT"}},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "PWK",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		RepeatCount: 10,
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"CT"}},
			"02": {AcceptValues: []string{"AB", "AD", "AF", "AG", "NS"}},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	5: rules.SegmentRule{
		Name:        "CR1",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"LB"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"A", "B", "C", "D", "E"}},
			"05": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"DH"}},
			"06": {Mask: rules.MASK_REQUIRED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
		},
	},
	6: rules.SegmentRule{
		Name:        "CR3",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"I", "R", "S"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"MO"}},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
		},
	},
	7: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"472"}},
			"02": {AcceptValues: []string{"D8", "RD8"}},
		},
	},
	8: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"471"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	9: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"607"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	10: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"463"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	11: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"461"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	12: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"304"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	13: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"738", "739"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	14: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"011"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	15: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"455"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
	16: rules.SegmentRule{
		Name:        "DTP",
		Description: "CLAIM SUPPLEMENTAL INFORMATION 2400",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"454"}},
			"02": {AcceptValues: []string{"D8"}},
		},
	},
}

var L2410Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "LIN",
		Description: "DRUG INFORMATION 2410",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_NOTUSED},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"N4", "EN", "EO", "HI", "ON", "UK", "UP"}},
			"03": {Mask: rules.MASK_REQUIRED},
		},
	},
	1: rules.SegmentRule{
		Name:        "CTP",
		Description: "DRUG INFORMATION 2410",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_NOTUSED},
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED},
		},
	},
	2: rules.SegmentRule{
		Name:        "REF",
		Description: "DRUG INFORMATION 2410",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 20,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"VV", "XZ"}},
		},
	},
}

var L2420ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "RENDERING PROVIDER 2420A",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"82"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "PRV",
		Description: "RENDERING PROVIDER 2420A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"PE"}},
			"02": {AcceptValues: []string{"PXC"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "REF",
		Description: "RENDERING PROVIDER 2420A",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 20,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2420BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "PURCHASED SERVICE PROVIDER 2420B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"QB"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"03": {Mask: rules.MASK_NOTUSED},
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
		Description: "PURCHASED SERVICE PROVIDER 2420B",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 20,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2"}},
		},
	},
}

var L2420CRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SERVICE FACILITY 2420C",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"77"}},
			"02": {AcceptValues: []string{"2"}},
			"03": {Mask: rules.MASK_NOTUSED},
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
		Description: "SERVICE FACILITY 2420C",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "SERVICE FACILITY 2420C",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE FACILITY 2420C",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"G2", "LU"}},
		},
	},
}

var L2420DRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUPERVISING PROVIDER IDENTIFICATION 2420D",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"DQ"}},
			"02": {AcceptValues: []string{"1"}},
			"03": {Mask: rules.MASK_NOTUSED},
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
		Description: "SUPERVISING PROVIDER IDENTIFICATION 2420D",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 20,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2", "LU"}},
		},
	},
}

var L2420ERule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "ORDERING PROVIDER 2420E",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"DK"}},
			"02": {AcceptValues: []string{"1"}},
			"03": {Mask: rules.MASK_NOTUSED},
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
		Description: "ORDERING PROVIDER 2420E",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "ORDERING PROVIDER 2420E",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "ORDERING PROVIDER 2420E",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 20,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2"}},
		},
	},
	4: rules.SegmentRule{
		Name:        "PER",
		Description: "ORDERING PROVIDER 2420E",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"IC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"EM", "FX", "TE"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EX", "EM", "FX", "TE"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EX", "EM", "FX", "TE"}},
			"08": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2420FRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "REFERRING PROVIDER 2420F",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"DN", "P3"}},
			"02": {AcceptValues: []string{"1"}},
			"03": {Mask: rules.MASK_NOTUSED},
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
		Description: "REFERRING PROVIDER 2420F",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 20,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0B", "1G", "G2"}},
		},
	},
}

var L2420GRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "AMBULANCE PICK UP 2420G",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"PW"}},
			"02": {AcceptValues: []string{"2"}},
			"03": {Mask: rules.MASK_NOTUSED},
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
		Description: "AMBULANCE PICK UP 2420G",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE PICK UP 2420G",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2420HRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "AMBULANCE DROP OFF 2420H",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"45"}},
			"02": {AcceptValues: []string{"2"}},
			"03": {Mask: rules.MASK_NOTUSED},
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
		Description: "AMBULANCE DROP OFF 2420H",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "AMBULANCE DROP OFF 2420H",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2430Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "SVD",
		Description: "ADJUDICATION INFORMATION 2430",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {Mask: rules.MASK_REQUIRED},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "CAS",
		Description: "ADJUDICATION INFORMATION 2430",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"CO", "CR", "OA", "PI", "PR"}},
		},
	},
	2: rules.SegmentRule{
		Name:        "DTP",
		Description: "ADJUDICATION INFORMATION 2430",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"573"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"D8"}},
			"03": {Mask: rules.MASK_REQUIRED},
		},
	},
	3: rules.SegmentRule{
		Name:        "AMT",
		Description: "ADJUDICATION INFORMATION 2430",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EAF"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2440Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "LQ",
		Description: "ADJUDICATION INFORMATION 2440",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"AS", "UT"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "CAS",
		Description: "ADJUDICATION INFORMATION 2440",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"02": {AcceptValues: []string{"N", "W", "Y"}},
		},
	},
}

var L2000CRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "HL",
		Description: "PATIENT HIERARCHICAL LEVEL",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"02": {Mask: rules.MASK_NOTUSED},
			"03": {AcceptValues: []string{"23"}},
			"04": {AcceptValues: []string{"0"}},
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
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"D8"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01"}},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"Y"}},
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
	6: rules.SegmentRule{
		Name:        "PER",
		Description: "CLAIM INFORMATION",
		Elements: rules.ElementSetRule{
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
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"837"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {AcceptValues: []string{"005010X222A1"}, Mask: rules.MASK_REQUIRED},
		},
	},
	BHT: rules.SegmentRule{
		Name:        "BHT",
		Description: "BEGINNING OF HIERARCHICAL TRANSACTION",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"0019"}},
			"02": {AcceptValues: []string{"00", "18"}},
			"06": {AcceptValues: []string{"31", "CH", "RP"}},
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
			Segments:    L2000ARule,
			Mask:        rules.MASK_REQUIRED,
			RepeatCount: MAXCOUNT,
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
					RepeatCount: MAXCOUNT,
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
							Mask:        rules.MASK_REQUIRED,
							RepeatCount: 100,
							Name:        "2300",
							Composite: rules.LoopSetRule{
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
									},
								},
								7: {
									Segments:    L2400Rule,
									Mask:        rules.MASK_REQUIRED,
									RepeatCount: 50,
									Name:        "2400",
									Composite: rules.LoopSetRule{
										0: {
											Segments: L2410Rule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2410",
										},
										1: {
											Segments: L2420ARule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2420A",
										},
										2: {
											Segments: L2420BRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2420B",
										},
										3: {
											Segments: L2420CRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2420C",
										},
										4: {
											Segments: L2420DRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2420D",
										},
										5: {
											Segments: L2420ERule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2420E",
										},
										6: {
											Segments: L2420FRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2420F",
										},
										7: {
											Segments: L2420GRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2420G",
										},
										8: {
											Segments: L2420HRule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2420H",
										},
										9: {
											Segments: L2430Rule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2430",
										},
										10: {
											Segments: L2440Rule,
											Mask:     rules.MASK_OPTIONAL,
											Name:     "2440",
										},
									},
								},
							},
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
									Composite: rules.LoopSetRule{
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
											},
										},
										7: {
											Segments:    L2400Rule,
											Mask:        rules.MASK_REQUIRED,
											RepeatCount: 50,
											Name:        "2400",
											Composite: rules.LoopSetRule{
												0: {
													Segments: L2410Rule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2410",
												},
												1: {
													Segments: L2420ARule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2420A",
												},
												2: {
													Segments: L2420BRule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2420B",
												},
												3: {
													Segments: L2420CRule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2420C",
												},
												4: {
													Segments: L2420DRule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2420D",
												},
												5: {
													Segments: L2420ERule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2420E",
												},
												6: {
													Segments: L2420FRule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2420F",
												},
												7: {
													Segments: L2420GRule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2420G",
												},
												8: {
													Segments: L2420HRule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2420H",
												},
												9: {
													Segments: L2430Rule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2430",
												},
												10: {
													Segments: L2440Rule,
													Mask:     rules.MASK_OPTIONAL,
													Name:     "2440",
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
		},
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
			"08": {AcceptValues: []string{"005010X222A1"}},
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

var InterchangeRule = rules.InterChangeRule{
	ISA: rules.SegmentRule{
		Name:        "ISA",
		Description: "INTERCHANGE CONTROL HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
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
		Elements:    rules.ElementSetRule{},
	},
	Group: GroupRule,
}
