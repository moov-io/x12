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

func NewSV2(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := SV2{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type SV2 struct {
	Field1  string            `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Field2  *ServiceProcedure `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Field3  string            `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Field4  string            `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Field5  string            `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Field6  string            `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Field7  string            `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	Field8  string            `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	Field9  string            `index:"09" json:"09,omitempty" xml:"09,omitempty"`
	Field10 string            `index:"10" json:"10,omitempty" xml:"10,omitempty"`

	Element
}

func (r SV2) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r SV2) fieldCount() int {
	return 10
}

func (r SV2) Name() string {
	return "SV2"
}

func (r *SV2) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r SV2) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *SV2) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var err error
		idx := fmt.Sprintf("%02d", i)

		if i == 1 {
			if r.Field2 != nil {
				cRule := rule.Get(idx).Composite
				err = r.Field2.Validate(&cRule)
			}
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask())
		}

		if err != nil {
			return fmt.Errorf("SV2's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *SV2) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data, args...)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("sv2 segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("sv2 segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		rule := r.GetRule().Get(idx)

		if value, size, err = util.ReadField(line, read, rule, r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse sv2's element (%s), %s", idx, err.Error())
		} else {
			read += size

			compositeRule := rule.Composite

			if i == 2 {
				var composite ServiceProcedure
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.Field2 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask())) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse sv2's element (%s), %s", idx, parseErr.Error())
				}
			} else {
				r.SetFieldByIndex(idx, value)
			}

		}
	}

	return read, nil
}

func (r SV2) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		var value any

		if i == 2 {
			if r.Field2 != nil {
				value = r.Field2.String(args...)
			}
		} else {
			value = r.GetFieldByIndex(idx)
		}

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
