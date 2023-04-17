// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

type HealthCareServiceLocation struct {
	FacilityCodeValue     string `index:"01" json:"01" xml:"01"`
	FacilityCodeQualifier string `index:"02" json:"02" xml:"02"`
	ClaimFacilityType     string `index:"03" json:"03" xml:"03"`

	Element
}

func (r HealthCareServiceLocation) defaultMask() string {
	return rules.MASK_REQUIRED
}

func (r HealthCareServiceLocation) fieldCount() int {
	return 3
}

func (r *HealthCareServiceLocation) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r HealthCareServiceLocation) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *HealthCareServiceLocation) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("service location's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *HealthCareServiceLocation) Parse(data string, args ...string) (int, error) {

	var err error
	var size, read int
	line := data

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse service location's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r HealthCareServiceLocation) String(args ...string) string {
	var buf string

	separator := util.SubElementSeparator
	if len(args) > 0 {
		separator = args[0]
	}

	for i := r.fieldCount(); i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		value := r.GetFieldByIndex(idx)

		if buf == "" {
			mask := r.GetRule().GetMask(idx, r.defaultMask())
			if mask == rules.MASK_NOTUSED {
				continue
			}
			if mask == rules.MASK_OPTIONAL && (value == nil || fmt.Sprintf("%v", value) == "") {
				continue
			}
		}

		if buf == "" {
			buf = fmt.Sprintf("%s", value)
		} else {
			buf = fmt.Sprintf("%v%s", value, separator) + buf
		}
	}

	return buf
}

type RelatedCausesInformation struct {
	Code1   string `index:"01" json:"01" xml:"01"`
	Code2   string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Code3   string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	State   string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Country string `index:"05" json:"05,omitempty" xml:"05,omitempty"`

	Element
}

func (r RelatedCausesInformation) defaultMask(index int) string {
	if index > 1 {
		return rules.MASK_OPTIONAL
	}
	return rules.MASK_REQUIRED
}

func (r RelatedCausesInformation) fieldCount() int {
	return 5
}

func (r *RelatedCausesInformation) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r RelatedCausesInformation) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *RelatedCausesInformation) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("causes information's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *RelatedCausesInformation) Parse(data string, args ...string) (int, error) {

	var err error
	var size, read int
	line := data

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse causes information's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r RelatedCausesInformation) String(args ...string) string {
	var buf string

	separator := util.SubElementSeparator
	if len(args) > 0 {
		separator = args[0]
	}

	for i := r.fieldCount(); i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		value := r.GetFieldByIndex(idx)

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
			buf = fmt.Sprintf("%s", value)
		} else {
			buf = fmt.Sprintf("%v%s", value, separator) + buf
		}
	}

	return buf
}
