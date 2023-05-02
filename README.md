[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://github.com/moov-io/x12">Project Documentation</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/x12?status.svg)](https://godoc.org/github.com/moov-io/x12)
[![Build Status](https://github.com/moov-io/x12/workflows/Go/badge.svg)](https://github.com/moov-io/x12/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/x12/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/x12)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/x12)](https://goreportcard.com/report/github.com/moov-io/x12)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/x12?label=project%20size)](https://github.com/moov-io/x12)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/x12/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/x12)](https://hub.docker.com/r/moov/x12)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/x12)](https://github.com/moov-io/x12)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/x12

moov-io/x12 is a Go library for parsing and generating Accredited Standards Committee X12 (ASC X12) documents. The goal is to support Electronic data interchange (EDI) Context Inspired Component Architecture (CICA) standards which support health care, insurance, transportation, finance, government, supply chains and many other industries.

## Project status

The x12 is actively being developed for use in production systems. Please star the project if you are interested in its progress. If you need other ASC X12 formats, find bugs, have comments, or suggstions we would appreciate an issue or pull request. Thanks!

## Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.19 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help in setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/iso8583/releases/latest) as well. We highly recommend you use a tagged release for production.

### Installation

```
go get github.com/moov-io/x12
```

### Goal of the moov-io/x12 project
X12 defines and maintains transaction sets that establish the data content exchanged for specific business purposes.
Transaction sets are identified by a numeric identifier and a name.
Each transaction set is maintained by a subcommittee operating within X12’s Accredited Standards Committee.

Moov-io/x12 project will support some kinds of x12 transaction sets
* X12F - Finance / 820 Payment Order & Remittance Advice
* X12C - Communications & Controls / 999 Implementation Acknowledgment
* X12N - Insurance / 837 Health Care Claim
* X12N - Insurance / 835 Health Care Claim Payment & Advice

To support these transaction sets, x12 project used following data unit
* Interchange
  * Interchange includes control header(ISA), control trailer(IEA), functional groups. i.e., segments and groups
* Group
  * Functional group includes group header(GS), group trailer(GE), transaction sets. i.e., segments and transaction sets
* Transaction Set
  * Transaction set includes transaction set header(SA), transaction set trailer(SE), groups, and segments. i.e., segments and loops
* Loop
  * Loop includes composite loops and segments. i.e., segments and loops
* Segment
  * Segment includes data fields

Any x12 transaction set will describe as structured data units with leveled order
Above data hierarchy, we can see that loop has nested sub-loop
Of course each x12 transaction will have different structured data units (loop structure and segments of transaction set)

How to specify structured data units according to x12 transaction type?
How to support specification of x12 transaction with above general data units?

Moov-io/x12 has rule feature to fix above problems.
Each interchange will perform validation check based on specified rule

### Define your rule

Currently, we support following rules for :

* [4010 820](./specs/rule_4010_820/rule.go) - ASC X12F 4010 820 (004010X061A1)
* [5010 837d](./specs/rule_5010_837d/rule.go) - ASC X12N 5010 837 (005010X224A2)
* [5010 837p](./specs/rule_5010_8837p/rule.go) - ASC X12N 5010 837 (005010X222A1)
* [STP 820](./specs/rule_stp_820/rule.go) - ASC X12F STP 820 (004010STP820)

We can specify rule using some kinds of rule struts.

#### InterchangeRule
```
type InterchangeRule struct {
	Name  string
	ISA   SegmentRule
	IEA   SegmentRule
	Group GroupRule
}
```
Interchange rule defined details of control header.

#### GroupRule
```
type GroupRule struct {
	GS    SegmentRule
	GE    SegmentRule
	Trans TransactionRule
}
```
Group rule defined details of group header.

#### TransactionRule
```
type TransactionRule struct {
	ST       SegmentRule
	SE       SegmentRule
	Loops    LoopSetRule
	Segments SegmentSetRule
}
```
Transaction rule defined details of transaction set.
Loops is ordered nested loop list.
order is index of rule map.

#### LoopSetRule
```
type LoopSetRule map[int]LoopRule

type LoopRule struct {
	Segments    SegmentSetRule
	Mask        string
	RepeatCount int
	Name        string
	Composite   LoopSetRule
}
```
Loop rule defined details of loop.
Mask is to specify loop's ability such as required, optional, non-used.
If the mask option is omitted, will deal as required.
RepeatCount is to specify how many times repeat the loop
If the RepeatCount option is omitted, will deal as RepeatCount=1.
Composite is ordered nested sub loops.

#### SegmentSetRule
```
type SegmentSetRule map[int]SegmentRule

type SegmentRule struct {
	Elements    ElementSetRule
	Mask        string
	RepeatCount int
	Name        string
	Description string
}
```
Mask and repeat count are same as loop's mask and repeat count
Elements is ordered fields

#### SegmentSetRule
```
type ElementSetRule map[string]ElementRule

type ElementRule struct {
	Mask         string
	Name         string
	AcceptValues []string
	Composite    ElementSetRule
}
```
Mask and repeat count are same as loop's mask and repeat count.
Composite is ordered nested fields.
Accept values is to specify available values that can use the field.

Element's mask is a bit different from segment and loop.
Each segment has a real struct such as IEA, ISA, GS, etc.
Field of segment is omitted field or required field (check out with json, xml tag).
If the mask option of element rule is omitted, will use original field's property (omitted field or required field).  

#### Rule example
```

import "github.com/moov-io/x12/pkg/rules"

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
	Group: rules.GroupRule{
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
		Trans: rules.TransactionRule{
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
		},
	},
}
```
In above rule, transaction set should have ST, BPR, TRN, NM1, NM1, SE, and can have rules

## Usage

User should create interchange, group, transaction, segments.
```
	isa := segments.ISA{
		AuthorizationQualifier:    "00",
		AuthorizationInformation:  " ",
		SecurityQualifier:         "00",
		SecurityInformation:       " ",
		SenderQualifier:           "30",
		SenderId:                  "227777777 ",
		ReceiverQualifier:         "14",
		ReceiverId:                "577777777 ",
		Date:                      "120530",
		Time:                      "1144",
		StandardsId:               "U",
		Version:                   "00401",
		InterchangeControlNumber:  "000000001",
		AcknowledgmentRequested:   "0",
		TestIndicator:             "P",
		ComponentElementSeparator: "~",
	}

	iea := segments.IEA{}
	iea.SetFieldByIndex("01", "1")
	iea.SetFieldByIndex("02", "000000001")

	gs := segments.GS{}
	gs.SetFieldByIndex("01", "RA")
	gs.SetFieldByIndex("02", "227777777")
	gs.SetFieldByIndex("03", "577777777")
	gs.SetFieldByIndex("04", "20120530")
	gs.SetFieldByIndex("05", "1144")
	gs.SetFieldByIndex("06", "1")
	gs.SetFieldByIndex("07", "X")
	gs.SetFieldByIndex("08", "004010")

	ge := segments.GE{}
	ge.SetFieldByIndex("01", "1")
	ge.SetFieldByIndex("02", "1")

	st := segments.ST{}
	st.SetFieldByIndex("01", "820")
	st.SetFieldByIndex("02", "0001")

	se := segments.SE{}
	se.SetFieldByIndex("01", "6")
	se.SetFieldByIndex("02", "0001")

	bpr := segments.BPR{}
	bpr.SetFieldByIndex("01", "C")
	bpr.SetFieldByIndex("02", "7989.73")
	bpr.SetFieldByIndex("03", "C")
	bpr.SetFieldByIndex("04", "ACH")
	bpr.SetFieldByIndex("05", "CTX")
	bpr.SetFieldByIndex("10", "1657777777")
	bpr.SetFieldByIndex("12", "01")
	bpr.SetFieldByIndex("13", "148529553")
	bpr.SetFieldByIndex("14", "DA")
	bpr.SetFieldByIndex("15", "92283334")
	bpr.SetFieldByIndex("16", "20120531")

	trn := segments.TRN{}
	trn.SetFieldByIndex("01", "1")
	trn.SetFieldByIndex("02", "12053011440000192")

	n11 := segments.N1{}
	n11.SetFieldByIndex("01", "PR")
	n11.SetFieldByIndex("02", "YOUR COMPANY")
	n11.SetFieldByIndex("03", "91")
	n11.SetFieldByIndex("04", "227777777")

	n12 := segments.N1{}
	n12.SetFieldByIndex("01", "PE")
	n12.SetFieldByIndex("02", "WALMART")

	newInterchange := file.Interchange{
		ISA: isa,
		IEA: &iea,
		FunctionalGroups: []file.FunctionalGroup{
			{
				GS: gs,
				GE: &ge,
				TransactionSets: []file.TransactionSet{
					{
						ST:       st,
						SE:       &se,
						Segments: []segments.SegmentInterface{&bpr, &trn, &n11, &n12},
					},
				},
			},
		},
	}

	copyRule := rule_stp_820.InterchangeRule
	if err := newInterchange.Validate(&copyRule); err != nil {
		fmt.Println(err)
	}

	raw := `ISA*00* *00* *30*227777777 *14*577777777 *120530*1144*U*00401*000000001*0*P*~\GS*RA*227777777*577777777*20120530*1144*1*X*004010\ST*820*0001\BPR*C*7989.73*C*ACH*CTX*****1657777777**01*148529553*DA*92283334*20120531\TRN*1*12053011440000192\N1*PR*YOUR COMPANY*91*227777777\N1*PE*WALMART\SE*6*0001\GE*1*1\IEA*1*000000001\`
	if newInterchange.String(segmentTerminator) != raw {
		fmt.Println("invalid string")
	}
	
```

We need to know hierarchy of file, interchange or rule when editing new interchange.

Please use Print() function to verify edited interchange or target rule structure

```

func (f File) Print(w io.Writer)

func (r Interchange) Print(w io.Writer)

func (r InterchangeRule) Print(w io.Writer, isRequiredOnly bool)

```

File print example
```
  DUMP EDI FILE WITH EPN STP 820

INDEX:    | 00 | 01       | 02       | 03      | 04  | 05        | 06 | 07        | 08    | 09  | 10 | 11   | 12       | 13 | 14 | 15 | 16 |
ISA       |00  |          |00        |         |30   |227777777  |14  |577777777  |120530 |1144 |U   |00401 |000000001 |0   |P   |~   |\   |
 GS       |RA  |227777777 |577777777 |20120530 |1144 |1          |X   |004010     |\      |     |    |      |          |    |    |    |    |
  ST      |820 |0001      |          |\        |     |           |    |           |       |     |    |      |          |    |    |    |    |
   Detail |    |          |          |         |     |           |    |           |       |     |    |      |          |    |    |    |    |
    ENT   |1   |          |          |         |\    |           |    |           |       |     |    |      |          |    |    |    |    |
    RMR   |IV  |7321239   |          |953.19   |     |           |\   |           |       |     |    |      |          |    |    |    |    |
    REF   |PO  |24305     |          |\        |     |           |    |           |       |     |    |      |          |    |    |    |    |
   Detail |    |          |          |         |     |           |    |           |       |     |    |      |          |    |    |    |    |
    RMR   |IV  |7321511   |          |7036.54  |     |           |\   |           |       |     |    |      |          |    |    |    |    |
    REF   |PO  |24333     |          |\        |     |           |    |           |       |     |    |      |          |    |    |    |    |
  SE      |11  |0001      |\         |         |     |           |    |           |       |     |    |      |          |    |    |    |    |
 GE       |1   |1         |\         |         |     |           |    |           |       |     |    |      |          |    |    |    |    |
IEA       |1   |000000001 |\         |         |     |           |    |           |       |     |    |      |          |    |    |    |    |
```

Rule print example
```
  DUMP RULE EPN STP 820

Segments & Rules Structure: | Usage   | Repeat Count | Description                                              |
ISA                         |REQUIRED |1             |INTERCHANGE CONTROL HEADER                                |
 GS                         |REQUIRED |1             |FUNCTIONAL GROUP HEADER                                   |
  ST                        |REQUIRED |1             |TRANSACTION SET HEADER                                    |
  BPR                       |REQUIRED |1             |Beginning Segment for Payment Order/Remittance Advice     |
  TRN                       |REQUIRED |1             |Trace                                                     |
  N1                        |REQUIRED |1             |Originator Name Identification                            |
  N1                        |REQUIRED |1             |Receiver Name Identification                              |
   Detail                   |OPTIONAL |>1            |                                                          |
    ENT                     |OPTIONAL |1             |Entity                                                    |
    RMR                     |REQUIRED |1             |Remittance Advice Accounts Receivable Open Item Reference |
    REF                     |OPTIONAL |>1            |Reference Identification                                  |
    DTM                     |OPTIONAL |1             |INDIVIDUAL COVERAGE PERIOD                                |
    ADX                     |OPTIONAL |1             |Adjustment                                                |
  SE                        |REQUIRED |1             |TRANSACTION SET TRAILER                                   |
 GE                         |REQUIRED |1             |FUNCTIONAL GROUP TRAILER                                  |
IEA                         |REQUIRED |1             |INTERCHANGE CONTROL TRAILER                               |
```

Rule print example with RequiredOnly option
```
  DUMP RULE EPN STP 820

Segments & Rules Structure: | Usage   | Repeat Count | Description                                          |
ISA                         |REQUIRED |1             |INTERCHANGE CONTROL HEADER                            |
 GS                         |REQUIRED |1             |FUNCTIONAL GROUP HEADER                               |
  ST                        |REQUIRED |1             |TRANSACTION SET HEADER                                |
  BPR                       |REQUIRED |1             |Beginning Segment for Payment Order/Remittance Advice |
  TRN                       |REQUIRED |1             |Trace                                                 |
  N1                        |REQUIRED |1             |Originator Name Identification                        |
  N1                        |REQUIRED |1             |Receiver Name Identification                          |
  SE                        |REQUIRED |1             |TRANSACTION SET TRAILER                               |
 GE                         |REQUIRED |1             |FUNCTIONAL GROUP TRAILER                              |
IEA                         |REQUIRED |1             |INTERCHANGE CONTROL TRAILER                           |
```


Checkout the [Go methods available](https://pkg.go.dev/github.com/moov-io/x12) for full details.

## Getting help

channel | info
------- | -------
[Project Documentation](https://github.com/moov-io/x12) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/x12/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`x12`) to have an interactive discussion about the development of the project.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! Checkout our [issues for first time contributors](https://github.com/moov-io/x12/contribute) for something to help out with.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
