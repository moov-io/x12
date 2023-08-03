// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"strings"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

func NewHI(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := HI{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type HI struct {
	HealthCareCodeInformation1 HealthCareCode  `index:"01" json:"01" xml:"01"`
	HealthCareCodeInformation2 *HealthCareCode `index:"02" json:"02,omitempty" xml:"02,omitempty"`
	HealthCareCodeInformation3 *HealthCareCode `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	HealthCareCodeInformation4 *HealthCareCode `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	HealthCareCodeInformation5 *HealthCareCode `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	HealthCareCodeInformation6 *HealthCareCode `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	HealthCareCodeInformation7 *HealthCareCode `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	HealthCareCodeInformation8 *HealthCareCode `index:"08" json:"08,omitempty" xml:"08,omitempty"`

	Element
}

func (r *HI) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r HI) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *HI) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= segmentFieldCount(r); i++ {
		var err error
		idx := util.GetFormattedIndex(i)

		switch i {
		case 1:
			cRule := rule.Get(idx).Composite
			err = r.HealthCareCodeInformation1.Validate(&cRule)
		case 2:
			if r.HealthCareCodeInformation2 != nil {
				cRule := rule.Get(idx).Composite
				err = r.HealthCareCodeInformation2.Validate(&cRule)
			}
		case 3:
			if r.HealthCareCodeInformation3 != nil {
				cRule := rule.Get(idx).Composite
				err = r.HealthCareCodeInformation3.Validate(&cRule)
			}
		case 4:
			if r.HealthCareCodeInformation4 != nil {
				cRule := rule.Get(idx).Composite
				err = r.HealthCareCodeInformation4.Validate(&cRule)
			}
		case 5:
			if r.HealthCareCodeInformation5 != nil {
				cRule := rule.Get(idx).Composite
				err = r.HealthCareCodeInformation5.Validate(&cRule)
			}
		case 6:
			if r.HealthCareCodeInformation6 != nil {
				cRule := rule.Get(idx).Composite
				err = r.HealthCareCodeInformation6.Validate(&cRule)
			}
		case 7:
			if r.HealthCareCodeInformation7 != nil {
				cRule := rule.Get(idx).Composite
				err = r.HealthCareCodeInformation7.Validate(&cRule)
			}
		case 8:
			if r.HealthCareCodeInformation8 != nil {
				cRule := rule.Get(idx).Composite
				err = r.HealthCareCodeInformation8.Validate(&cRule)
			}
		}

		if err != nil {
			return util.NewValidateElementError(util.GetStructName(r), idx, err.Error())
		}

	}

	return nil
}

func (r *HI) Parse(data string, args ...string) (int, error) {
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
			var composite HealthCareCode
			if compositeRule != nil {
				composite.SetRule(&compositeRule)
			}

			_, parseErr := composite.Parse(value, args...)
			switch i {
			case 1:
				if parseErr == nil {
					r.HealthCareCodeInformation1 = composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			case 2:
				if parseErr == nil {
					r.HealthCareCodeInformation2 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			case 3:
				if parseErr == nil {
					r.HealthCareCodeInformation3 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			case 4:
				if parseErr == nil {
					r.HealthCareCodeInformation4 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			case 5:
				if parseErr == nil {
					r.HealthCareCodeInformation5 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			case 6:
				if parseErr == nil {
					r.HealthCareCodeInformation6 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			case 7:
				if parseErr == nil {
					r.HealthCareCodeInformation7 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			case 8:
				if parseErr == nil {
					r.HealthCareCodeInformation8 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, getFieldMask(r, i))) && parseErr != nil {
					return 0, util.NewParseSegmentError(name, idx, parseErr.Error())
				}
			}

		}
	}

	return returnRead(read, data, name, args...)
}

func (r HI) String(args ...string) string {
	var buf string

	for i := segmentFieldCount(r); i > 0; i-- {
		var value any
		idx := util.GetFormattedIndex(i)
		mask := r.GetRule().GetMask(idx, getFieldMask(r, i))

		switch i {
		case 1:
			value = r.HealthCareCodeInformation1.String(args...)
		case 2:
			if r.HealthCareCodeInformation2 != nil {
				value = r.HealthCareCodeInformation2.String(args...)
			}
		case 3:
			if r.HealthCareCodeInformation3 != nil {
				value = r.HealthCareCodeInformation3.String(args...)
			}
		case 4:
			if r.HealthCareCodeInformation4 != nil {
				value = r.HealthCareCodeInformation4.String(args...)
			}
		case 5:
			if r.HealthCareCodeInformation5 != nil {
				value = r.HealthCareCodeInformation5.String(args...)
			}
		case 6:
			if r.HealthCareCodeInformation6 != nil {
				value = r.HealthCareCodeInformation6.String(args...)
			}
		case 7:
			if r.HealthCareCodeInformation7 != nil {
				value = r.HealthCareCodeInformation7.String(args...)
			}
		case 8:
			if r.HealthCareCodeInformation8 != nil {
				value = r.HealthCareCodeInformation8.String(args...)
			}
		}

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), value)
	}

	return r.TerminateString(buf, util.GetStructName(r))
}
