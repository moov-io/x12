// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"
	"strings"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

func NewCLP(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := CLP{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type CLP struct {
	ClaimSubmitterIdentifier  string `index:"01" json:"01" xml:"01"`
	ClaimStatusCode           string `index:"02" json:"02" xml:"02"`
	MonetaryAmount1           string `index:"03" json:"03" xml:"03"`
	MonetaryAmount2           string `index:"04" json:"04" xml:"04"`
	MonetaryAmount3           string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	ClaimFilingIndicatorCode  string `index:"06" json:"06" xml:"06"`
	ReferenceIdentification   string `index:"07" json:"07" xml:"07"`
	FacilityCodeValue         string `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	ClaimFrequencyTypeCode    string `index:"09" json:"09,omitempty" xml:"09,omitempty"`
	PatientStatusCode         string `index:"10" json:"10,omitempty" xml:"10,omitempty"`
	DiagnosisRelatedGroupCode string `index:"11" json:"11,omitempty" xml:"11,omitempty"`
	Quantity                  string `index:"12" json:"12,omitempty" xml:"12,omitempty"`
	Percent                   string `index:"13" json:"13,omitempty" xml:"13,omitempty"`

	Element
}

func (r CLP) defaultMask(index int) string {
	if index < 5 || index == 6 || index == 7 {
		return rules.MASK_REQUIRED
	}
	return rules.MASK_OPTIONAL
}

func (r CLP) fieldCount() int {
	return 13
}

func (r CLP) Name() string {
	return "CLP"
}

func (r *CLP) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r CLP) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *CLP) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("clp's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *CLP) Parse(data string, args ...string) (int, error) {
	var size int
	name := strings.ToLower(r.Name())
	read, line, err := r.VerifyCode(data, name, args...)
	if err != nil {
		return 0, err
	}

	for i := 1; i <= r.fieldCount(); i++ {
		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r CLP) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
