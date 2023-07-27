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

func NewDTM(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := DTM{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type DTM struct {
	Qualifier       string `index:"01" json:"01" xml:"01"`
	Date            string `index:"02" json:"02" xml:"02"`
	Time            string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	TimeCode        string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	FormatQualifier string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	DateTimePeriod  string `index:"06" json:"06,omitempty" xml:"06,omitempty"`

	Element
}

func (r DTM) defaultMask(index int) string {
	if index < 3 {
		return rules.MASK_REQUIRED
	}
	return rules.MASK_OPTIONAL
}

func (r DTM) fieldCount() int {
	return 6
}

func (r DTM) Name() string {
	return "DTM"
}

func (r *DTM) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r DTM) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *DTM) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("dtm's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *DTM) Parse(data string, args ...string) (int, error) {
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

func (r DTM) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
