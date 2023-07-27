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

func NewIT1(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := IT1{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type IT1 struct {
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
	Field17 string `index:"17" json:"17,omitempty" xml:"17,omitempty"`
	Field18 string `index:"18" json:"18,omitempty" xml:"18,omitempty"`
	Field19 string `index:"19" json:"19,omitempty" xml:"19,omitempty"`
	Field20 string `index:"20" json:"20,omitempty" xml:"20,omitempty"`
	Field21 string `index:"21" json:"21,omitempty" xml:"21,omitempty"`
	Field22 string `index:"22" json:"22,omitempty" xml:"22,omitempty"`
	Field23 string `index:"23" json:"23,omitempty" xml:"23,omitempty"`
	Field24 string `index:"24" json:"24,omitempty" xml:"24,omitempty"`
	Field25 string `index:"25" json:"25,omitempty" xml:"25,omitempty"`

	Element
}

func (r IT1) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r IT1) fieldCount() int {
	return 25
}

func (r IT1) Name() string {
	return "IT1"
}

func (r *IT1) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r IT1) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *IT1) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("it1's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *IT1) Parse(data string, args ...string) (int, error) {
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

func (r IT1) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask())

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
