// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"errors"
	"github.com/moov-io/x12/pkg/loops"
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/moov-io/x12/pkg/util"
	"strconv"
)

func NewTransactionSet(rule *rules.TransactionRule) *TransactionSet {
	newTransaction := TransactionSet{rule: rule}

	return &newTransaction
}

type TransactionSet struct {
	ST        segments.ST         `json:"ST" xml:"ST"`
	Composite loops.CompositeLoop `json:"composite" xml:"composite"`
	SE        *segments.SE        `json:"SE,omitempty" xml:"SE,omitempty"`

	rule *rules.TransactionRule
}

func (r *TransactionSet) GetTransactionControlNumber() string {
	return r.ST.TransactionSetControlNumber
}

func (r *TransactionSet) Validate(transRule *rules.TransactionRule) error {
	if transRule == nil && r.rule != nil {
		transRule = r.rule
	}

	var err error

	// Validating ST Segment
	stRule := transRule.ST
	{
		err = r.ST.Validate(&stRule.Elements)
		if err != nil && rules.IsMaskRequired(stRule.Mask) {
			util.AppendErrorStack(err, util.GetStructName(r))
			return err
		}
	}

	// Validating Composite
	err = r.Composite.Validate(&transRule.Composite)
	if err != nil {
		util.AppendErrorStack(err, util.GetStructName(r))
		return err
	}

	// Validating SE Segment
	seRule := transRule.SE
	{
		err = r.SE.Validate(&seRule.Elements)
		if err != nil && rules.IsMaskRequired(seRule.Mask) {
			util.AppendErrorStack(err, util.GetStructName(r))
			return err
		}
	}

	// Validating Transaction Set
	if r.SE != nil { // compare control set number
		if r.SE.TransactionSetControlNumber != r.ST.TransactionSetControlNumber {
			return errors.New("has invalid transaction set control number")
		}

		// getting segments count
		segmentCnt := 2

		// segments
		segmentCnt += len(r.Composite.GetSegments())

		// compare number of segments
		if v, conErr := strconv.ParseInt(r.SE.NumberOfSegments, 10, 32); conErr == nil {
			if v != int64(segmentCnt) {
				return errors.New("has invalid number of segments")
			}
		}
	}

	return nil
}

func (r *TransactionSet) Parse(data string, args ...string) (int, error) {
	if r.rule == nil {
		return 0, util.NewSpecifiedRuleError(util.GetStructName(r))
	}

	var size, read int
	var err error

	// Parsing ST Segment
	stRule := r.rule.ST
	{
		r.ST.SetRule(&stRule.Elements)
		size, err = r.ST.Parse(data[read:], args...)
		read += size
		if err != nil {
			util.AppendErrorStack(err, util.GetStructName(r))
			return read, err
		}
	}

	// Parsing Composite
	{
		r.Composite.SetRule(&r.rule.Composite)
		size, err = r.Composite.Parse(data[read:], args...)
		read += size
		if err != nil {
			util.AppendErrorStack(err, util.GetStructName(r))
			return read, err
		}
	}

	// Parsing SE Segment
	seRule := r.rule.SE
	if seRule.Name == "SE" {
		newSE := segments.NewSE(&seRule.Elements)
		size, err = newSE.Parse(data[read:], args...)
		if err != nil && rules.IsMaskRequired(seRule.Mask) {
			util.AppendErrorStack(err, util.GetStructName(r))
			return read, err
		} else if err == nil {
			read += size
			if s, ok := newSE.(*segments.SE); ok {
				r.SE = s
			}
		}
	}

	return read, nil
}

func (r TransactionSet) String(args ...string) string {
	var buf bytes.Buffer

	buf.WriteString(r.ST.String(args...))

	segList := r.Composite.GetSegments()
	for index := range segList {
		buf.WriteString(segList[index].String(args...))
	}
	if r.SE != nil {
		buf.WriteString(r.SE.String(args...))
	}

	return buf.String()
}

func (r TransactionSet) DumpStructInfo(level int) []util.ElementInfo {
	var selfDumps []util.ElementInfo

	dump := util.DumpStructInfo(r.ST, level)
	selfDumps = append(selfDumps, dump)

	dumps := r.Composite.DumpStructInfo(level)
	selfDumps = append(selfDumps, dumps...)

	if r.SE != nil {
		dump = util.DumpStructInfo(r.SE, level)
		selfDumps = append(selfDumps, dump)
	}

	return selfDumps
}
