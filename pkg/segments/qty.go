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

func NewQTY(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := QTY{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type QTY struct {
	Qualifier string        `index:"01" json:"01" xml:"01"`
	Quantity  string        `index:"02" json:"02" xml:"02"`
	Field03   *QtyComposite `index:"03" json:"03" xml:"03"`
	Field04   string        `index:"04" json:"04" xml:"04"`

	Element
}

func (r QTY) defaultMask(index int) string {
	if index < 3 {
		return rules.MASK_REQUIRED
	}
	return rules.MASK_OPTIONAL
}

func (r QTY) fieldCount() int {
	return 4
}

func (r QTY) Name() string {
	return "QTY"
}

func (r *QTY) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r QTY) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *QTY) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		var err error
		idx := fmt.Sprintf("%02d", i)

		if i == 3 {
			if r.Field03 != nil {
				cRule := rule.Get(idx).Composite
				err = r.Field03.Validate(&cRule)
			}
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i))
		}

		if err != nil {
			return fmt.Errorf("qty's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *QTY) Parse(data string, args ...string) (int, error) {
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

			if i == 3 {
				var composite QtyComposite
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.Field03 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}
			} else {
				r.SetFieldByIndex(idx, value)
			}
		}
	}

	return read, nil
}

func (r QTY) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		var value any
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		if i == 3 {
			if r.Field03 != nil {
				value = r.Field03.String(args...)
			}
		} else {
			value = r.GetFieldByIndex(idx)
		}

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), value)
	}

	return r.TerminateString(buf, r.Name())
}
