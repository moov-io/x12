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

func NewSBR(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := SBR{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type SBR struct {
	PayerResponsibilitySequenceCode string `index:"01" json:"01" xml:"01"`
	IndividualRelationshipCode      string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	ReferenceIdentification         string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	GroupName                       string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	InsuranceTypeCode               string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	CoordinationOfBenefitsCode      string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	ConditionOrResponseCode         string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	EmploymentStatusCode            string `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	ClaimFilingIndicator            string `index:"09" json:"09,omitempty" xml:"09,omitempty"`

	Element
}

func (r SBR) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index > 1 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r SBR) fieldCount() int {
	return 9
}

func (r SBR) Name() string {
	return "SBR"
}

func (r *SBR) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r SBR) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *SBR) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("sbr's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *SBR) Parse(data string, args ...string) (int, error) {
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

func (r SBR) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
