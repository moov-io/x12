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

func NewSVD(rule *rules.ElementSetRule) SegmentInterface {

	newSegment := SVD{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type SVD struct {
	IdentifierCode      string              `index:"01" json:"01" xml:"01"`
	MonetaryAmount      string              `index:"02" json:"02" xml:"02"`
	ProcedureIdentifier ProcedureIdentifier `index:"03" json:"03" xml:"03"`
	ServiceId           string              `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Quantity            string              `index:"05" json:"05" xml:"05"`
	AssignedNumber      string              `index:"06" json:"06,omitempty" xml:"06,omitempty"`

	Element
}

func (r SVD) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index == 4 || index == 6 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r SVD) fieldCount() int {
	return 6
}

func (r SVD) Name() string {
	return "SVD"
}

func (r *SVD) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r SVD) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *SVD) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var err error
		idx := fmt.Sprintf("%02d", i)

		if i == 3 {
			cRule := rule.Get(idx).Composite
			err = r.ProcedureIdentifier.Validate(&cRule)
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i))
		}

		if err != nil {
			return fmt.Errorf("svd's element (%s) has invalid value, %s", idx, err.Error())
		}

	}

	return nil
}

func (r *SVD) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data, args...)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("svd segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("svd segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		rule := r.GetRule().Get(idx)

		if value, size, err = util.ReadField(line, read, rule, r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse svd's element (%s), %s", idx, err.Error())
		} else {
			read += size

			compositeRule := rule.Composite

			if i == 3 { // HealthCareServiceLocation
				var composite ProcedureIdentifier
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.ProcedureIdentifier = composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse svd's element (%s), %s", idx, parseErr.Error())
				}
			} else {
				r.SetFieldByIndex(idx, value)
			}

		}
	}

	return read, nil
}

func (r SVD) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {

		var value any
		idx := fmt.Sprintf("%02d", i)

		if i == 3 { // HealthCareServiceLocation
			value = r.ProcedureIdentifier.String(args...)
		} else {
			value = r.GetFieldByIndex(idx)
		}

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
