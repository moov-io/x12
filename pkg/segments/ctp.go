// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"strings"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

func NewCTP(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := CTP{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type CTP struct {
	Field01 string            `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Field02 string            `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Field03 string            `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Field04 string            `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Field05 *PricingComposite `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Field06 string            `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Field07 string            `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	Field08 string            `index:"08" json:"08,omitempty" xml:"08,omitempty"`

	Element
}

func (r *CTP) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r CTP) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *CTP) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= segmentFieldCount(r); i++ {
		var err error
		idx := util.GetFormattedIndex(i)

		if i == 5 {
			cRule := rule.Get(idx).Composite
			err = r.Field05.Validate(&cRule)
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), getFieldMask(r, i))
		}

		if err != nil {
			return util.NewValidateElementError(util.GetStructName(r), idx, err.Error())
		}
	}

	return nil
}

func (r *CTP) Parse(data string, args ...string) (int, error) {
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

			if i == 5 {
				var composite PricingComposite
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.Field05 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}

			} else {
				r.SetFieldByIndex(idx, value)
			}
		}
	}

	return returnRead(read, data, name, args...)
}

func (r CTP) String(args ...string) string {
	var buf string

	for i := segmentFieldCount(r); i > 0; i-- {
		var value any
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		if i == 5 {
			if r.Field05 != nil {
				value = r.Field05.String(args...)
			}
		} else {
			value = r.GetFieldByIndex(idx)
		}

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), value)
	}

	return r.TerminateString(buf, util.GetStructName(r))
}
