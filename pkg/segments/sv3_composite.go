// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"

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

func (r ServiceProcedure) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r ServiceProcedure) fieldCount() int {
	return 8
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

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("dental procedure's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *ServiceProcedure) Parse(data string, args ...string) (int, error) {

	var err error
	var size, read int
	line := data

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse dental procedure's element (%s), %s", idx, err.Error())
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

type DentalServiceCode struct {
	Code1 string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Code2 string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Code3 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Code4 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Code5 string `index:"05" json:"05,omitempty" xml:"05,omitempty"`

	Element
}

func (r DentalServiceCode) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r DentalServiceCode) fieldCount() int {
	return 5
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

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("dental code's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *DentalServiceCode) Parse(data string, args ...string) (int, error) {

	var err error
	var size, read int
	line := data

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse dental code's element (%s), %s", idx, err.Error())
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

type DentalServiceCodePointer struct {
	Pointer1 string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Pointer2 string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Pointer3 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Pointer4 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`

	Element
}

func (r DentalServiceCodePointer) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r DentalServiceCodePointer) fieldCount() int {
	return 4
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

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("dental code pointer's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *DentalServiceCodePointer) Parse(data string, args ...string) (int, error) {

	var err error
	var size, read int
	line := data

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse dental code pointer's element (%s), %s", idx, err.Error())
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
