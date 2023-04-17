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

func NewHL(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := HL{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type HL struct {
	HierarchicalIdNumber       string `index:"01" json:"01" xml:"01"`
	HierarchicalParentIdNumber string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	HierarchicalLevelCode      string `index:"03" json:"03" xml:"03"`
	HierarchicalChildCode      string `index:"04" json:"04" xml:"04"`

	Element
}

func (r *HL) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index == 2 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r HL) fieldCount() int {
	return 4
}

func (r HL) Name() string {
	return "HL"
}

func (r *HL) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r HL) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *HL) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("hl's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *HL) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("hl segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("hl segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		mask := rules.MASK_REQUIRED
		idx := fmt.Sprintf("%02d", i)

		if i == 2 {
			mask = rules.MASK_OPTIONAL
		}

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), mask); err != nil {
			return 0, fmt.Errorf("unable to parse hl's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r HL) String(args ...string) string {
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
