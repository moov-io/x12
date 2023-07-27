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

func NewHL(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := HL{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type HL struct {
	HierarchicalIdNumber       string `index:"01" json:"01" xml:"01"`
	HierarchicalParentIdNumber string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	HierarchicalLevelCode      string `index:"03" json:"03" xml:"03"`
	HierarchicalChildCode      string `index:"04" json:"04" xml:"04"`

	Element
}

func (r *HL) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index == 2 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r HL) fieldCount() int {
	return 4
}

func (r HL) Name() string {
	return "HL"
}

func (r *HL) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r HL) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *HL) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("hl's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *HL) Parse(data string, args ...string) (int, error) {
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

func (r HL) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
