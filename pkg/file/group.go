// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"errors"
	"fmt"
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

func (r *FunctionalGroup) Validate(groupRule *rules.GroupRule) error {

	if groupRule == nil && r.rule != nil {
		groupRule = r.rule
	}

	var err error

	// Validating GS Segment
	gsRule := groupRule.GS
	{
		if gsRule.Name != "GS" {
			return errors.New("invalid gs rule")
		}

		err = r.GS.Validate(&gsRule.Elements)
		if err != nil {
			return errors.New("unable to validate gs segment")
		}
	}

	// Validating transaction sets
	{
		for index := 0; index < len(r.TransactionSets); index++ {
			set := r.TransactionSets[index]
			if err = set.Validate(&groupRule.Trans); err != nil {
				return fmt.Errorf("transaction set(%02d) is not invalid", index)
			}
		}

	}

	// Validating GE Segment
	geRule := groupRule.GE
	if rules.IsMaskRequired(geRule.Mask) && r.GE == nil {
		return errors.New("ge segment is required segment")
	}

	if r.GE != nil {
		if geRule.Name != "GE" {
			return errors.New("invalid ge rule")
		}

		err = r.GE.Validate(&geRule.Elements)
		if err != nil && rules.IsMaskRequired(geRule.Mask) {
			return errors.New("unable to validate ge segment")
		}
	}

	// Validating Group
	if r.GE != nil {

		// compare control set number
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
		return 0, errors.New("please specify rules for this group")
	}

	var size, read int
	var err error

	line := data[read:]

	// Parsing GS Segment
	gsRule := r.rule.GS
	{
		if gsRule.Name != "GS" {
			return 0, errors.New("invalid gs rule")
		}

		r.GS.SetRule(&gsRule.Elements)
		size, err = r.GS.Parse(line, args...)
		if err != nil {
			return 0, errors.New("unable to parse gs segment")
		} else {
			read += size
			line = data[read:]
		}
	}

	// Parsing Transaction sets
	trRule := r.rule.Trans
	for err == nil {
		newTrans := NewTransactionSet(&trRule)
		size, err = newTrans.Parse(line, args...)
		if err == nil {
			read += size
			line = data[read:]
			r.TransactionSets = append(r.TransactionSets, *newTrans)
		} else {
			if len(r.TransactionSets) == 0 && (len(line) > 2 && line[0:2] == "ST") {
				return 0, errors.New("unable to parse transaction set")
			}
		}
	}

	// Parsing GE Segment
	geRule := r.rule.GE
	if geRule.Name == "GE" {
		newGE := segments.NewGE(&geRule.Elements)
		size, err = newGE.Parse(line, args...)
		if err != nil && rules.IsMaskRequired(geRule.Mask) {
			return 0, errors.New("unable to parse ge segment")
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
