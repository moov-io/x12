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

func NewOI(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := OI{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type OI struct {
	IndicatorCode   string `index:"01" json:"01,omitempty" xml:"01,omitempty"`
	ReasonCode      string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	Response        string `index:"03" json:"03" xml:"03"`
	SourceCode      string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	AgreementCode   string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	InformationCode string `index:"06" json:"06" xml:"06"`

	Element
}

func (r *OI) defaultMask(index int) string {
	mask := rules.MASK_OPTIONAL
	if index == 3 || index == 6 {
		mask = rules.MASK_REQUIRED
	}
	return mask
}

func (r OI) Name() string {
	return "OI"
}

func (r *OI) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r OI) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *OI) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 6; i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("oi's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *OI) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size, read int

	length := util.GetRecordSize(data)
	if length < 2 {
		return 0, errors.New("oi segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:2] {
		return 0, errors.New("oi segment contains invalid code")
	}
	read += 3

	for i := 1; i <= 6; i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(i)); err != nil {
			return 0, fmt.Errorf("unable to parse oi's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r *OI) String(args ...string) string {
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
