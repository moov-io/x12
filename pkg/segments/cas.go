// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"strings"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

func NewCAS(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := CAS{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type CAS struct {
	AdjustmentGroupCode   string `index:"01" json:"01" xml:"01"`
	AdjustmentReasonCode  string `index:"02" json:"02" xml:"02"`
	MonetaryAmount        string `index:"03" json:"03" xml:"03"`
	Quantity              string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	AdjustmentReasonCode1 string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	MonetaryAmount1       string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Quantity1             string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	AdjustmentReasonCode2 string `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	MonetaryAmount2       string `index:"09" json:"09,omitempty" xml:"09,omitempty"`
	Quantity2             string `index:"10" json:"10,omitempty" xml:"10,omitempty"`
	AdjustmentReasonCode3 string `index:"11" json:"11,omitempty" xml:"11,omitempty"`
	MonetaryAmount3       string `index:"12" json:"12,omitempty" xml:"12,omitempty"`
	Quantity3             string `index:"13" json:"13,omitempty" xml:"13,omitempty"`
	AdjustmentReasonCode4 string `index:"14" json:"14,omitempty" xml:"14,omitempty"`
	MonetaryAmount4       string `index:"15" json:"15,omitempty" xml:"15,omitempty"`
	Quantity4             string `index:"16" json:"16,omitempty" xml:"16,omitempty"`
	AdjustmentReasonCode5 string `index:"17" json:"17,omitempty" xml:"17,omitempty"`
	MonetaryAmount5       string `index:"18" json:"18,omitempty" xml:"18,omitempty"`
	Quantity5             string `index:"19" json:"19,omitempty" xml:"19,omitempty"`

	Element
}

func (r *CAS) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r CAS) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *CAS) Validate(rule *rules.ElementSetRule) error {
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

func (r *CAS) Parse(data string, args ...string) (int, error) {
	var size int
	name := strings.ToLower(util.GetStructName(r))
	read, line, err := r.VerifyCode(data, name, args...)
	if err != nil {
		return 0, err
	}

	for i := 1; i <= segmentFieldCount(r); i++ {
		var value string
		idx := util.GetFormattedIndex(i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), getFieldMask(r, i), args...); err != nil {
			return 0, util.NewParseSegmentError(name, idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return returnRead(read, data, name)
}

func (r CAS) String(args ...string) string {
	var buf string

	for i := segmentFieldCount(r); i > 0; i-- {
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, util.GetStructName(r))
}
