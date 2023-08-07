// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

type MonetaryIdentification struct {
	AdjustmentReasonCode    string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	ReferenceIdentification string `index:"02" json:"02,omitempty" xml:"02,omitempty"`

	Element
}

func (r *MonetaryIdentification) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r MonetaryIdentification) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *MonetaryIdentification) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= segmentFieldCount(r); i++ {
		idx := util.GetFormattedIndex(i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), getFieldMask(r, i)); err != nil {
			return util.NewValidateElementError(util.GetStructName(r), idx, err.Error())
		}
	}

	return nil
}

func (r *MonetaryIdentification) Parse(data string, args ...string) (int, error) {
	var err error
	var size, read int
	line := data

	for i := 1; i <= segmentFieldCount(r); i++ {
		var value string
		idx := util.GetFormattedIndex(i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), getFieldMask(r, i), args...); err != nil {
			return 0, util.NewParseSegmentError(util.GetStructName(r), idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r MonetaryIdentification) String(args ...string) string {
	var buf string
	separator := util.GetElementSeparator(args...)

	for i := segmentFieldCount(r); i > 0; i-- {
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		buf = r.CompositeString(buf, mask, separator, "", r.GetFieldByIndex(idx))
	}

	return buf
}
