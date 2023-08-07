// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"strings"

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

	for i := 1; i <= segmentFieldCount(r); i++ {
		idx := util.GetFormattedIndex(i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), getFieldMask(r, i)); err != nil {
			return util.NewValidateElementError(util.GetStructName(r), idx, err.Error())
		}
	}

	return nil
}

func (r *ISA) Parse(data string, args ...string) (int, error) {
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

	return returnRead(read, data, name, args...)
}

func (r ISA) String(args ...string) string {
	var buf string

	for i := segmentFieldCount(r); i > 0; i-- {
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, util.GetStructName(r))
}
