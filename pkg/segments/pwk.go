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

func NewPWK(rule *rules.Elements) SegmentInterface {

	newSegment := PWK{}

	if rule == nil {
		newRule := make(rules.Elements)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type PWK struct {
	ReportTypeCode          string `index:"01" json:"01" xml:"01"`
	ReportTransmissionCode  string `index:"02" json:"02" xml:"02"`
	ReportCopiesNeeded      string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	EntityIdentifierCode    string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	IdentifierCodeQualifier string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	IdentifierCode          string `index:"06" json:"06,omitempty" xml:"06,omitempty"`

	Element
}

func (r *PWK) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index > 2 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r PWK) Name() string {
	return "PWK"
}

func (r *PWK) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r PWK) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *PWK) Validate(rule *rules.Elements) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 6; i++ {
		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("pwk's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *PWK) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size, read int

	length := util.GetRecordSize(data)
	if length < 3 {
		return 0, errors.New("pwk segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:3] {
		return 0, errors.New("pwk segment contains invalid code")
	}
	read += 4

	for i := 1; i <= 6; i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(i)); err != nil {
			return 0, fmt.Errorf("unable to parse pwk's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r PWK) String(args ...string) string {

	var buf string

	for i := 6; i > 0; i-- {

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
