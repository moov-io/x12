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

func NewNM1(rule *rules.Elements) SegmentInterface {

	newSegment := NM1{}

	if rule == nil {
		newRule := make(rules.Elements)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type NM1 struct {
	EntityIdentifierCode        string `index:"01" json:"01" xml:"01"`
	EntityTypeQualifier         string `index:"02" json:"02" xml:"02"`
	OrganizationName            string `index:"03" json:"03" xml:"03"`
	NameFirst                   string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	NameMiddle                  string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	NamePrefix                  string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	NameSuffix                  string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	IdentificationCodeQualifier string `index:"08" json:"08" xml:"08"`
	IdentificationCode          string `index:"09" json:"09" xml:"09"`

	Element
}

func (r *NM1) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index >= 5 && index <= 7 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r NM1) Name() string {
	return "NM1"
}

func (r *NM1) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r NM1) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *NM1) Validate(rule *rules.Elements) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 9; i++ {

		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("nm1's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *NM1) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size, read int

	length := util.GetRecordSize(data)
	if length < 3 {
		return 0, errors.New("nm1 segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:3] {
		return 0, errors.New("nm1 segment contains invalid code")
	}
	read += 4

	for i := 1; i <= 9; i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(i)); err != nil {
			return 0, fmt.Errorf("unable to parse nm1's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r NM1) String(args ...string) string {
	var buf string

	for i := 9; i > 0; i-- {

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
