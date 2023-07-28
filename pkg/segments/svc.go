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

func NewSVC(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := SVC{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type SVC struct {
	Field1 ServiceProcedure   `index:"01" json:"01" xml:"01"`
	Field2 string             `index:"02" json:"02" xml:"02"`
	Field3 string             `index:"03" json:"03" xml:"03"`
	Field4 string             `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Field5 string             `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Field6 *DentalServiceCode `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Field7 string             `index:"07" json:"07,omitempty" xml:"07,omitempty"`

	Element
}

func (r SVC) defaultMask(index int) string {
	if index > 3 {
		return rules.MASK_OPTIONAL
	}
	return rules.MASK_REQUIRED
}

func (r SVC) fieldCount() int {
	return 7
}

func (r SVC) Name() string {
	return "SVC"
}

func (r *SVC) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r SVC) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *SVC) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		var err error
		idx := fmt.Sprintf("%02d", i)

		if i == 1 {
			err = r.Field1.Validate(nil)
		} else if i == 6 {
			if r.Field6 != nil {
				err = r.Field6.Validate(nil)
			}
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i))
		}

		if err != nil {
			return fmt.Errorf("svc's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *SVC) Parse(data string, args ...string) (int, error) {
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

		if value, size, err = util.ReadField(line, read, rule, r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, err.Error())
		} else {
			read += size
			compositeRule := rule.Composite

			if i == 1 {
				var composite ServiceProcedure
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				if _, parseErr := composite.Parse(value, args...); parseErr == nil {
					r.Field1 = composite
				} else if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}

			} else if i == 6 {
				var composite DentalServiceCode
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				if _, parseErr := composite.Parse(value, args...); parseErr == nil {
					r.Field6 = &composite
				} else if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}
			} else {
				r.SetFieldByIndex(idx, value)
			}

		}
	}

	return read, nil
}

func (r SVC) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		var value any
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		if i == 1 {
			value = r.Field1.String(args...)
		} else if i == 6 {
			if r.Field6 != nil {
				value = r.Field6.String(args...)
			}
		} else {
			value = r.GetFieldByIndex(idx)
		}

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), value)
	}

	return r.TerminateString(buf, r.Name())
}
