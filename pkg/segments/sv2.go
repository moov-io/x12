// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"strings"

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

	for i := 1; i <= segmentFieldCount(r); i++ {
		var err error
		idx := util.GetFormattedIndex(i)

		if i == 1 {
			if r.Field2 != nil {
				cRule := rule.Get(idx).Composite
				err = r.Field2.Validate(&cRule)
			}
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), getFieldMask(r, i))
		}

		if err != nil {
			return util.NewValidateElementError(util.GetStructName(r), idx, err.Error())
		}
	}

	return nil
}

func (r *SV2) Parse(data string, args ...string) (int, error) {
	var size int
	name := strings.ToLower(util.GetStructName(r))
	read, line, err := r.VerifyCode(data, name, args...)
	if err != nil {
		return 0, err
	}

	for i := 1; i <= segmentFieldCount(r); i++ {
		var value string
		idx := util.GetFormattedIndex(i)

		rule := r.GetRule().Get(idx)

		if value, size, err = util.ReadField(line, read, rule, getFieldMask(r, i), args...); err != nil {
			return 0, util.NewParseSegmentError(name, idx, err.Error())
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

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			} else {
				r.SetFieldByIndex(idx, value)
			}

		}
	}

	return returnRead(read, data, name)
}

func (r SV2) String(args ...string) string {
	var buf string

	for i := segmentFieldCount(r); i > 0; i-- {
		var value any
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		if i == 2 {
			if r.Field2 != nil {
				value = r.Field2.String(args...)
			}
		} else {
			value = r.GetFieldByIndex(idx)
		}

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), value)
	}

	return r.TerminateString(buf, util.GetStructName(r))
}
