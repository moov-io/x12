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

func NewNM1(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := NM1{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type NM1 struct {
	EntityIdentifierCode        string `index:"01" json:"01" xml:"01"`
	EntityTypeQualifier         string `index:"02" json:"02" xml:"02"`
	OrganizationName            string `index:"03" json:"03" xml:"03"`
	NameFirst                   string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	NameMiddle                  string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	NamePrefix                  string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	NameSuffix                  string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	IdentificationCodeQualifier string `index:"08" json:"08" xml:"08"`
	IdentificationCode          string `index:"09" json:"09" xml:"09"`

	Element
}

func (r NM1) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index >= 5 && index <= 7 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r NM1) fieldCount() int {
	return 9
}

func (r NM1) Name() string {
	return "NM1"
}

func (r *NM1) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r NM1) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *NM1) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("nm1's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *NM1) Parse(data string, args ...string) (int, error) {
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

func (r NM1) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
