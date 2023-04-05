// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

type HealthCareCode struct {
	TypeCode     string `index:"01" json:"01" xml:"01"`
	Code         string `index:"02" json:"02" xml:"02"`
	PeriodFormat string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Period       string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Amount       string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Quantity     string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Identifier   string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	IndustryCode string `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	ResponseCode string `index:"09" json:"09,omitempty" xml:"09,omitempty"`

	Element
}

func (r *HealthCareCode) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index > 2 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r *HealthCareCode) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r HealthCareCode) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *HealthCareCode) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 9; i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("health care code's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *HealthCareCode) Parse(data string, args ...string) (int, error) {

	var err error
	var size, read int
	line := data

	for i := 1; i <= 9; i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse health care code's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r *HealthCareCode) String(args ...string) string {
	var buf string

	separator := util.SubElementSeparator
	if len(args) > 0 {
		separator = args[0]
	}

	for i := 9; i > 0; i-- {

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
