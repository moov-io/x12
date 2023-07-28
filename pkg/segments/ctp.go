// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"
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

func (r CTP) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r CTP) fieldCount() int {
	return 8
}

func (r CTP) Name() string {
	return "CTP"
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

	for i := 1; i <= r.fieldCount(); i++ {
		var err error
		idx := fmt.Sprintf("%02d", i)

		if i == 5 {
			cRule := rule.Get(idx).Composite
			err = r.Field05.Validate(&cRule)
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask())
		}

		if err != nil {
			return fmt.Errorf("ctp's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *CTP) Parse(data string, args ...string) (int, error) {
	var size int
	name := strings.ToLower(r.Name())
	read, line, err := r.VerifyCode(data, name, args...)
	if err != nil {
		return 0, err
	}

	for i := 1; i <= r.fieldCount(); i++ {
		var value string
		idx := fmt.Sprintf("%02d", i)
		rule := r.GetRule().Get(idx)

		if value, size, err = util.ReadField(line, read, rule, r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, err.Error())
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

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask())) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}

			} else {
				r.SetFieldByIndex(idx, value)
			}
		}
	}

	return read, nil
}

func (r CTP) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		var value any
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask())

		if i == 5 {
			if r.Field05 != nil {
				value = r.Field05.String(args...)
			}
		} else {
			value = r.GetFieldByIndex(idx)
		}

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), value)
	}

	return r.TerminateString(buf, r.Name())
}
