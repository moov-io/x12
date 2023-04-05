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

func NewN3(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := N3{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type N3 struct {
	Street1 string `index:"01" json:"01" xml:"01"`
	Street2 string `index:"02" json:"02" xml:"02"`

	Element
}

func (r N3) Name() string {
	return "N3"
}

func (r *N3) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r N3) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *N3) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 2; i++ {

		idx := fmt.Sprintf("%02d", i)
		mask := rules.MASK_REQUIRED

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), mask); err != nil {
			return fmt.Errorf("n3's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *N3) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size, read int

	length := util.GetRecordSize(data)
	if length < 2 {
		return 0, errors.New("n3 segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:2] {
		return 0, errors.New("n3 segment contains invalid code")
	}
	read += 3

	for i := 1; i <= 2; i++ {

		var value string
		mask := rules.MASK_REQUIRED
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), mask); err != nil {
			return 0, fmt.Errorf("unable to parse n3's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r N3) String(args ...string) string {
	var buf string

	for i := 2; i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		value := r.GetFieldByIndex(idx)

		if buf == "" {
			mask := r.GetRule().GetMask(idx, rules.MASK_REQUIRED)
			if mask == rules.MASK_NOTUSED {
				continue
			}
			if mask == rules.MASK_OPTIONAL && (value == nil || fmt.Sprintf("%v", value) == "") {
				continue
			}
		}

		if buf == "" {
			buf = fmt.Sprintf("%v%s", value, util.SegmentTerminator)
		} else {
			buf = fmt.Sprintf("%v%s", value, util.DataElementSeparator) + buf
		}
	}

	if buf == "" {
		buf = fmt.Sprintf("%s%s", r.Name(), util.SegmentTerminator)
	} else {
		buf = fmt.Sprintf("%s%s", r.Name(), util.DataElementSeparator) + buf
	}

	return buf
}
