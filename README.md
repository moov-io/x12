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

[![Go Reference](https://pkg.go.dev/badge/github.com/moov-io/x12.svg)](https://pkg.go.dev/github.com/moov-io/x12)
[![Build Status](https://github.com/moov-io/x12/workflows/Go/badge.svg)](https://github.com/moov-io/x12/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/x12/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/x12)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/x12)](https://goreportcard.com/report/github.com/moov-io/x12)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/x12?label=project%20size)](https://github.com/moov-io/x12)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/x12/master/LICENSE)
[![Slack Channel](https://img.shields.io/badge/slack-moov--io-e01563?logo=slack)](https://slack.moov.io/)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/x12)](https://github.com/moov-io/x12)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/x12

moov-io/x12 is a Go library for parsing and generating Accredited Standards Committee X12 (ASC X12) documents. The goal is to support Electronic data interchange (EDI) Context Inspired Component Architecture (CICA) standards which support health care, insurance, transportation, finance, government, supply chains and many other industries.

## Project status

The x12 is actively being developed for use in production systems. Please star the project if you are interested in its progress. If you need other ASC X12 formats, find bugs, have comments, or suggestions we would appreciate an issue or pull request. Thanks!

## Go library

This project uses [Go Modules](https://go.dev/blog/using-go-modules) and Go v1.19 or newer. See [Golang's install instructions](https://golang.org/doc/install) for help in setting up Go. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/x12/releases) as well. We highly recommend you use a tagged release for production.

### Installation

```
go get github.com/moov-io/x12
```

### Quick start: parse and validate a file

Create a file with the rule for your transaction set, parse it with a scanner, then call `Validate()`. This program reads an 835 file and prints the first segments of each transaction set:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moov-io/x12/pkg/file"
	"github.com/moov-io/x12/rules/rule_5010_835"
)

func main() {
	reader, err := os.Open("examples/example_5010_835/sample2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	f := file.NewFile(&rule_5010_835.InterchangeRule)
	if err := f.Parse(file.NewScanner(reader)); err != nil {
		log.Fatal(err)
	}
	if err := f.Validate(); err != nil {
		log.Fatal(err)
	}

	for _, interchange := range f.Interchanges {
		for _, group := range interchange.FunctionalGroups {
			fmt.Println("GS08:", group.GS.GetFieldByIndex("08")) // 005010X221A1

			for _, txn := range group.TransactionSets {
				fmt.Println("ST01:", txn.ST.GetFieldByIndex("01")) // 835

				// Segments of the transaction set (BPR, TRN, ...) are in txn.Composite.Loop.Segments.
				// Nested loops are in txn.Composite.SubLoops.
				for _, seg := range txn.Composite.Loop.Segments {
					fmt.Println(seg.String())
				}
			}
		}
	}
}
```

The segment terminator is `~` by default. If your files use a different terminator, pass it to `file.NewFile` and `file.NewScanner`, for example `file.NewFile(&rule, "\\")` and `file.NewScanner(reader, "\\")`.

A tested version of this program is `Example_parseAndValidate` in [example_test.go](./example_test.go).

### Examples

The [examples](./examples) folder has a program for each of these transaction sets. Each program parses and validates a sample file, writes the file again, and prints its structure.

| Example | Rule |
|---------|------|
| [example_4010_810](./examples/example_4010_810) | 4010 810 Invoice |
| [example_5010_835](./examples/example_5010_835) | 5010 835 Health Care Claim Payment/Advice |
| [example_5010_837d](./examples/example_5010_837d) | 5010 837D Health Care Claim: Dental |
| [example_5010_837p](./examples/example_5010_837p) | 5010 837P Health Care Claim: Professional |
| [example_stp_820](./examples/example_stp_820) | STP 820 Payment Order/Remittance Advice |

Run an example from the repository root:

```
go run ./examples/example_5010_835
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
* X12 - Supply Chain / 810 Invoice

To support these transaction sets, x12 project used following data unit
* Interchange
  * Interchange includes control header(ISA), control trailer(IEA), functional groups. i.e., segments and groups
* Group
  * Functional group includes group header(GS), group trailer(GE), transaction sets. i.e., segments and transaction sets
* Transaction Set
  * Transaction set includes transaction set header(ST), transaction set trailer(SE), groups, and segments. i.e., segments and loops
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

* [4010 810](./rules/rule_4010_810/rule.go) - ASC X12F 4010 810
* [4010 820](./rules/rule_4010_820/rule.go) - ASC X12F 4010 820 (004010X061A1)
* [4010 835](./rules/rule_4010_835/rule.go) - ASC X12N 4010 835 (004010X091A1)
* [4010 837I](./rules/rule_4010_837I/rule.go) - ASC X12N 4010 837I (004010X096A1)
* [5010 835](./rules/rule_5010_835/rule.go) - ASC X12N 5010 835 (005010X221A1)
* [5010 837d](./rules/rule_5010_837d/rule.go) - ASC X12N 5010 837 (005010X224A2)
* [5010 837p](./rules/rule_5010_837p/rule.go) - ASC X12N 5010 837 (005010X222A1)
* [STP 820](./rules/rule_stp_820/rule.go) - ASC X12F STP 820 (004010STP820)

We can specify rule using some kinds of rule structs.

#### InterchangeRule
```go
type InterchangeRule struct {
	Name  string
	ISA   SegmentRule
	IEA   SegmentRule
	Group GroupRule
}
```
Interchange rule defined details of control header.

#### GroupRule
```go
type GroupRule struct {
	GS    SegmentRule
	GE    SegmentRule
	Trans TransactionRule
}
```
Group rule defined details of group header.

#### TransactionRule
```go
type TransactionRule struct {
	ST        SegmentRule
	SE        SegmentRule
	Composite LoopRule
}
```
Transaction rule defined details of transaction set.
Composite is the loop between ST and SE. Its segments and nested loops (Composite) are ordered by the index of the rule map.

#### LoopSetRule
```go
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
```go
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

#### ElementSetRule
```go
type ElementSetRule map[string]ElementRule

type ElementRule struct {
	Mask         string
	Name         string
	AcceptRegex  string
	AcceptValues []string
	Composite    ElementSetRule
}
```
Mask is same as loop's mask.
Composite is ordered nested fields.
Accept values is to specify available values that can use the field.
Accept regex is a regular expression that the value of the field must match.

Element's mask is a bit different from segment and loop.
Each segment has a real struct such as IEA, ISA, GS, etc.
Field of segment is omitted field or required field (check out with json, xml tag).
If the mask option of element rule is omitted, will use original field's property (omitted field or required field).

#### Rule example

This is the rule for STP 820 files, from [rules/rule_stp_820/rule.go](./rules/rule_stp_820/rule.go):

```go
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
	Composite: rules.LoopRule{
		Name: "Transaction Loop",
		Mask: rules.MASK_REQUIRED,
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
				Name:        "N1",
				Description: "Originator Name Identification",
				Mask:        rules.MASK_REQUIRED,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PR"}},
					"02": {Mask: rules.MASK_OPTIONAL},
					"03": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"91"}},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
				},
			},
			3: rules.SegmentRule{
				Name:        "N1",
				Description: "Receiver Name Identification",
				Mask:        rules.MASK_REQUIRED,
				Elements: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"PE"}},
					"02": {Mask: rules.MASK_OPTIONAL},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
				},
			},
		},
		Composite: map[int]rules.LoopRule{
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
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"RA"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED},
			"06": {Mask: rules.MASK_REQUIRED},
			"07": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"X"}},
			"08": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"004010STP820"}},
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
	Name: "EPN STP 820",
	ISA: rules.SegmentRule{
		Name:        "ISA",
		Description: "INTERCHANGE CONTROL HEADER",
		Mask:        rules.MASK_REQUIRED,
		Elements: rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00"}},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00"}},
			"04": {Mask: rules.MASK_REQUIRED},
			"05": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "14", "17", "30"}},
			"06": {Mask: rules.MASK_REQUIRED},
			"07": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"01", "14", "17", "30"}},
			"08": {Mask: rules.MASK_REQUIRED},
			"09": {Mask: rules.MASK_REQUIRED},
			"10": {Mask: rules.MASK_REQUIRED},
			"11": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"U"}},
			"12": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"00401"}},
			"13": {Mask: rules.MASK_REQUIRED},
			"14": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"0"}},
			"15": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"P"}},
			"16": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"~"}},
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
```
In above rule, transaction set should have ST, a transaction loop with BPR, TRN, N1, N1 and optional Detail loops, and SE.

## Usage

User should create interchange, group, transaction, segments. This program builds an STP 820 interchange, validates it, and writes it:

```go
package main

import (
	"fmt"
	"strings"

	"github.com/moov-io/x12/pkg/file"
	"github.com/moov-io/x12/pkg/loops"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/moov-io/x12/rules/rule_stp_820"
)

func main() {
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
	gs.SetFieldByIndex("08", "004010STP820")

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

	interchange := file.Interchange{
		ISA: isa,
		IEA: &iea,
		FunctionalGroups: []file.FunctionalGroup{
			{
				GS: gs,
				GE: &ge,
				TransactionSets: []file.TransactionSet{
					{
						ST: st,
						SE: &se,
						Composite: loops.CompositeLoop{
							Loop: loops.Loop{
								Segments: []segments.SegmentInterface{&bpr, &trn, &n11, &n12},
							},
						},
					},
				},
			},
		},
	}

	if err := interchange.Validate(&rule_stp_820.InterchangeRule); err != nil {
		panic(err)
	}

	segmentTerminator := "\\"
	fmt.Println(strings.ReplaceAll(interchange.String(segmentTerminator), segmentTerminator, segmentTerminator+"\n"))
}
```

A tested version of this program is `Example_create` in [example_test.go](./example_test.go).

We need to know hierarchy of file, interchange or rule when editing new interchange.

Please use Print() function to verify edited interchange or target rule structure

```go
func (f File) Print(w io.Writer)

func (r Interchange) Print(w io.Writer)

func (r InterchangeRule) Print(w io.Writer, isRequiredOnly bool)
```

File print example
```
  DUMP EDI FILE WITH EPN STP 820

INDEX:             | 00 | 01               | 02       | 03       | 04  | 05        | 06 | 07        | 08    | 09        | 10 | 11   | 12       | 13 | 14      | 15      | 16 | 17 | 18 | 19 | 20 | 21 |
ISA                |00  |                  |00        |          |30   |227777777  |14  |577777777  |120530 |1144       |U   |00401 |000000001 |0   |P        |~        |\   |    |    |    |    |    |
 GS                |RA  |227777777         |577777777 |20120530  |1144 |1          |X   |004010     |\      |           |    |      |          |    |         |         |    |    |    |    |    |    |
  ST               |820 |0001              |          |\         |     |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
  Transaction Loop |    |                  |          |          |     |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
   BPR             |C   |7989.73           |C         |ACH       |CTX  |           |    |           |       |1657777777 |    |01    |148529553 |DA  |92283334 |20120531 |    |    |    |    |    |\   |
   TRN             |1   |12053011440000192 |          |          |\    |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
   N1              |PR  |YOUR COMPANY      |91        |227777777 |     |           |\   |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
   N1              |PE  |WALMART           |          |          |     |           |\   |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
   Detail          |    |                  |          |          |     |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
    ENT            |1   |                  |          |          |\    |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
    RMR            |IV  |7321239           |          |953.19    |     |           |\   |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
    REF            |PO  |24305             |          |\         |     |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
   Detail          |    |                  |          |          |     |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
    RMR            |IV  |7321511           |          |7036.54   |     |           |\   |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
    REF            |PO  |24333             |          |\         |     |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
  SE               |11  |0001              |\         |          |     |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
 GE                |1   |1                 |\         |          |     |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
IEA                |1   |000000001         |\         |          |     |           |    |           |       |           |    |      |          |    |         |         |    |    |    |    |    |    |
```

Rule print example
```
  DUMP RULE EPN STP 820

Segments & Rules Structure: | Usage   | Repeat Count | Description                                              |
ISA                         |REQUIRED |1             |INTERCHANGE CONTROL HEADER                                |
 GS                         |REQUIRED |1             |FUNCTIONAL GROUP HEADER                                   |
  ST                        |REQUIRED |1             |TRANSACTION SET HEADER                                    |
  Transaction Loop          |REQUIRED |1             |                                                          |
   BPR                      |REQUIRED |1             |Beginning Segment for Payment Order/Remittance Advice     |
   TRN                      |REQUIRED |1             |Trace                                                     |
   N1                       |REQUIRED |1             |Originator Name Identification                            |
   N1                       |REQUIRED |1             |Receiver Name Identification                              |
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
  Transaction Loop          |REQUIRED |1             |                                                      |
   BPR                      |REQUIRED |1             |Beginning Segment for Payment Order/Remittance Advice |
   TRN                      |REQUIRED |1             |Trace                                                 |
   N1                       |REQUIRED |1             |Originator Name Identification                        |
   N1                       |REQUIRED |1             |Receiver Name Identification                          |
  SE                        |REQUIRED |1             |TRANSACTION SET TRAILER                               |
 GE                         |REQUIRED |1             |FUNCTIONAL GROUP TRAILER                              |
IEA                         |REQUIRED |1             |INTERCHANGE CONTROL TRAILER                           |
```


Checkout the [Go methods available](https://pkg.go.dev/github.com/moov-io/x12) for full details.


### Error Messages
X12 project used 2 error styles for parsing logic. 
First style used when getting error on any segment inside composite loop or loop, second type used for normal segment
* reason, error line, error stack
```
  tds segment can't parse all input data, error segment line: 'TDS*139550*1*1*1*1~', error stack: 'functionalgroup->transactionset->summary->tds'
```

* reason, error line
```
  ge segment contains invalid code, error segment line: 'ST*810*0002~'
```

X12 project used 2 error styles for validating logic.
First style used when getting validate error on any segment inside composite loop or loop, second type used for normal segment
* reason, error stack
```
  segment(amt) don't accept specified rule(tds), please verify segment orders or has dirty segments as previous segment, error stack: 'functionalgroup->transactionset->summary->amt'
```

* reason
```
  segment(amt) don't accept specified rule(tds), please verify segment orders or has dirty segments as previous segment
```

## Further Reading

- [What makes EDI so hard?](https://www.stedi.com/blog/what-makes-edi-so-hard)
- EDI [inspector](https://www.stedi.com/edi/inspector) / [reference](https://www.stedi.com/edi) (Stedi)

## Getting help

channel | info
------- | -------
[Project Documentation](https://github.com/moov-io/x12) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/x12/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`x12`) to have an interactive discussion about the development of the project.

## Command line

The `x12` command line tool prints EDI files and rules in a human-readable format. Install it with:

```
go install github.com/moov-io/x12/cmd/x12@latest
```

Usage:

```
x12 describe_file [-rule <name>] [-validate] <files>
x12 describe_rule [-rule <name>] [-required]
x12 -version
```

`describe_file` parses and prints each file. With `-validate` it also validates each file. When `-rule` is not given, it uses the built-in rule that accepts the GS08 value of the file. It exits with a non-zero code when a file cannot be read, parsed, or validated.

`describe_rule` prints a rule. The default rule is `5010_837d`. With `-required` it prints only the required segments.

Supported rules: `4010_810`, `4010_820`, `4010_835`, `4010_837i`, `5010_835`, `5010_837d`, `5010_837p`, `stp_820`.

```
x12 describe_file -validate examples/example_5010_835/sample2.txt
x12 describe_rule -rule 5010_835 -required
```

## Contributing

Yes please! Please review our [Contributing guide](https://github.com/moov-io/ach/blob/master/CONTRIBUTING.md) and [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md) to get started! Checkout our [issues for first time contributors](https://github.com/moov-io/x12/contribute) for something to help out with.

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
