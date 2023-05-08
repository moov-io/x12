// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"errors"
	"fmt"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

func NewTS2(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := TS2{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type TS2 struct {
	MonetaryAmount1  string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	MonetaryAmount2  string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	MonetaryAmount3  string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	MonetaryAmount4  string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	MonetaryAmount5  string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	MonetaryAmount6  string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Quantity1        string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	MonetaryAmount7  string `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	MonetaryAmount8  string `index:"09" json:"09,omitempty" xml:"09,omitempty"`
	Quantity2        string `index:"10" json:"10,omitempty" xml:"10,omitempty"`
	Quantity3        string `index:"11" json:"11,omitempty" xml:"11,omitempty"`
	Quantity4        string `index:"12" json:"12,omitempty" xml:"12,omitempty"`
	Quantity5        string `index:"13" json:"13,omitempty" xml:"13,omitempty"`
	Quantity6        string `index:"14" json:"14,omitempty" xml:"14,omitempty"`
	MonetaryAmount9  string `index:"15" json:"15,omitempty" xml:"15,omitempty"`
	Quantity7        string `index:"16" json:"16,omitempty" xml:"16,omitempty"`
	MonetaryAmount10 string `index:"17" json:"17,omitempty" xml:"17,omitempty"`
	MonetaryAmount11 string `index:"18" json:"18,omitempty" xml:"18,omitempty"`
	MonetaryAmount12 string `index:"19" json:"19,omitempty" xml:"19,omitempty"`

	Element
}

func (r TS2) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r TS2) fieldCount() int {
	return 19
}

func (r TS2) Name() string {
	return "TS2"
}

func (r *TS2) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r TS2) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *TS2) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("ts2's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *TS2) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data, args...)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("ts2 segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("ts2 segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse ts2's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r TS2) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		value := r.GetFieldByIndex(idx)

		if buf == "" {
			mask := r.GetRule().GetMask(idx, r.defaultMask())
			if mask == rules.MASK_NOTUSED {
				continue
			}
			if mask == rules.MASK_OPTIONAL && (value == nil || fmt.Sprintf("%v", value) == "") {
				continue
			}
		}

		if buf == "" {
			buf = fmt.Sprintf("%v%s", value, util.GetSegmentTerminator(args...))
		} else {
			buf = fmt.Sprintf("%v%s", value, util.DataElementSeparator) + buf
		}
	}

	if buf == "" {
		buf = fmt.Sprintf("%s%s", r.Name(), util.GetSegmentTerminator(args...))
	} else {
		buf = fmt.Sprintf("%s%s", r.Name(), util.DataElementSeparator) + buf
	}

	return buf
}
