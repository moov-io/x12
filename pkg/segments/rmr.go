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

func NewRMR(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := RMR{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type RMR struct {
	ReferenceIdentificationQualifier string `index:"01" json:"01" xml:"01"`
	ReferenceIdentification          string `index:"02" json:"02" xml:"02"`
	PaymentActionCode                string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	MonetaryAmount1                  string `index:"04" json:"04" xml:"04"`
	MonetaryAmount2                  string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	MonetaryAmount3                  string `index:"06" json:"06,omitempty" xml:"06,omitempty"`

	Element
}

func (r RMR) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index >= 3 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r RMR) fieldCount() int {
	return 4
}

func (r RMR) Name() string {
	return "RMR"
}

func (r *RMR) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r RMR) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *RMR) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("rmr's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *RMR) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data, args...)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("rmr segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("rmr segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse rmr's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r RMR) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		value := r.GetFieldByIndex(idx)

		if buf == "" {
			mask := r.GetRule().GetMask(idx, r.defaultMask(i))
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
