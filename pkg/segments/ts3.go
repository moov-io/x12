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

func NewTS3(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := TS3{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type TS3 struct {
	ReferenceIdentification string `index:"01" json:"01" xml:"01"`
	FacilityCodeValue       string `index:"02" json:"02" xml:"02"`
	Date                    string `index:"03" json:"03" xml:"03"`
	Quantity1               string `index:"04" json:"04" xml:"04"`
	MonetaryAmount1         string `index:"05" json:"05" xml:"05"`
	MonetaryAmount2         string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	MonetaryAmount3         string `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	MonetaryAmount4         string `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	MonetaryAmount5         string `index:"09" json:"09,omitempty" xml:"09,omitempty"`
	MonetaryAmount6         string `index:"10" json:"10,omitempty" xml:"10,omitempty"`
	MonetaryAmount7         string `index:"11" json:"11,omitempty" xml:"11,omitempty"`
	MonetaryAmount8         string `index:"12" json:"12,omitempty" xml:"12,omitempty"`
	MonetaryAmount9         string `index:"13" json:"13,omitempty" xml:"13,omitempty"`
	MonetaryAmount10        string `index:"14" json:"14,omitempty" xml:"14,omitempty"`
	MonetaryAmount11        string `index:"15" json:"15,omitempty" xml:"15,omitempty"`
	MonetaryAmount12        string `index:"16" json:"16,omitempty" xml:"16,omitempty"`
	MonetaryAmount13        string `index:"17" json:"17,omitempty" xml:"17,omitempty"`
	MonetaryAmount14        string `index:"18" json:"18,omitempty" xml:"18,omitempty"`
	MonetaryAmount15        string `index:"19" json:"19,omitempty" xml:"19,omitempty"`
	MonetaryAmount16        string `index:"20" json:"20,omitempty" xml:"20,omitempty"`
	MonetaryAmount17        string `index:"21" json:"21,omitempty" xml:"21,omitempty"`
	MonetaryAmount18        string `index:"22" json:"22,omitempty" xml:"22,omitempty"`
	Quantity2               string `index:"23" json:"23,omitempty" xml:"23,omitempty"`
	MonetaryAmount19        string `index:"24" json:"24,omitempty" xml:"24,omitempty"`

	Element
}

func (r TS3) defaultMask(index int) string {
	if index < 6 {
		return rules.MASK_REQUIRED
	}
	return rules.MASK_OPTIONAL
}

func (r TS3) fieldCount() int {
	return 24
}

func (r TS3) Name() string {
	return "TS3"
}

func (r *TS3) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r TS3) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *TS3) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("ts3's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *TS3) Parse(data string, args ...string) (int, error) {
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

func (r TS3) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
