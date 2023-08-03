// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

type ServiceProcedure struct {
	ProductIdQualifier string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	ProductId1         string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	ProcedureModifier1 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	ProcedureModifier2 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	ProcedureModifier3 string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	ProcedureModifier4 string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Description        string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	ProductId2         string `index:"08" json:"08,omitempty" xml:"08,omitempty"`

	Element
}

func (r *ServiceProcedure) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r ServiceProcedure) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *ServiceProcedure) Validate(rule *rules.ElementSetRule) error {
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

func (r *ServiceProcedure) Parse(data string, args ...string) (int, error) {
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

func (r ServiceProcedure) String(args ...string) string {
	var buf string
	separator := util.GetElementSeparator(args...)

	for i := segmentFieldCount(r); i > 0; i-- {
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		buf = r.CompositeString(buf, mask, separator, "", r.GetFieldByIndex(idx))
	}

	return buf
}

type DentalServiceCode struct {
	Code1 string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Code2 string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Code3 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Code4 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Code5 string `index:"05" json:"05,omitempty" xml:"05,omitempty"`

	Element
}

func (r *DentalServiceCode) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r DentalServiceCode) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *DentalServiceCode) Validate(rule *rules.ElementSetRule) error {
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

func (r *DentalServiceCode) Parse(data string, args ...string) (int, error) {
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

func (r DentalServiceCode) String(args ...string) string {
	var buf string
	separator := util.GetElementSeparator(args...)

	for i := segmentFieldCount(r); i > 0; i-- {
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		buf = r.CompositeString(buf, mask, separator, "", r.GetFieldByIndex(idx))
	}

	return buf
}

type DentalServiceCodePointer struct {
	Pointer1 string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Pointer2 string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Pointer3 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Pointer4 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`

	Element
}

func (r *DentalServiceCodePointer) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r DentalServiceCodePointer) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *DentalServiceCodePointer) Validate(rule *rules.ElementSetRule) error {
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

func (r *DentalServiceCodePointer) Parse(data string, args ...string) (int, error) {
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

func (r DentalServiceCodePointer) String(args ...string) string {
	var buf string
	separator := util.GetElementSeparator(args...)

	for i := segmentFieldCount(r); i > 0; i-- {
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		buf = r.CompositeString(buf, mask, separator, "", r.GetFieldByIndex(idx))
	}

	return buf
}
