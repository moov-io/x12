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

func NewSV3(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := SV3{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type SV3 struct {
	Field1  *DentalServiceProcedure   `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Field2  string                    `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Field3  string                    `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Field4  *DentalServiceCode        `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Field5  string                    `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Field6  string                    `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Field7  string                    `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	Field8  string                    `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	Field9  string                    `index:"09" json:"09,omitempty" xml:"09,omitempty"`
	Field10 string                    `index:"10" json:"10,omitempty" xml:"10,omitempty"`
	Field11 *DentalServiceCodePointer `index:"11" json:"11,omitempty" xml:"11,omitempty"`

	Element
}

func (r *SV3) defaultMask(index int) string {
	return rules.MASK_OPTIONAL
}

func (r SV3) Name() string {
	return "SV3"
}

func (r *SV3) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r SV3) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *SV3) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 21; i++ {

		var err error
		idx := fmt.Sprintf("%02d", i)

		if i == 1 {
			if r.Field1 != nil {
				err = r.Field1.Validate(nil)
			}
		} else if i == 4 {
			if r.Field4 != nil {
				err = r.Field4.Validate(nil)
			}
		} else if i == 11 {
			if r.Field11 != nil {
				err = r.Field11.Validate(nil)
			}
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i))
		}

		if err != nil {
			return fmt.Errorf("sv3's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *SV3) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size, read int

	length := util.GetRecordSize(data)
	if length < 3 {
		return 0, errors.New("sv3 segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:3] {
		return 0, errors.New("sv3 segment contains invalid code")
	}
	read += 4

	for i := 1; i <= 21; i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		rule := r.GetRule().Get(idx)

		if value, size, err = util.ReadField(line, read, rule, r.defaultMask(i)); err != nil {
			return 0, fmt.Errorf("unable to parse sv3's element (%s), %s", idx, err.Error())
		} else {
			read += size

			compositeRule := rule.Composite

			if i == 1 {
				var composite DentalServiceProcedure
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}
				if _, parseErr := composite.Parse(value, args...); parseErr == nil {
					r.Field1 = &composite
				}
			} else if i == 4 {
				var composite DentalServiceCode
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}
				if _, parseErr := composite.Parse(value, args...); parseErr == nil {
					r.Field4 = &composite
				}
			} else if i == 11 {
				var composite DentalServiceCodePointer
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}
				if _, parseErr := composite.Parse(value, args...); parseErr == nil {
					r.Field11 = &composite
				}
			} else {
				r.SetFieldByIndex(idx, value)
			}

		}
	}

	return read, nil
}

func (r *SV3) String(args ...string) string {
	var buf string

	for i := 21; i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		var value any

		if i == 1 {
			if r.Field1 != nil {
				value = r.Field1.String(args...)
			}
		} else if i == 4 {
			if r.Field4 != nil {
				value = r.Field4.String(args...)
			}
		} else if i == 11 {
			if r.Field11 != nil {
				value = r.Field11.String(args...)
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
			buf = fmt.Sprintf("%v%s", value, util.SegmentTerminator)
		} else {
			buf = fmt.Sprintf("%v%s", value, util.DataElementSeparator) + buf
		}
	}

	if buf == "" {
		buf = fmt.Sprintf("%s%s", r.Name(), util.SegmentTerminator)
	} else {
		buf = fmt.Sprintf("%s%s", r.Name(), util.DataElementSeparator) + buf
	}

	return buf
}
