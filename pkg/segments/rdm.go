// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"
	"strings"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

func NewRDM(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := RDM{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type RDM struct {
	ReportTransmissionCode string `index:"01" json:"01" xml:"01"`
	MethodName             string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	CommunicationNumber    string `index:"03" json:"03,omitempty" xml:"03,omitempty"`

	Element
}

func (r RDM) defaultMask(index int) string {
	if index > 2 {
		return rules.MASK_OPTIONAL
	}
	return rules.MASK_REQUIRED
}

func (r RDM) fieldCount() int {
	return 3
}

func (r RDM) Name() string {
	return "RDM"
}

func (r *RDM) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r RDM) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *RDM) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("rdm's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *RDM) Parse(data string, args ...string) (int, error) {
	var size int
	name := strings.ToLower(r.Name())
	read, line, err := r.VerifyCode(data, name, args...)
	if err != nil {
		return 0, err
	}

	for i := 1; i <= r.fieldCount(); i++ {
		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadField(line, read, r.GetRule().Get(idx), r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r RDM) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
