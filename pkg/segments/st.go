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

func NewST(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := ST{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type ST struct {
	TransactionSetIdentifierCode      string `index:"01" json:"01" xml:"01"`
	TransactionSetControlNumber       string `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	ImplementationConventionReference string `index:"03" json:"03,omitempty" xml:"03,omitempty"`

	Element
}

func (r ST) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index >= 2 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r ST) fieldCount() int {
	return 3
}

func (r ST) Name() string {
	return "ST"
}

func (r *ST) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r ST) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *ST) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		idx := fmt.Sprintf("%02d", i)
		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("st's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *ST) Parse(data string, args ...string) (int, error) {
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

func (r ST) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), r.GetFieldByIndex(idx))
	}

	return r.TerminateString(buf, r.Name())
}
