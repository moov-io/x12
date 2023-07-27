// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

type MonetaryIdentification struct {
	AdjustmentReasonCode    string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	ReferenceIdentification string `index:"02" json:"02,omitempty" xml:"02,omitempty"`

	Element
}

func (r MonetaryIdentification) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r MonetaryIdentification) fieldCount() int {
	return 2
}

func (r *MonetaryIdentification) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r MonetaryIdentification) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *MonetaryIdentification) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("monetary identification's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *MonetaryIdentification) Parse(data string, args ...string) (int, error) {
	var err error
	var size, read int
	line := data

	for i := 1; i <= r.fieldCount(); i++ {
		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse monetary identification's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r MonetaryIdentification) String(args ...string) string {
	var buf string
	separator := util.GetElementSeparator(args...)

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask())

		buf = r.CompositeString(buf, mask, separator, "", r.GetFieldByIndex(idx))
	}

	return buf
}
