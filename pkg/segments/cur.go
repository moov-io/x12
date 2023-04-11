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

func NewCUR(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := CUR{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type CUR struct {
	EntityIdentifierCode string `index:"01" json:"01" xml:"01"`
	CurrencyCode         string `index:"02" json:"02" xml:"02"`

	Element
}

func (r CUR) Name() string {
	return "CUR"
}

func (r *CUR) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r CUR) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *CUR) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 2; i++ {

		idx := fmt.Sprintf("%02d", i)
		mask := rules.MASK_REQUIRED

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), mask); err != nil {
			return fmt.Errorf("cur's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *CUR) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size, read int

	length := util.GetRecordSize(data)
	if length < 3 {
		return 0, errors.New("cur segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:3] {
		return 0, errors.New("cur segment contains invalid code")
	}
	read += 4

	for i := 1; i <= 2; i++ {

		var value string
		mask := rules.MASK_REQUIRED
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), mask); err != nil {
			return 0, fmt.Errorf("unable to parse cur's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r *CUR) String(args ...string) string {
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