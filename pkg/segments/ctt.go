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

func NewCTT(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := CTT{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type CTT struct {
	Field01 string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Field02 string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Field03 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Field04 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Field05 string `index:"05" json:"05,omitempty" xml:"05,omitempty"`

	Element
}

func (r CTT) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r CTT) fieldCount() int {
	return 5
}

func (r CTT) Name() string {
	return "CTT"
}

func (r *CTT) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r CTT) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *CTT) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("ctt's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *CTT) Parse(data string, args ...string) (int, error) {
	var size int
	name := strings.ToLower(r.Name())
	read, line, err := r.VerifyCode(data, name, args...)
	if err != nil {
		return 0, err
	}

	for i := 1; i <= r.fieldCount(); i++ {
		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r CTT) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask())

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
