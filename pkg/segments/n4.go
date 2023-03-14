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

func NewN4(rule *rules.Elements) SegmentInterface {

	newSegment := N4{}

	if rule == nil {
		newRule := make(rules.Elements)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type N4 struct {
	City       string `index:"01" json:"01" xml:"01"`
	State      string `index:"02" json:"02" xml:"02"`
	PostalCode string `index:"03" json:"03" xml:"03"`

	Element
}

func (r N4) Name() string {
	return "N4"
}

func (r *N4) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r N4) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *N4) Validate(rule *rules.Elements) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 3; i++ {

		idx := fmt.Sprintf("%02d", i)
		mask := rules.MASK_REQUIRED

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), mask); err != nil {
			return fmt.Errorf("n4's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *N4) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size, read int

	length := util.GetRecordSize(data)
	if length < 2 {
		return 0, errors.New("n4 segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:2] {
		return 0, errors.New("n4 segment contains invalid code")
	}
	read += 3

	for i := 1; i <= 3; i++ {

		var value string
		mask := rules.MASK_REQUIRED
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), mask); err != nil {
			return 0, fmt.Errorf("unable to parse n4's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r N4) String(args ...string) string {
	var buf string

	for i := 3; i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		value := r.GetFieldByIndex(idx)

		if buf == "" {
			if r.GetRule().Get(idx).Mask == rules.MASK_NOTUSED {
				continue
			}
			if r.GetRule().Get(idx).Mask == rules.MASK_OPTIONAL && value == nil {
				continue
			}
		}

		if buf == "" {
			buf = fmt.Sprintf("%v%s", value, util.SegmentTerminator)
		} else {
			buf = fmt.Sprintf("%v%s", value, util.DataElementSeparator) + buf
		}
	}

	buf = fmt.Sprintf("%s%s", r.Name(), util.DataElementSeparator) + buf

	return buf
}
