// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

type QtyComposite struct {
	Field01 string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Field02 string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Field03 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Field04 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Field05 string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Field06 string `index:"06" json:"06,omitempty" xml:"06,omitempty"`

	Element
}

func (r QtyComposite) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r QtyComposite) fieldCount() int {
	return 6
}

func (r *QtyComposite) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r QtyComposite) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *QtyComposite) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("qty composite's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *QtyComposite) Parse(data string, args ...string) (int, error) {
	var err error
	var size, read int
	line := data

	for i := 1; i <= r.fieldCount(); i++ {
		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse qty composite's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r QtyComposite) String(args ...string) string {
	var buf string
	separator := util.GetElementSeparator(args...)

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask())

		buf = r.CompositeString(buf, mask, separator, "", r.GetFieldByIndex(idx))
	}

	return buf
}
