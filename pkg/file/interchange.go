// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/moov-io/x12/pkg/util"
)

func NewInterchange(rule *rules.InterchangeRule, args ...string) *Interchange {
	newChange := Interchange{rule: rule}

	if len(args) > 0 && args[0] != "" {
		newChange.terminator = args[0]
	}

	return &newChange
}

type Interchange struct {
	ISA              segments.ISA `json:"ISA" xml:"ISA"`
	FunctionalGroups []FunctionalGroup
	IEA              *segments.IEA `json:"IEA,omitempty" xml:"IEA,omitempty"`

	terminator string

	rule *rules.InterchangeRule
}

func (r Interchange) getTerminator() string {
	if r.terminator != "" {
		return r.terminator
	}

	return util.SegmentTerminator
}

func (r *Interchange) GetTransactionControlNumbers() []string {
	var numbers []string
	for _, t := range r.FunctionalGroups {
		numbers = append(numbers, t.GetTransactionControlNumbers()...)
	}
	return numbers
}

func (r *Interchange) GetGroupControlNumbers() []string {
	var numbers []string
	for _, t := range r.FunctionalGroups {
		numbers = append(numbers, t.GetGroupControlNumber())
	}
	return numbers
}

func (r *Interchange) Validate(validateRule *rules.InterchangeRule) error {
	changeRule := r.rule
	if validateRule != nil {
		changeRule = validateRule
	}

	if changeRule == nil {
		return util.NewFindRuleError(util.GetStructName(r))
	}

	var err error

	// Validating ISA Segment
	isaRule := changeRule.ISA
	{
		err = r.ISA.Validate(&isaRule.Elements)
		if err != nil {
			return util.UpdateErrorReason(err)
		}
	}

	// Validating groups
	{
		if len(r.FunctionalGroups) == 0 {
			return util.NewFindElementError("group")
		}

		for index := 0; index < len(r.FunctionalGroups); index++ {
			group := r.FunctionalGroups[index]
			if err = group.Validate(&changeRule.Group); err != nil {
				return util.UpdateErrorReason(err)
			}
		}

	}

	ieaRule := changeRule.IEA
	if r.IEA != nil {
		err = r.IEA.Validate(&ieaRule.Elements)
		if err != nil && rules.IsMaskRequired(ieaRule.Mask) {
			return util.UpdateErrorReason(err)
		}
	}

	// Validating
	if r.IEA != nil { // compare control set number
		if r.IEA.InterchangeControlNumber != r.ISA.InterchangeControlNumber {
			return errors.New("has invalid interchange control number")
		}

		// compare number of groups
		if v, conErr := strconv.ParseInt(r.IEA.NumberOfFunctionalGroups, 10, 32); conErr == nil {
			if v != int64(len(r.FunctionalGroups)) {
				return errors.New("has invalid number of functional groups")
			}
		}
	}

	// Validating transaction control numbers
	if exist, number := util.GetDuplicateControlNumber(r.GetTransactionControlNumbers()); exist {
		return fmt.Errorf("transaction control number(%s) should be unique per interchange", number)
	}

	// Validating group control numbers
	if exist, number := util.GetDuplicateControlNumber(r.GetGroupControlNumbers()); exist {
		return fmt.Errorf("group control number(%s) should be unique per interchange", number)
	}

	return nil
}

func (r *Interchange) Parse(data string) (int, error) {
	if r.rule == nil {
		return 0, util.NewFindRuleError(util.GetStructName(r))
	}

	var size, read int
	var err error

	// Parsing ISA Segment
	isaRule := r.rule.ISA
	{
		r.ISA.SetRule(&isaRule.Elements)
		size, err = r.ISA.Parse(data[read:], r.getTerminator())
		if err != nil {
			util.AppendErrorSegmentLine(err, data[read:], r.getTerminator())
			return 0, util.UpdateErrorReason(err)
		} else {
			read += size
		}
	}

	// Parsing groups
	grRule := r.rule.Group
	for err == nil {
		newTrans := NewGroup(&grRule)
		size, err = newTrans.Parse(data[read:], r.getTerminator(), r.ISA.ComponentElementSeparator)
		if err == nil {
			read += size
			r.FunctionalGroups = append(r.FunctionalGroups, *newTrans)
		} else {
			line := data[read:]
			if len(r.FunctionalGroups) == 0 && (len(line) > 2 && line[0:2] == "GS") {
				return 0, util.UpdateErrorReason(err)
			}
		}
	}

	// Parsing IEA Segment
	ieaRule := r.rule.IEA
	if ieaRule.Name == "IEA" {
		newIEA := segments.NewIEA(&ieaRule.Elements)
		size, err = newIEA.Parse(data[read:], r.getTerminator())
		if err != nil && rules.IsMaskRequired(ieaRule.Mask) {
			util.AppendErrorSegmentLine(err, data[read:], r.getTerminator())
			return 0, util.UpdateErrorReason(err)
		} else if err == nil {
			read += size
			if s, ok := newIEA.(*segments.IEA); ok {
				r.IEA = s
			}
		}
	} else if rules.IsMaskRequired(ieaRule.Mask) {
		return 0, util.NewFindSegmentError("iea")
	}

	return read, nil
}

func (r Interchange) String(args ...string) string {
	var buf bytes.Buffer

	if len(args) == 0 {
		args = append(args, r.getTerminator(), r.ISA.ComponentElementSeparator)
	} else if len(args) == 1 {
		args = append(args, r.ISA.ComponentElementSeparator)
	}

	buf.WriteString(r.ISA.String(args...))
	for index := range r.FunctionalGroups {
		buf.WriteString(r.FunctionalGroups[index].String(args...))
	}
	if r.IEA != nil {
		buf.WriteString(r.IEA.String(args...))
	}

	return buf.String()
}

//  The function will print human-readable format of edi file
//  User should be fund segments and loops with level, indexes are normal spec indexes
//
//	INDEX:       | 00           | 01               | 02                     | 03        | 04    | 05             | 06 | 07             | 08        | 09  | 10 | 11   | 12       | 13 | 14 | 15 | 16 | 17 | 18 |
//	ISA          |00            |                  |00                      |           |ZZ     |133052274       |ZZ  |311279999       |120419     |2125 |^   |00501 |000002120 |0   |P   |:   |~   |    |    |
//	 GS          |HC            |133052274         |311279999               |20120419   |212549 |2120            |X   |005010X224A2    |~          |     |    |      |          |    |    |    |    |    |    |
//	  ST         |837           |3456              |005010X224A2            |~          |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |
//	   1000A     |              |                  |                        |           |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |
//	    NM1      |41            |2                 |PREMIER BILLING SERVICE |           |       |                |    |46              |TGJ23      |~    |    |      |          |    |    |    |    |    |    |
//	    PER      |IC            |JERRY             |TE                      |7176149999 |       |                |    |                |~          |     |    |      |          |    |    |    |    |    |    |
//   ... ...
//	  SE         |31            |3456              |~                       |           |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |
//	 GE          |1             |2120              |~                       |           |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |
//	IEA          |1             |000002120         |~                       |           |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |

func (r Interchange) Print(w io.Writer) {
	padChar := byte(' ')
	padding := 1
	level := 0

	var selfDumps []util.ElementInfo

	dump := util.DumpStructInfo(r.ISA, level)
	selfDumps = append(selfDumps, dump)

	for _, g := range r.FunctionalGroups {
		dumps := g.DumpStructInfo(level + 1)
		selfDumps = append(selfDumps, dumps...)
	}

	if r.IEA != nil {
		dump = util.DumpStructInfo(r.IEA, level)
		selfDumps = append(selfDumps, dump)
	}

	maxItemCnt := 0
	for _, d := range selfDumps {
		if len(d.Items) > maxItemCnt {
			maxItemCnt = len(d.Items)
		}
	}

	tw := tabwriter.NewWriter(w, 0, 0, padding, padChar, tabwriter.Debug)

	// header line
	{
		printStr := "INDEX:\t"
		for i := 0; i < maxItemCnt+1; i++ {
			idx := fmt.Sprintf(" %02d", i)
			printStr = printStr + idx + "\t"
		}
		fmt.Fprintln(tw, printStr)
	}

	for _, d := range selfDumps {
		printStr := d.Name + "\t"
		for _, item := range d.Items {
			printStr = printStr + item + "\t"
		}

		if len(d.Items) > 0 {
			printStr = printStr + r.getTerminator() + "\t"
		} else {
			printStr = printStr + "\t"
		}

		for i := 0; i < maxItemCnt-len(d.Items); i++ {
			printStr = printStr + "\t"
		}

		if d.Level > 0 {
			printStr = strings.Repeat(string(padChar), d.Level) + printStr
		}

		fmt.Fprintln(tw, printStr)
	}

	tw.Flush()
}
