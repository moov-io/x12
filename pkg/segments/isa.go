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

func NewISA(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := ISA{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type ISA struct {
	AuthorizationQualifier    string `index:"01" json:"01" xml:"01"`
	AuthorizationInformation  string `index:"02" json:"02" xml:"02"`
	SecurityQualifier         string `index:"03" json:"03" xml:"03"`
	SecurityInformation       string `index:"04" json:"04" xml:"04"`
	SenderQualifier           string `index:"05" json:"05" xml:"05"`
	SenderId                  string `index:"06" json:"06" xml:"06"`
	ReceiverQualifier         string `index:"07" json:"07" xml:"07"`
	ReceiverId                string `index:"08" json:"08" xml:"08"`
	Date                      string `index:"09" json:"09" xml:"09"`
	Time                      string `index:"10" json:"10" xml:"10"`
	StandardsId               string `index:"11" json:"11" xml:"11"`
	Version                   string `index:"12" json:"12" xml:"12"`
	InterchangeControlNumber  string `index:"13" json:"13" xml:"13"`
	AcknowledgmentRequested   string `index:"14" json:"14" xml:"14"`
	TestIndicator             string `index:"15" json:"15" xml:"15"`
	ComponentElementSeparator string `index:"16" json:"16" xml:"16"`

	Element
}

func (r ISA) defaultMask() string {
	return rules.MASK_REQUIRED
}

func (r ISA) fieldCount() int {
	return 16
}

func (r ISA) Name() string {
	return "ISA"
}

func (r *ISA) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r ISA) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *ISA) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask()); err != nil {
			return fmt.Errorf("isa's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *ISA) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data, args...)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("isa segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("isa segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(), args...); err != nil {
			return 0, fmt.Errorf("unable to parse isa's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r ISA) String(args ...string) string {
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
