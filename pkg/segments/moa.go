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

func NewMOA(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := MOA{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type MOA struct {
	Percentage      string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	MonetaryAmount1 string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Identification1 string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Identification2 string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Identification3 string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Identification4 string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Identification5 string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	MonetaryAmount2 string `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	MonetaryAmount3 string `index:"09" json:"09,omitempty" xml:"09,omitempty"`

	Element
}

func (r MOA) defaultMask() string {
	return rules.MASK_OPTIONAL
}

func (r MOA) fieldCount() int {
	return 9
}

func (r MOA) Name() string {
	return "MOA"
}

func (r *MOA) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r MOA) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *MOA) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("moa's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *MOA) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data, args...)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("moa segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("moa segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse moa's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r MOA) String(args ...string) string {
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
