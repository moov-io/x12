// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x12_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/moov-io/x12/pkg/file"
	"github.com/moov-io/x12/pkg/loops"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/moov-io/x12/rules/rule_5010_835"
	"github.com/moov-io/x12/rules/rule_stp_820"
)

// This example reads an 835 file, validates it, and reads its segments.
func Example_parseAndValidate() {
	reader, err := os.Open(filepath.Join("examples", "example_5010_835", "sample2.txt"))
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	f := file.NewFile(&rule_5010_835.InterchangeRule)
	if err := f.Parse(file.NewScanner(reader)); err != nil {
		panic(err)
	}
	if err := f.Validate(); err != nil {
		panic(err)
	}

	for _, interchange := range f.Interchanges {
		for _, group := range interchange.FunctionalGroups {
			fmt.Println("GS08:", group.GS.GetFieldByIndex("08"))

			for _, txn := range group.TransactionSets {
				fmt.Println("ST01:", txn.ST.GetFieldByIndex("01"))

				// The first segments of an 835 are BPR and TRN
				for _, seg := range txn.Composite.Loop.Segments[:2] {
					fmt.Println(strings.TrimSuffix(seg.String(), "~"))
				}
			}
		}
	}
	// Output:
	// GS08: 005010X221A1
	// ST01: 835
	// BPR*I*4000.00*C*ACH*CCP*01*123456789*DA*0000000000*1234567890**01*123456789*DA*12345678*20990080
	// TRN*1*822070000341871*1066033492
}

// This example builds an STP 820 interchange, validates it, and writes it.
func Example_create() {
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
	// Output:
	// ISA*00* *00* *30*227777777 *14*577777777 *120530*1144*U*00401*000000001*0*P*~\
	// GS*RA*227777777*577777777*20120530*1144*1*X*004010STP820\
	// ST*820*0001\
	// BPR*C*7989.73*C*ACH*CTX*****1657777777**01*148529553*DA*92283334*20120531\
	// TRN*1*12053011440000192\
	// N1*PR*YOUR COMPANY*91*227777777\
	// N1*PE*WALMART\
	// SE*6*0001\
	// GE*1*1\
	// IEA*1*000000001\
}
