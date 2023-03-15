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

func NewST(rule *rules.Elements) SegmentInterface {

	newSegment := ST{}

	if rule == nil {
		newRule := make(rules.Elements)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type ST struct {
	TransactionSetIdentifierCode      string `index:"01" json:"01" xml:"01"`
	TransactionSetControlNumber       string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	ImplementationConventionReference string `index:"03" json:"03,omitempty" xml:"03,omitempty"`

	Element
}

func (r *ST) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index >= 2 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r ST) Name() string {
	return "ST"
}

func (r *ST) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r ST) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *ST) Validate(rule *rules.Elements) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 3; i++ {

		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("st's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *ST) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size, read int

	length := util.GetRecordSize(data)
	if length < 2 {
		return 0, errors.New("st segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:2] {
		return 0, errors.New("st segment contains invalid code")
	}
	read += 3

	for i := 1; i <= 3; i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(i)); err != nil {
			return 0, fmt.Errorf("unable to parse st's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r *ST) String(args ...string) string {

	var buf string

	for i := 3; i > 0; i-- {

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
