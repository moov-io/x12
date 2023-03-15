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

func NewGS(rule *rules.Elements) SegmentInterface {

	newSegment := GS{}

	if rule == nil {
		newRule := make(rules.Elements)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type GS struct {
	FunctionalIdentifierCode string `index:"01" json:"01" xml:"01"`
	ApplicationSenderCode    string `index:"02" json:"02" xml:"02"`
	ApplicationReceiverCode  string `index:"03" json:"03" xml:"03"`
	Date                     string `index:"04" json:"04" xml:"04"`
	Time                     string `index:"05" json:"05" xml:"05"`
	GroupControlNumber       string `index:"06" json:"06" xml:"06"`
	ResponsibleAgencyCode    string `index:"07" json:"07" xml:"07"`
	Version                  string `index:"08" json:"08" xml:"08"`

	Element
}

func (r GS) Name() string {
	return "GS"
}

func (r *GS) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r GS) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *GS) Validate(rule *rules.Elements) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 8; i++ {

		idx := fmt.Sprintf("%02d", i)
		mask := rules.MASK_REQUIRED

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), mask); err != nil {
			return fmt.Errorf("gs's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *GS) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size, read int

	length := util.GetRecordSize(data)
	if length < 2 {
		return 0, errors.New("gs segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:2] {
		return 0, errors.New("gs segment contains invalid code")
	}
	read += 3

	for i := 1; i <= 8; i++ {

		var value string
		mask := rules.MASK_REQUIRED
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), mask); err != nil {
			return 0, fmt.Errorf("unable to parse gs's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r *GS) String(args ...string) string {
	var buf string

	for i := 8; i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		value := r.GetFieldByIndex(idx)

		if buf == "" {
			mask := r.GetRule().GetMask(idx, rules.MASK_REQUIRED)
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
