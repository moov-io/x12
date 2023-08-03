package rule_4010_810

import "github.com/moov-io/x12/pkg/rules"

var SummaryRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "TDS",
		Description: "TOTAL MONETARY VALUE SUMMARY",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "TXI",
		Description: "TAX INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
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
		},
	},
	2: rules.SegmentRule{
		Name:        "AMT",
		Description: "TOTAL AMOUNT FOR INVOICE BATCH",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: rules.GREATER_THAN_ONE,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"2"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"C", "D"}},
		},
	},
	3: rules.SegmentRule{
		Name:        "CTT",
		Description: "TRANSACTION TOTAL",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var PayerRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "PAYER PARTY",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PR"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"PI"}},
			"04": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"NOT PROVIDED"}},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "PAYER PARTY ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAYER PARTY LOCATION",
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
	3: rules.SegmentRule{
		Name:        "PER",
		Description: "PAYER ADMINISTRATIVE CONTACT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CN", "OC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var PayeeRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "PAYEE PARTY",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PE"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"1"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "PAYEE PARTY ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "PAYEE PARTY LOCATION",
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
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "INVOICE NUMBER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"I5"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "PER",
		Description: "PAYEE ADMINISTRATIVE CONTACT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CN"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var BuyingRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "BUYING PARTY",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"BY"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"PI"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "BUYING PARTY ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "BUYER PARTY LOCATION",
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
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "BUSINESS INFORMATION FOR SHIPMENT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"BE"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "PER",
		Description: "BUYER ADMINISTRATIVE CONTACT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CN"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var SellingRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "SELLING PARTY",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"SE"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"ZZ"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "SELLING PARTY ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "SELLING PARTY LOCATION",
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
	3: rules.SegmentRule{
		Name:        "PER",
		Description: "SELLER ADMINISTRATIVE CONTACT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CN", "SU"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var ShipFromRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "SHIP FROM PARTY",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"SF"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"58"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "SHIP FROM ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "SHIP FROM LOCATION",
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
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "SHIP FROM FACILITY IDENTIFIER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1J"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "MOVE TYPE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"4M"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	5: rules.SegmentRule{
		Name:        "PER",
		Description: "SHIP FROM ADMINISTRATIVE CONTACT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CN", "SH"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var ShipToRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "SHIP TO PARTY",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ST"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"59"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "SHIP TO ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "SHIP TO LOCATION",
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
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "SHIP TO FACILITY IDENTIFIER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"LU"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "PER",
		Description: "SHIP TO ADMINISTRATIVE CONTACT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CN", "DC", "RE"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var PropertyRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "USER DEFINED NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "USER DEFINED ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "USER DEFINED LOCATION",
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
	3: rules.SegmentRule{
		Name:        "PER",
		Description: "USER DEFINED ADMINISTRATIVE CONTACT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var StopOffRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "STOP OFF NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"RC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"ZZ"}},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N4",
		Description: "STOP OFF LOCATION",
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
	2: rules.SegmentRule{
		Name:        "PER",
		Description: "STOP OFF ADMINISTRATIVE CONTACT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"CN"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var CycleRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "ITD",
		Description: "CYCLE START DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"16"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "ITD",
		Description: "CYCLE END DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"ZZ"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "ITD",
		Description: "FINANCIAL STATUS",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"4"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "ITD",
		Description: "SERVICE COMPLETION DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"11"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "ITD",
		Description: "ADDED 1",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"3"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	5: rules.SegmentRule{
		Name:        "ITD",
		Description: "ADDED 2",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"7"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	6: rules.SegmentRule{
		Name:        "DTM",
		Description: "SCHEDULED PICKUP DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"118"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	7: rules.SegmentRule{
		Name:        "DTM",
		Description: "ACTUAL SHIP DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"011"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	8: rules.SegmentRule{
		Name:        "DTM",
		Description: "REQUESTED DELIVERY DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"002"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	9: rules.SegmentRule{
		Name:        "DTM",
		Description: "ACTUAL DELIVERY DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"035"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	10: rules.SegmentRule{
		Name:        "DTM",
		Description: "SELLER INVOICE DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"003"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	11: rules.SegmentRule{
		Name:        "DTM",
		Description: "PURCHASE ORDER DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"004", "095"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	12: rules.SegmentRule{
		Name:        "DTM",
		Description: "CREATION DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"097", "COM", "007"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var ExtendedRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N9",
		Description: "EXTENDED REFERENCE IDENTIFICATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ZZ"}},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "MSG",
		Description: "MESSAGE TEXT",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var LIT1Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "IT1",
		Description: "BASE LINE ITEM DATA (INVLICE)",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"1"}},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EA"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"PL"}},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"IN"}},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"VN"}},
			"11": {Mask: rules.MASK_OPTIONAL},
			"12": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"KF"}},
			"13": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"FREIGHT", "PRODUCT", "TAX", "SERVICE CHARGE", "EBILL", "FEE"}},
			"14": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"OT", "ZZ"}},
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
			"25": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "QTY",
		Description: "BILL/RATED QUANTITY AND UNIT OF MEASURE",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"OC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "QTY",
		Description: "LOADING QUANTITY AND PACKAGING FORM CODE",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"63"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
		},
	},
	3: rules.SegmentRule{
		Name:        "QTY",
		Description: "CUBIC VOLUME AND UNIT OF MEASURE",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"38"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "QTY",
		Description: "BILL/RATED AS QUANTITY AND UNIT OF MEASURE",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"T5"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
		},
	},
	5: rules.SegmentRule{
		Name:        "QTY",
		Description: "LOADING QUANTITY AND PACKAGING FORM CODE",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"39"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
		},
	},
	6: rules.SegmentRule{
		Name:        "QTY",
		Description: "CUBIC VOLUME AND UNIT OF MEASURE",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"12"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
		},
	},
	7: rules.SegmentRule{
		Name:        "QTY",
		Description: "TOTAL WIGHT",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"TO"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
		},
	},
	8: rules.SegmentRule{
		Name:        "TXI",
		Description: "TAX INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"A"}},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
		},
	},
	9: rules.SegmentRule{
		Name:        "CTP",
		Description: "PRICING INFORMATION - INVOICE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"INV"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {Mask: rules.MASK_OPTIONAL},
					"02": {Mask: rules.MASK_OPTIONAL},
					"03": {Mask: rules.MASK_OPTIONAL},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
					"06": {Mask: rules.MASK_OPTIONAL},
				}},
			"06": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"DIS"}},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
		},
	},
	10: rules.SegmentRule{
		Name:        "CTP",
		Description: "PRICING INFORMATION - ORDER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"OPP"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {Mask: rules.MASK_OPTIONAL},
					"02": {Mask: rules.MASK_OPTIONAL},
					"03": {Mask: rules.MASK_OPTIONAL},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
					"06": {Mask: rules.MASK_OPTIONAL},
				}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
		},
	},
	11: rules.SegmentRule{
		Name:        "CTP",
		Description: "ALTERNATE PRICING INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"ALT"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {Mask: rules.MASK_OPTIONAL},
					"02": {Mask: rules.MASK_OPTIONAL},
					"03": {Mask: rules.MASK_OPTIONAL},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
					"06": {Mask: rules.MASK_OPTIONAL},
				}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var LPIDRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "PID",
		Description: "PRODUCT/ITEM DESCRIPTION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"F"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var LLineItemRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "REF",
		Description: "LINE ITEM PROCESSING TYPE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ZZ"}},
			"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"INF", "LOT", "MRC", "NRC", "OTC", "USG"}},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "REF",
		Description: "CONTRACT NUMBER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"CT"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "REF",
		Description: "PARENT LINE NUMBER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"FJ"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	// TODO
	// need to verify
	/*
		3: rules.SegmentRule{
			Name:        "REF",
			Description: "REFERENCE IDENTIFIER",
			Mask:        rules.MASK_OPTIONAL,
			RepeatCount: 3,
			Elements: rules.ElementSetRule{
				"01": {Mask: rules.MASK_REQUIRED},
				"02": {Mask: rules.MASK_REQUIRED},
				"03": {Mask: rules.MASK_OPTIONAL},
			},
		},
	*/
	3: rules.SegmentRule{
		Name:        "REF",
		Description: "REFERENCE NUMBER",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"9F"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	4: rules.SegmentRule{
		Name:        "REF",
		Description: "DEPARTMENT NUMBER",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"19"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	5: rules.SegmentRule{
		Name:        "REF",
		Description: "REFERENCE TABLE DATA",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 3,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ZZ"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	6: rules.SegmentRule{
		Name:        "REF",
		Description: "MESSAGE INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"L1"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	7: rules.SegmentRule{
		Name:        "REF",
		Description: "BUYER PO/TCN",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"L1"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	8: rules.SegmentRule{
		Name:        "REF",
		Description: "COMMODITY CODE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PG"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	9: rules.SegmentRule{
		Name:        "REF",
		Description: "COMMODITY CODE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PG"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
		},
	},
	10: rules.SegmentRule{
		Name:        "DTM",
		Description: "USER DEFINED DATE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var LSACRule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "SAC",
		Description: "SERVICE",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"S"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"ZZ"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"DO"}},
			"10": {Mask: rules.MASK_OPTIONAL},
			"11": {Mask: rules.MASK_OPTIONAL},
			"12": {Mask: rules.MASK_OPTIONAL},
			"13": {Mask: rules.MASK_OPTIONAL},
			"14": {Mask: rules.MASK_OPTIONAL},
			"15": {Mask: rules.MASK_OPTIONAL},
			"16": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "TXI",
		Description: "TAX INFORMATION",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 10,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"A"}},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
			"10": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var LN1Rule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "N1",
		Description: "USER DEFINED NAME",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
		},
	},
	1: rules.SegmentRule{
		Name:        "N3",
		Description: "USER DEFINED ADDRESS",
		Mask:        rules.MASK_OPTIONAL,
		RepeatCount: 2,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_OPTIONAL},
		},
	},
	2: rules.SegmentRule{
		Name:        "N4",
		Description: "USER DEFINED LOCATION",
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
	3: rules.SegmentRule{
		Name:        "PER",
		Description: "USER DEFINED ADMINISTRATIVE CONTACT",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"TE"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_OPTIONAL},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_OPTIONAL},
		},
	},
}

var LFARule = rules.SegmentSetRule{
	0: rules.SegmentRule{
		Name:        "FA1",
		Description: "TYPE OF FINANCIAL ACCOUNTING DATA",
		Mask:        rules.MASK_OPTIONAL,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ZZ"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "FA2",
		Description: "ACCOUNTING DATA",
		Mask:        rules.MASK_REQUIRED,
		RepeatCount: rules.GREATER_THAN_ONE,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
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
			"01": {AcceptValues: []string{"810"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_NOTUSED},
		},
	},
	Composite: rules.LoopRule{
		Name: "Transaction Loop",
		Mask: rules.MASK_REQUIRED,
		Segments: rules.SegmentSetRule{
			0: rules.SegmentRule{
				Name:        "BIG",
				Description: "BILLING INFORMATION",
				Mask:        rules.MASK_REQUIRED,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_OPTIONAL},
					"02": {Mask: rules.MASK_OPTIONAL},
					"03": {Mask: rules.MASK_OPTIONAL},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
					"06": {Mask: rules.MASK_OPTIONAL},
					"07": {Mask: rules.MASK_OPTIONAL},
					"08": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"54", "22", "53", "SU"}},
					"09": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"9"}},
					"10": {Mask: rules.MASK_OPTIONAL},
				},
			},
			1: rules.SegmentRule{
				Name:        "NTE",
				Description: "SERVICE LEVEL REQUESTED",
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 100,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ORI"}},
					"02": {Mask: rules.MASK_OPTIONAL},
				},
			},
			2: rules.SegmentRule{
				Name:        "NTE",
				Description: "SERVICE LEVEL PROVIDED",
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 100,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"DOI"}},
					"02": {Mask: rules.MASK_OPTIONAL},
				},
			},
			3: rules.SegmentRule{
				Name:        "NTE",
				Description: "PRODUCT NAME",
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 100,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PDS"}},
					"02": {Mask: rules.MASK_OPTIONAL},
				},
			},
			4: rules.SegmentRule{
				Name:        "NTE",
				Description: "FILE REQUEST NAME",
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 100,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"OTH"}},
					"02": {Mask: rules.MASK_OPTIONAL},
				},
			},
			5: rules.SegmentRule{
				Name:        "NTE",
				Description: "SOURCE DOCUMENT USED FOR HEADER DATA",
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 100,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PCS"}},
					"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"MERGED", "ORDER", "INVOICE"}},
				},
			},
			6: rules.SegmentRule{
				Name:        "NTE",
				Description: "SOURCE DOCUMENT USED FOR LINE ITEM DATA",
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 100,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"LIN", "GPI"}},
					"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"ORDER", "DOC TYPE", "INVOICE"}},
				},
			},
			7: rules.SegmentRule{
				Name:        "NTE",
				Description: "SOURCE DOCUMENT USED FOR SERVICE CHANGES",
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 100,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"INV", "APN", "GEN"}},
					"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"INVOICE", "ORDER", "DOCTYPE"}},
				},
			},
			8: rules.SegmentRule{
				Name:        "NTE",
				Description: "SOURCE DOCUMENT USED FOR EBILL LINE ITEM DATA",
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 100,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ADD", "CAR"}},
					"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"EBILL", "GL CODE"}},
				},
			},
			9: rules.SegmentRule{
				Name:        "CUR",
				Description: "CURRENCY",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ZZ"}},
					"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"AUD", "CAD", "CHF", "EUR", "GDP", "JPY", "USD"}},
					"03": {Mask: rules.MASK_OPTIONAL},
				},
			},
			10: rules.SegmentRule{
				Name:        "REF",
				Description: "CUSTOMER ORDER NUMBER/BOL NUMBER",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"CO", "BM"}},
					"02": {Mask: rules.MASK_REQUIRED},
					"03": {Mask: rules.MASK_OPTIONAL},
				},
			},
			11: rules.SegmentRule{
				Name:        "REF",
				Description: "INVOICE/PRO NUMBER",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"IV"}},
					"02": {Mask: rules.MASK_REQUIRED},
					"03": {Mask: rules.MASK_OPTIONAL},
				},
			},
			12: rules.SegmentRule{
				Name:        "REF",
				Description: "BILL NUMBER",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"TN"}},
					"02": {Mask: rules.MASK_REQUIRED},
					"03": {Mask: rules.MASK_OPTIONAL},
				},
			},
			13: rules.SegmentRule{
				Name:        "REF",
				Description: "DOCUMENT TYPE",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"8X"}},
					"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ORDER", "BOL", "EBILL", "INVOICE", "PRO", "FEE", "SUPPLEMENTAL INVOICE"}},
					"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"MATCHING", "BUYER SELF INVOICING", "SELLER INVOICING"}},
				},
			},
			14: rules.SegmentRule{
				Name:        "REF",
				Description: "SYSTEM NUMBER",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"06"}},
					"02": {Mask: rules.MASK_REQUIRED},
					"03": {Mask: rules.MASK_OPTIONAL},
				},
			},
			15: rules.SegmentRule{
				Name:        "REF",
				Description: "INBOUND/OUTBOUND INDICATOR OR BUYER RELEASE NUMBER",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"4C", "RE"}},
					"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"Inbound", "Outbound", "INBOUND", "OUTBOUND"}},
					"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"YES", "NO"}},
				},
			},
			16: rules.SegmentRule{
				Name:        "REF",
				Description: "FINANCIAL STATUS",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1Z"}},
					"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PS", "PI", "AE", "AF", "CA", "DN", "HD", "ON", "OP", "MA"}},
					"03": {Mask: rules.MASK_OPTIONAL},
				},
			},
			17: rules.SegmentRule{
				Name:        "REF",
				Description: "COMPLIANCE COMPLETION EVENT",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ACC"}},
					"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"DV", "AF", "VD"}},
					"03": {Mask: rules.MASK_OPTIONAL},
				},
			},
			18: rules.SegmentRule{
				Name:        "REF",
				Description: "SHIPMENT MODE",
				Mask:        rules.MASK_OPTIONAL,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"ZZ"}},
					"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"LT", "M", "A", "O", "R", "HHG"}},
					"03": {Mask: rules.MASK_OPTIONAL},
				},
			},
			19: rules.SegmentRule{
				Name:        "REF",
				Description: "USER DEFINED REFERENCE DATA",
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 3,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED},
					"02": {Mask: rules.MASK_REQUIRED},
					"03": {Mask: rules.MASK_OPTIONAL},
				},
			},
		},
		Composite: map[int]rules.LoopRule{
			0: {
				Segments: PayerRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "Payer",
			},
			1: {
				Segments: PayeeRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "Payee",
			},
			2: {
				Segments: BuyingRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "Buyer",
			},
			3: {
				Segments: SellingRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "Seller",
			},
			4: {
				Segments: ShipFromRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "ShipFrom",
			},
			5: {
				Segments: ShipToRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "ShipFrom",
			},
			6: {
				Segments: PropertyRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "UserDefinedProperty",
			},
			7: {
				Segments: StopOffRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "StopOff",
			},
			8: {
				Segments: CycleRule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "Cycle",
			},
			9: {
				Segments:    ExtendedRule,
				Mask:        rules.MASK_OPTIONAL,
				RepeatCount: 200,
				Name:        "Extended",
			},
			10: {
				Segments: LIT1Rule,
				Mask:     rules.MASK_OPTIONAL,
				Name:     "LOOP IT1",
				Composite: rules.LoopSetRule{
					0: {
						Segments:    LPIDRule,
						Mask:        rules.MASK_OPTIONAL,
						Name:        "LOOP PID",
						RepeatCount: 100,
					},
					1: {
						Segments: LLineItemRule,
						Mask:     rules.MASK_OPTIONAL,
						Name:     "LOOP LINE ITEM",
					},
					2: {
						Segments:    LSACRule,
						Mask:        rules.MASK_OPTIONAL,
						Name:        "LOOP SAC",
						RepeatCount: 25,
					},
					3: {
						Segments:    LN1Rule,
						Mask:        rules.MASK_OPTIONAL,
						Name:        "LOOP N1",
						RepeatCount: 200,
					},
					4: {
						Segments:    LFARule,
						Mask:        rules.MASK_OPTIONAL,
						Name:        "LOOP FA",
						RepeatCount: rules.GREATER_THAN_ONE,
					},
				},
			},
			11: {
				Segments: SummaryRule,
				Mask:     rules.MASK_REQUIRED,
				Name:     "Summary",
			},
		},
	},
	SE: rules.SegmentRule{
		Name:        "SE",
		Description: "TRANSACTION SET TRAILER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
		},
	},
}

var GroupRule = rules.GroupRule{
	GS: rules.SegmentRule{
		Name:        "GS",
		Description: "FUNCTIONAL GROUP HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"IN"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED},
			"06": {Mask: rules.MASK_REQUIRED},
			"07": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"X"}},
			"08": {Mask: rules.MASK_REQUIRED, AcceptRegex: "^004010.*$"},
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
	Name: "810(004010)",
	ISA: rules.SegmentRule{
		Name:        "ISA",
		Description: "INTERCHANGE CONTROL HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00", "03"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00", "01"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "12", "14", "20", "27", "28", "29", "30", "33", "ZZ"}},
			"06": {Mask: rules.MASK_REQUIRED},
			"07": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "12", "14", "20", "27", "28", "29", "30", "33", "ZZ"}},
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
