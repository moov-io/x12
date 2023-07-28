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

func NewTOO(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := TOO{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type TOO struct {
	CodeListQualCode string            `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	IndustryCode     string            `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	ToothSurfaceCode *ToothSurfaceCode `index:"03" json:"03,omitempty" xml:"03,omitempty"`

	Element
}

func (r TOO) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r TOO) fieldCount() int {
	return 3
}

func (r TOO) Name() string {
	return "TOO"
}

func (r *TOO) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r TOO) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *TOO) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		var err error
		idx := fmt.Sprintf("%02d", i)
		cRule := rule.Get(idx).Composite

		if i == 3 {
			if r.ToothSurfaceCode != nil {
				err = r.ToothSurfaceCode.Validate(&cRule)
			}
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask())
		}

		if err != nil {
			return fmt.Errorf("sv3's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *TOO) Parse(data string, args ...string) (int, error) {
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

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)

			compositeRule := rule.Composite

			if i == 3 {
				var composite ToothSurfaceCode
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.ToothSurfaceCode = &composite
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

func (r TOO) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		var value any
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask())

		if i == 3 {
			if r.ToothSurfaceCode != nil {
				value = r.ToothSurfaceCode.String(args...)
			}
		} else {
			value = r.GetFieldByIndex(idx)
		}

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), value)
	}

	return r.TerminateString(buf, r.Name())
}
