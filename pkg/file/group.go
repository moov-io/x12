// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"errors"
	"strconv"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/moov-io/x12/pkg/util"
)

func NewGroup(rule *rules.GroupRule) *FunctionalGroup {
	newTransaction := FunctionalGroup{rule: rule}

	return &newTransaction
}

type FunctionalGroup struct {
	GS              segments.GS `json:"GS" xml:"GS"`
	TransactionSets []TransactionSet
	GE              *segments.GE `json:"GE,omitempty" xml:"GE,omitempty"`

	rule *rules.GroupRule
}

func (r *FunctionalGroup) GetTransactionControlNumbers() []string {
	var numbers []string
	for _, t := range r.TransactionSets {
		numbers = append(numbers, t.GetTransactionControlNumber())
	}
	return numbers
}

func (r *FunctionalGroup) GetGroupControlNumber() string {
	return r.GS.GroupControlNumber
}

func (r *FunctionalGroup) Validate(groupRule *rules.GroupRule) error {
	if groupRule == nil && r.rule != nil {
		groupRule = r.rule
	}

	var err error

	// Validating GS Segment
	gsRule := groupRule.GS
	{
		err = r.GS.Validate(&gsRule.Elements)
		if err != nil {
			util.AppendErrorStack(err, util.GetStructName(r))
			return err
		}
	}

	// Validating transaction sets
	{
		if len(r.TransactionSets) == 0 {
			return util.NewFindRuleError(util.GetStructName(r))
		}

		for index := 0; index < len(r.TransactionSets); index++ {
			set := r.TransactionSets[index]
			if err = set.Validate(&groupRule.Trans); err != nil {
				util.AppendErrorStack(err, util.GetStructName(r))
				return err
			}
		}

	}

	// Validating GE Segment
	geRule := groupRule.GE
	if r.GE != nil {
		err = r.GE.Validate(&geRule.Elements)
		if err != nil && rules.IsMaskRequired(geRule.Mask) {
			return err
		}
	}

	// Validating Group
	if r.GE != nil { // compare control set number
		if r.GE.GroupControlNumber != r.GS.GroupControlNumber {
			return errors.New("has invalid group control number")
		}

		// compare number of transaction set
		if v, conErr := strconv.ParseInt(r.GE.NumberOfTransactionSet, 10, 32); conErr == nil {
			if v != int64(len(r.TransactionSets)) {
				return errors.New("has invalid number of transaction set")
			}
		}
	}

	return nil
}

func (r *FunctionalGroup) Parse(data string, args ...string) (int, error) {
	if r.rule == nil {
		return 0, util.NewFindRuleError(util.GetStructName(r))
	}

	var size, read int
	var err error

	// Parsing GS Segment
	gsRule := r.rule.GS
	{
		r.GS.SetRule(&gsRule.Elements)
		size, err = r.GS.Parse(data[read:], args...)
		if err != nil {
			util.AppendErrorSegmentLine(err, data[read:], args...)
			return 0, err
		} else {
			read += size
		}
	}

	// Parsing Transaction sets
	trRule := r.rule.Trans
	for err == nil {
		newTrans := NewTransactionSet(&trRule)
		size, err = newTrans.Parse(data[read:], args...)
		if err == nil {
			read += size
			r.TransactionSets = append(r.TransactionSets, *newTrans)
		} else {
			line := data[read:]
			if len(r.TransactionSets) == 0 && (len(line) > 2 && line[0:2] == "ST") {
				util.AppendErrorStack(err, util.GetStructName(r))
				return 0, err
			}
		}
	}

	// Parsing GE Segment
	geRule := r.rule.GE
	if geRule.Name == "GE" {
		newGE := segments.NewGE(&geRule.Elements)
		size, err = newGE.Parse(data[read:], args...)
		if err != nil && rules.IsMaskRequired(geRule.Mask) {
			util.AppendErrorSegmentLine(err, data[read:], args...)
			return 0, err
		} else if err == nil {
			read += size
			if s, ok := newGE.(*segments.GE); ok {
				r.GE = s
			}
		}
	}

	return read, nil
}

func (r FunctionalGroup) String(args ...string) string {
	var buf bytes.Buffer

	buf.WriteString(r.GS.String(args...))
	for index := range r.TransactionSets {
		buf.WriteString(r.TransactionSets[index].String(args...))
	}
	if r.GE != nil {
		buf.WriteString(r.GE.String(args...))
	}

	return buf.String()
}

func (r FunctionalGroup) DumpStructInfo(level int) []util.ElementInfo {
	var selfDumps []util.ElementInfo

	dump := util.DumpStructInfo(r.GS, level)
	selfDumps = append(selfDumps, dump)

	for _, t := range r.TransactionSets {
		dumps := t.DumpStructInfo(level + 1)
		selfDumps = append(selfDumps, dumps...)
	}

	if r.GE != nil {
		dump = util.DumpStructInfo(r.GE, level)
		selfDumps = append(selfDumps, dump)
	}

	return selfDumps
}
