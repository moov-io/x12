// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

type ToothSurfaceCode struct {
	Code1 string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Code2 string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Code3 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Code4 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Code5 string `index:"05" json:"05,omitempty" xml:"05,omitempty"`

	Element
}

func (r ToothSurfaceCode) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r ToothSurfaceCode) fieldCount() int {
	return 5
}

func (r *ToothSurfaceCode) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r ToothSurfaceCode) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *ToothSurfaceCode) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("tooth surface code's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *ToothSurfaceCode) Parse(data string, args ...string) (int, error) {
	var err error
	var size, read int
	line := data

	for i := 1; i <= r.fieldCount(); i++ {
		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse tooth surface code's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r ToothSurfaceCode) String(args ...string) string {
	var buf string
	separator := util.GetElementSeparator(args...)

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask())

		buf = r.CompositeString(buf, mask, separator, "", r.GetFieldByIndex(idx))
	}

	return buf
}
