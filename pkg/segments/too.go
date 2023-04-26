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

func NewTOO(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := TOO{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type TOO struct {
	CodeListQualCode string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	IndustryCode     string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Identification1  string `index:"03" json:"03,omitempty" xml:"03,omitempty"`

	Element
}

func (r TOO) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r TOO) fieldCount() int {
	return 3
}

func (r TOO) Name() string {
	return "TOO"
}

func (r *TOO) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r TOO) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *TOO) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("too's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *TOO) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data, args...)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("too segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("too segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse too's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r TOO) String(args ...string) string {
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
