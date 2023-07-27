// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// 004010X091A1

package rule_4010_835

import "github.com/moov-io/x12/pkg/rules"

var L1000ARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "PAYER IDENTIFICATION - 1000A",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PR"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"XV"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "PAYER ADDRESS - 1000A",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAYER CITY - 1000A",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "ADDITIONAL PAYER IDENTIFICATION - 1000A",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 4,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"2U", "EO", "HI", "NF"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "PER",
		Description: "PAYER CONTACT INFORMATION - 1000A",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"CX"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EM", "FX", "TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EM", "EX", "FX", "TE"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EX"}},
			"08": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L1000BRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "PAYEE IDENTIFICATION - 1000B",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PE"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"FI", "XX"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "PAYEE ADDRESS - 1000B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAYEE CITY - 1000B",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "ADDITIONAL PAYER IDENTIFICATION - 1000B",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: rules.GREATER_THAN_ONE,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"0B", "1A", "1B", "1C", "1D", "1E", "1F", "1G", "1H", "D3", "G2", "N5", "PQ", "TJ"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var L2000Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "LX",
		Description: "HEADER NUMBER - 2000",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
		},
	},
	1: rules.SegmentRule{
		Name:        "TS3",
		Description: "PROVIDER SUMMARY INFORMATION - 2000",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
			"11": {Mask: rules.MASK_OPTIONAL},
			"12": {Mask: rules.MASK_OPTIONAL},
			"13": {Mask: rules.MASK_OPTIONAL},
			"14": {Mask: rules.MASK_OPTIONAL},
			"15": {Mask: rules.MASK_OPTIONAL},
			"16": {Mask: rules.MASK_OPTIONAL},
			"17": {Mask: rules.MASK_OPTIONAL},
			"18": {Mask: rules.MASK_OPTIONAL},
			"19": {Mask: rules.MASK_OPTIONAL},
			"20": {Mask: rules.MASK_OPTIONAL},
			"21": {Mask: rules.MASK_OPTIONAL},
			"22": {Mask: rules.MASK_OPTIONAL},
			"23": {Mask: rules.MASK_OPTIONAL},
			"24": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "TS2",
		Description: "PROVIDER SUPPLEMENTAL SUMMARY INFORMATION - 2000",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
			"11": {Mask: rules.MASK_OPTIONAL},
			"12": {Mask: rules.MASK_OPTIONAL},
			"13": {Mask: rules.MASK_OPTIONAL},
			"14": {Mask: rules.MASK_OPTIONAL},
			"15": {Mask: rules.MASK_OPTIONAL},
			"16": {Mask: rules.MASK_OPTIONAL},
			"17": {Mask: rules.MASK_OPTIONAL},
			"18": {Mask: rules.MASK_OPTIONAL},
			"19": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var L2100Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "CLP",
		Description: "CLAIM PAYMENT INFORMATION - 2100",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: rules.GREATER_THAN_ONE,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1", "2", "3", "4", "5", "10", "13", "15", "16", "17", "19", "20", "21", "22", "23", "25", "27"}},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"12", "13", "14", "15", "16", "AM", "CH", "DS", "HM", "LM", "MA", "MB", "MC", "OF", "TV", "VA", "WC"}},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_NOTUSED},
			"11": {Mask: rules.MASK_OPTIONAL},
			"12": {Mask: rules.MASK_OPTIONAL},
			"13": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "CAS",
		Description: "CLAIM ADJUSTMENT - 2100",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 99,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"CO", "CR", "OA", "PI", "PR"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
			"11": {Mask: rules.MASK_OPTIONAL},
			"12": {Mask: rules.MASK_OPTIONAL},
			"13": {Mask: rules.MASK_OPTIONAL},
			"14": {Mask: rules.MASK_OPTIONAL},
			"15": {Mask: rules.MASK_OPTIONAL},
			"16": {Mask: rules.MASK_OPTIONAL},
			"17": {Mask: rules.MASK_OPTIONAL},
			"18": {Mask: rules.MASK_OPTIONAL},
			"19": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "NM1",
		Description: "PATIENT NAME - 2100",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"QC"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"34", "HN", "II", "MI", "MR"}},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "NM1",
		Description: "INSURED NAME - 2100",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"IL"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1", "2"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"34", "HN", "MI"}},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "NM1",
		Description: "INSURED NAME - 2100",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"74"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1", "2"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"C"}},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	5: rules.SegmentRule{
		Name:        "NM1",
		Description: "SERVICE PROVIDER NAME - 2100",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"82"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1", "2"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"BD", "BS", "FI", "MC", "PC", "SL", "UP", "XX"}},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	6: rules.SegmentRule{
		Name:        "NM1",
		Description: "CROSSOVER CARRIER NAME - 2100",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"TT"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"2"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AD", "FI", "NI", "PI", "PP", "XV"}},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	7: rules.SegmentRule{
		Name:        "NM1",
		Description: "PRIORITY PAYER NAME - 2100",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PR"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"2"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AD", "FI", "NI", "PI", "PP", "XV"}},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	8: rules.SegmentRule{
		Name:        "MIA",
		Description: "INPATIENT ADJUSTMENT INFORMATION - 2100",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
			"11": {Mask: rules.MASK_OPTIONAL},
			"12": {Mask: rules.MASK_OPTIONAL},
			"13": {Mask: rules.MASK_OPTIONAL},
			"14": {Mask: rules.MASK_OPTIONAL},
			"15": {Mask: rules.MASK_OPTIONAL},
			"16": {Mask: rules.MASK_OPTIONAL},
			"17": {Mask: rules.MASK_OPTIONAL},
			"18": {Mask: rules.MASK_OPTIONAL},
			"19": {Mask: rules.MASK_OPTIONAL},
			"20": {Mask: rules.MASK_OPTIONAL},
			"21": {Mask: rules.MASK_OPTIONAL},
			"22": {Mask: rules.MASK_OPTIONAL},
			"23": {Mask: rules.MASK_OPTIONAL},
			"24": {Mask: rules.MASK_OPTIONAL},
		},
	},
	9: rules.SegmentRule{
		Name:        "MOA",
		Description: "OUTPATIENT ADJUSTMENT INFORMATION - 2100",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
	10: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER CLAIM RELATED INFORMATION - 2100",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 5,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1L", "1W", "9A", "9C", "A6", "BB", "CE", "EA", "F8", "G1", "G3", "IG", "SY"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	11: rules.SegmentRule{
		Name:        "REF",
		Description: "OTHER CLAIM RELATED IDENTIFICATION - 2100",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 5,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1L", "1W", "9A", "9C", "A6", "BB", "CE", "EA", "F8", "G1", "G3", "IG", "SY"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	12: rules.SegmentRule{
		Name:        "REF",
		Description: "RENDERING PROVIDER INFORMATION - 2100",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1A", "1B", "1C", "1D", "1G", "1H", "D3", "G2"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	13: rules.SegmentRule{
		Name:        "DTM",
		Description: "CLAIM DATE - 2100",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 4,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"036", "050", "232", "233"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	14: rules.SegmentRule{
		Name:        "PER",
		Description: "CLAIM CONTACT INFORMATION - 2100",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"CX"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EM", "FX", "TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EM", "EX", "FX", "TE"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EX"}},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_NOTUSED},
		},
	},
	15: rules.SegmentRule{
		Name:        "AMT",
		Description: "CLAIM SUPPLEMENT INFORMATION - 2100",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 14,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"AU", "D8", "DY", "F5", "I", "NL", "T", "T2", "ZK", "ZL", "ZM", "ZN", "ZO", "ZZ"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	16: rules.SegmentRule{
		Name:        "QTY",
		Description: "CLAIM SUPPLEMENT INFORMATION QUANTITY - 2100",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 14,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"CA", "CD", "LA", "LE", "NA", "NE", "NR", "OU", "PS", "VS", "ZK", "ZL", "ZM", "ZN", "ZO"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
}

var L2110Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "SVC",
		Description: "SERVICE PAYMENT INFORMATION - 2110",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {
				Mask: rules.MASK_REQUIRED,
				Composite: rules.ElementSetRule{
					"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AD", "ER", "HC", "ID", "IV", "N4", "NU", "RB", "ZZ"}},
					"02": {Mask: rules.MASK_OPTIONAL},
					"03": {Mask: rules.MASK_OPTIONAL},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
					"06": {Mask: rules.MASK_OPTIONAL},
					"07": {Mask: rules.MASK_OPTIONAL},
					"08": {Mask: rules.MASK_OPTIONAL},
				},
			},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AD", "ER", "HC", "ID", "IV", "N4", "NU", "RB", "ZZ"}},
					"02": {Mask: rules.MASK_OPTIONAL},
					"03": {Mask: rules.MASK_OPTIONAL},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
					"06": {Mask: rules.MASK_OPTIONAL},
					"07": {Mask: rules.MASK_OPTIONAL},
					"08": {Mask: rules.MASK_OPTIONAL},
				},
			},
			"07": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "DTM",
		Description: "SERVICE DATE - 2110",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"150", "151", "472"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "CAS",
		Description: "SERVICE ADJUSTMENT - 2110",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 99,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"CO", "CR", "OA", "PI", "PR"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
			"11": {Mask: rules.MASK_OPTIONAL},
			"12": {Mask: rules.MASK_OPTIONAL},
			"13": {Mask: rules.MASK_OPTIONAL},
			"14": {Mask: rules.MASK_OPTIONAL},
			"15": {Mask: rules.MASK_OPTIONAL},
			"16": {Mask: rules.MASK_OPTIONAL},
			"17": {Mask: rules.MASK_OPTIONAL},
			"18": {Mask: rules.MASK_OPTIONAL},
			"19": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "SERVICE IDENTIFICATION - 2110",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 7,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1S", "6R", "BB", "E9", "G1", "G3", "LU", "RB"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "RENDERING PROVIDER INFORMATION - 2110",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1A", "1B", "1C", "1D", "1G", "1H", "1J", "HPI", "SY", "TJ"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	5: rules.SegmentRule{
		Name:        "AMT",
		Description: "SERVICE SUPPLEMENT INFORMATION - 2110",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 12,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"B6", "DY", "KH", "NE", "T", "T2", "ZK", "ZL", "ZM", "ZN", "ZO"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	6: rules.SegmentRule{
		Name:        "QTY",
		Description: "SERVICE SUPPLEMENT INFORMATION QUANTITY - 2110",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 6,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"NE", "ZK", "ZL", "ZM", "ZN", "ZO"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
	7: rules.SegmentRule{
		Name:        "LQ",
		Description: "HEALTHCARE REMARK CODES - 2110",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 99,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"HE", "RX"}},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
}

var TransactionSetRule = rules.TransactionRule{
	ST: rules.SegmentRule{
		Name:        "ST",
		Description: "TRANSACTION SET HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {AcceptValues: []string{"835"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	Segments: rules.SegmentSetRule{
		0: rules.SegmentRule{
			Name:        "BPR",
			Description: "FINANCIAL HEADER",
			Mask:        rules.MASK_REQUIRED,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"C", "D", "H", "I", "P", "U", "X"}},
				"02": {Mask: rules.MASK_REQUIRED},
				"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"C", "D"}},
				"04": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ACH", "BOP", "CHK", "FWT", "NON"}},
				"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CCP", "CTX"}},
				"06": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01", "04"}},
				"07": {Mask: rules.MASK_OPTIONAL},
				"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"DA"}},
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
			Description: "REASSOCIATION TRACE NUMBER",
			Mask:        rules.MASK_REQUIRED,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1"}},
				"02": {Mask: rules.MASK_REQUIRED},
				"03": {Mask: rules.MASK_OPTIONAL},
				"04": {Mask: rules.MASK_OPTIONAL},
			},
		},
		2: rules.SegmentRule{
			Name:        "CUR",
			Description: "FOREIGN CURRENCY INFORMATION",
			Mask:        rules.MASK_OPTIONAL,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PR"}},
				"02": {Mask: rules.MASK_REQUIRED},
				"03": {Mask: rules.MASK_OPTIONAL},
			},
		},
		3: rules.SegmentRule{
			Name:        "REF",
			Description: "RECEIVER IDENTIFICATION",
			Mask:        rules.MASK_OPTIONAL,
			RepeatCount: rules.GREATER_THAN_ONE,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"EV"}},
				"02": {Mask: rules.MASK_REQUIRED},
			},
		},
		4: rules.SegmentRule{
			Name:        "REF",
			Description: "VERSION IDENTIFICATION",
			Mask:        rules.MASK_OPTIONAL,
			RepeatCount: rules.GREATER_THAN_ONE,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"F2"}},
				"02": {Mask: rules.MASK_REQUIRED},
			},
		},
		5: rules.SegmentRule{
			Name:        "DTM",
			Description: "PRODUCTION DATE",
			Mask:        rules.MASK_OPTIONAL,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"405"}},
				"02": {Mask: rules.MASK_REQUIRED},
				"03": {Mask: rules.MASK_OPTIONAL},
				"04": {Mask: rules.MASK_OPTIONAL},
				"05": {Mask: rules.MASK_OPTIONAL},
				"06": {Mask: rules.MASK_OPTIONAL},
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
			Segments:    L2000Rule,
			Mask:        rules.MASK_OPTIONAL,
			Name:        "2000",
			RepeatCount: rules.GREATER_THAN_ONE,
			Composite: rules.LoopSetRule{
				0: {
					Segments:    L2100Rule,
					Mask:        rules.MASK_REQUIRED,
					Name:        "2100",
					RepeatCount: rules.GREATER_THAN_ONE,
					Composite: rules.LoopSetRule{
						0: {
							Segments:    L2110Rule,
							Mask:        rules.MASK_OPTIONAL,
							Name:        "2110",
							RepeatCount: 999,
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
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"HP"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED},
			"06": {Mask: rules.MASK_REQUIRED},
			"07": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"X"}},
			"08": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"004010X091A1"}},
		},
	},
	GE: rules.SegmentRule{
		Name:        "GE",
		Description: "FUNCTIONAL GROUP TRAILER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
	Trans: TransactionSetRule,
}

var InterchangeRule = rules.InterchangeRule{
	Name: "835(004010X091A1)",
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
			"11": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"U"}},
			"12": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00401"}},
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
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
	Group: GroupRule,
}