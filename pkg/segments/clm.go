// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"errors"
	"fmt"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

func NewCLM(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := CLM{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type CLM struct {
	ClaimSubmitterIdentifier    string                    `index:"01" json:"01" xml:"01"`
	MonetaryAmount              string                    `index:"02" json:"02" xml:"02"`
	ClaimFilingIndicatorCode    string                    `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	NonInstitutionClaimTypeCode string                    `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	HealthCareServiceLocation   HealthCareServiceLocation `index:"05" json:"05" xml:"05"`
	Code1                       string                    `index:"06" json:"06" xml:"06"`
	ProvideAcceptAssign         string                    `index:"07" json:"07" xml:"07"`
	Code2                       string                    `index:"08" json:"08" xml:"08"`
	ReleaseInformation          string                    `index:"09" json:"09" xml:"09"`
	PatientSignatureSourceCode  string                    `index:"10" json:"10,omitempty" xml:"10,omitempty"`
	RelatedCausesInformation    *RelatedCausesInformation `index:"11" json:"11,omitempty" xml:"11,omitempty"`
	Code3                       string                    `index:"12" json:"12,omitempty" xml:"12,omitempty"`
	Code4                       string                    `index:"13" json:"13,omitempty" xml:"13,omitempty"`
	Code5                       string                    `index:"14" json:"14,omitempty" xml:"14,omitempty"`
	Code6                       string                    `index:"15" json:"15,omitempty" xml:"15,omitempty"`
	Code7                       string                    `index:"16" json:"16,omitempty" xml:"16,omitempty"`
	Code8                       string                    `index:"17" json:"17,omitempty" xml:"17,omitempty"`
	Code9                       string                    `index:"18" json:"18,omitempty" xml:"18,omitempty"`
	Code10                      string                    `index:"19" json:"19,omitempty" xml:"19,omitempty"`
	Code11                      string                    `index:"20" json:"20,omitempty" xml:"20,omitempty"`

	Element
}

func (r CLM) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index > 9 || (index == 3 || index == 4) {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r CLM) fieldCount() int {
	return 20
}

func (r CLM) Name() string {
	return "CLM"
}

func (r *CLM) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r CLM) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *CLM) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var err error
		idx := fmt.Sprintf("%02d", i)

		if i == 5 { // HealthCareServiceLocation
			err = r.HealthCareServiceLocation.Validate(nil)
		} else if i == 11 { // RelatedCausesInformation
			if r.RelatedCausesInformation != nil {
				err = r.RelatedCausesInformation.Validate(nil)
			}
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i))
		}

		if err != nil {
			return fmt.Errorf("clm's element (%s) has invalid value, %s", idx, err.Error())
		}

	}

	return nil
}

func (r *CLM) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data, args...)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("clm segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("clm segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		rule := r.GetRule().Get(idx)

		if value, size, err = util.ReadField(line, read, rule, r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse clm's element (%s), %s", idx, err.Error())
		} else {
			read += size

			compositeRule := rule.Composite

			if i == 5 { // HealthCareServiceLocation
				var composite HealthCareServiceLocation
				if compositeRule != nil {
					composite.SetRule((*rules.ElementSetRule)(&compositeRule))
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.HealthCareServiceLocation = composite
				}

				if rules.IsMaskRequired(util.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse clm's element (%s), %s", idx, parseErr.Error())
				}

			} else if i == 11 { // RelatedCausesInformation
				var composite RelatedCausesInformation
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}
				if _, parseErr := composite.Parse(value, args...); parseErr == nil {
					r.RelatedCausesInformation = &composite
				}
			} else {
				r.SetFieldByIndex(idx, value)
			}

		}
	}

	return read, nil
}

func (r CLM) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {

		var value any
		idx := fmt.Sprintf("%02d", i)

		if i == 5 { // HealthCareServiceLocation
			value = r.HealthCareServiceLocation.String(args...)
		} else if i == 11 { // RelatedCausesInformation
			if r.RelatedCausesInformation != nil {
				value = r.RelatedCausesInformation.String(args...)
			}
		} else {
			value = r.GetFieldByIndex(idx)
		}

		if buf == "" {
			mask := r.GetRule().GetMask(idx, r.defaultMask(i))
			if mask == rules.MASK_NOTUSED {
				continue
			}
			if mask == rules.MASK_OPTIONAL && (value == nil || fmt.Sprintf("%v", value) == "") {
				continue
			}
		}

		if buf == "" {
			buf = fmt.Sprintf("%v%s", value, util.GetSegmentTerminator(args...))
		} else {
			buf = fmt.Sprintf("%v%s", value, util.DataElementSeparator) + buf
		}
	}

	if buf == "" {
		buf = fmt.Sprintf("%s%s", r.Name(), util.GetSegmentTerminator(args...))
	} else {
		buf = fmt.Sprintf("%s%s", r.Name(), util.DataElementSeparator) + buf
	}

	return buf
}
