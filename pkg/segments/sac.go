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

func NewSAC(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := SAC{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type SAC struct {
	Field01 string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	Field02 string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Field03 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Field04 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Field05 string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Field06 string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Field07 string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	Field08 string `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	Field09 string `index:"09" json:"09,omitempty" xml:"09,omitempty"`
	Field10 string `index:"10" json:"10,omitempty" xml:"10,omitempty"`
	Field11 string `index:"11" json:"11,omitempty" xml:"11,omitempty"`
	Field12 string `index:"12" json:"12,omitempty" xml:"12,omitempty"`
	Field13 string `index:"13" json:"13,omitempty" xml:"13,omitempty"`
	Field14 string `index:"14" json:"14,omitempty" xml:"14,omitempty"`
	Field15 string `index:"15" json:"15,omitempty" xml:"15,omitempty"`
	Field16 string `index:"16" json:"16,omitempty" xml:"16,omitempty"`

	Element
}

func (r SAC) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r SAC) fieldCount() int {
	return 16
}

func (r SAC) Name() string {
	return "SAC"
}

func (r *SAC) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r SAC) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *SAC) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("sac's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *SAC) Parse(data string, args ...string) (int, error) {
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

func (r SAC) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask())

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
