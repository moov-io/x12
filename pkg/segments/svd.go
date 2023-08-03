// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"strings"

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

	for i := 1; i <= segmentFieldCount(r); i++ {
		var err error
		idx := util.GetFormattedIndex(i)

		if i == 3 {
			cRule := rule.Get(idx).Composite
			err = r.ProcedureIdentifier.Validate(&cRule)
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), getFieldMask(r, i))
		}

		if err != nil {
			return util.NewValidateElementError(util.GetStructName(r), idx, err.Error())
		}

	}

	return nil
}

func (r *SVD) Parse(data string, args ...string) (int, error) {
	var size int
	name := strings.ToLower(util.GetStructName(r))
	read, line, err := r.VerifyCode(data, name, args...)
	if err != nil {
		return 0, err
	}

	for i := 1; i <= segmentFieldCount(r); i++ {
		var value string
		idx := util.GetFormattedIndex(i)

		rule := r.GetRule().Get(idx)

		if value, size, err = util.ReadField(line, read, rule, getFieldMask(r, i), args...); err != nil {
			return 0, util.NewParseSegmentError(name, idx, err.Error())
		} else {
			read += size

			compositeRule := rule.Composite

			if i == 3 {
				var composite ProcedureIdentifier
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.ProcedureIdentifier = composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			} else {
				r.SetFieldByIndex(idx, value)
			}

		}
	}

	return returnRead(read, data, name)
}

func (r SVD) String(args ...string) string {
	var buf string

	for i := segmentFieldCount(r); i > 0; i-- {
		var value any
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		if i == 3 {
			value = r.ProcedureIdentifier.String(args...)
		} else {
			value = r.GetFieldByIndex(idx)
		}

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), value)
	}

	return r.TerminateString(buf, util.GetStructName(r))
}
