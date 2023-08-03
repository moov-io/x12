// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

type HealthCareServiceLocation struct {
	FacilityCodeValue     string `index:"01" json:"01" xml:"01"`
	FacilityCodeQualifier string `index:"02" json:"02" xml:"02"`
	ClaimFacilityType     string `index:"03" json:"03" xml:"03"`

	Element
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

	for i := 1; i <= segmentFieldCount(r); i++ {
		idx := util.GetFormattedIndex(i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), getFieldMask(r, i)); err != nil {
			return util.NewValidateElementError(util.GetStructName(r), idx, err.Error())
		}
	}

	return nil
}

func (r *HealthCareServiceLocation) Parse(data string, args ...string) (int, error) {
	var err error
	var size, read int
	line := data

	for i := 1; i <= segmentFieldCount(r); i++ {
		var value string
		idx := util.GetFormattedIndex(i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), getFieldMask(r, i), args...); err != nil {
			return 0, util.NewParseSegmentError(util.GetStructName(r), idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r HealthCareServiceLocation) String(args ...string) string {
	var buf string

	separator := util.GetElementSeparator(args...)

	for i := segmentFieldCount(r); i > 0; i-- {
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		buf = r.CompositeString(buf, mask, separator, "", r.GetFieldByIndex(idx))
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

	for i := 1; i <= segmentFieldCount(r); i++ {
		idx := util.GetFormattedIndex(i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), getFieldMask(r, i)); err != nil {
			return util.NewValidateElementError(util.GetStructName(r), idx, err.Error())
		}
	}

	return nil
}

func (r *RelatedCausesInformation) Parse(data string, args ...string) (int, error) {
	var err error
	var size, read int
	line := data

	for i := 1; i <= segmentFieldCount(r); i++ {
		var value string
		idx := util.GetFormattedIndex(i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), getFieldMask(r, i), args...); err != nil {
			return 0, util.NewParseSegmentError(util.GetStructName(r), idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r RelatedCausesInformation) String(args ...string) string {
	var buf string
	separator := util.GetElementSeparator(args...)

	for i := segmentFieldCount(r); i > 0; i-- {
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		buf = r.CompositeString(buf, mask, separator, "", r.GetFieldByIndex(idx))
	}

	return buf
}
