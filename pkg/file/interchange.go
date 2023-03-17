// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
)

func NewInterchange(rule *rules.InterChangeRule) *Interchange {

	newChange := Interchange{rule: rule}

	return &newChange
}

type Interchange struct {
	ISA              segments.ISA `json:"ISA" xml:"ISA"`
	FunctionalGroups []FunctionalGroup
	IEA              *segments.IEA `json:"IEA,omitempty" xml:"IEA,omitempty"`

	rule *rules.InterChangeRule
}

func (r *Interchange) Validate(changeRule *rules.InterChangeRule) error {

	if changeRule == nil && r.rule != nil {
		changeRule = r.rule
	}

	var err error

	// Validating ISA Segment
	isaRule := changeRule.ISA
	{
		if isaRule.Name != "ISA" {
			return errors.New("invalid isa rule")
		}

		err = r.ISA.Validate(&isaRule.Elements)
		if err != nil {
			return errors.New("unable to validate isa segment")
		}
	}

	// Validating groups
	{
		for index := 0; index < len(r.FunctionalGroups); index++ {
			group := r.FunctionalGroups[index]
			if err = group.Validate(&changeRule.Group); err != nil {
				return fmt.Errorf("group(%02d) is not invalid", index)
			}
		}

	}

	// Validating IEA Segment
	ieaRule := changeRule.IEA
	if rules.IsMaskRequired(ieaRule.Mask) && r.IEA == nil {
		return errors.New("iea segment is required segment")
	}

	if r.IEA != nil {
		if ieaRule.Name != "IEA" {
			return errors.New("invalid iea rule")
		}

		err = r.IEA.Validate(&ieaRule.Elements)
		if err != nil && rules.IsMaskRequired(ieaRule.Mask) {
			return errors.New("unable to validate iea segment")
		}
	}

	return nil
}

func (r *Interchange) Parse(data string, args ...string) (int, error) {

	if r.rule == nil {
		return 0, errors.New("please specify rules for this group")
	}

	var size, read int
	var err error

	line := data[read:]

	// Parsing ISA Segment
	isaRule := r.rule.ISA
	{
		if isaRule.Name != "ISA" {
			return 0, errors.New("invalid isa rule")
		}

		r.ISA.SetRule(&isaRule.Elements)
		size, err = r.ISA.Parse(line, args...)
		if err != nil {
			return 0, errors.New("unable to parse isa segment")
		} else {
			read += size
			line = data[read:]
		}
	}

	// Parsing groups
	grRule := r.rule.Group
	for err == nil {
		newTrans := NewGroup(&grRule)
		size, err = newTrans.Parse(line, args...)
		if err == nil {
			read += size
			line = data[read:]
			r.FunctionalGroups = append(r.FunctionalGroups, *newTrans)
		} else {
			break
		}
	}

	// Parsing IEA Segment
	ieaRule := r.rule.IEA
	if ieaRule.Name == "IEA" {
		newIEA := segments.NewIEA(&ieaRule.Elements)
		size, err = newIEA.Parse(line, args...)
		if err != nil && rules.IsMaskRequired(ieaRule.Mask) {
			return 0, errors.New("unable to parse iea segment")
		} else if err == nil {
			read += size
			if s, ok := newIEA.(*segments.IEA); ok {
				r.IEA = s
			}
		}
	}

	return read, nil
}

func (r Interchange) String(args ...string) string {
	var buf bytes.Buffer

	buf.WriteString(r.ISA.String(args...))
	for index := range r.FunctionalGroups {
		buf.WriteString(r.FunctionalGroups[index].String(args...))
	}
	if r.IEA != nil {
		buf.WriteString(r.IEA.String(args...))
	}

	return buf.String()
}
