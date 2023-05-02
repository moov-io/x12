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

func NewN1(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := N1{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type N1 struct {
	EntityIdentifierCode        string `index:"01" json:"01" xml:"01"`
	OriginalName                string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	IdentificationCodeQualifier string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	IdentificationCodeUniqueID  string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Field05                     string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Field06                     string `index:"06" json:"06,omitempty" xml:"06,omitempty"`

	Element
}

func (r N1) defaultMask(index int) string {
	if index > 2 {
		return rules.MASK_OPTIONAL
	}
	return rules.MASK_REQUIRED
}

func (r N1) fieldCount() int {
	return 6
}

func (r N1) Name() string {
	return "N1"
}

func (r *N1) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r N1) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *N1) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("n1's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *N1) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data, args...)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("n1 segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("n1 segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse n1's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r N1) String(args ...string) string {
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

	buf = fmt.Sprintf("%s%s", r.Name(), util.DataElementSeparator) + buf

	return buf
}
