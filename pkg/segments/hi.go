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

func (r HI) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index > 1 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r HI) fieldCount() int {
	return 8
}

func (r HI) Name() string {
	return "HI"
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

	for i := 1; i <= r.fieldCount(); i++ {

		var err error
		idx := fmt.Sprintf("%02d", i)

		switch i {
		case 1:
			err = r.HealthCareCodeInformation1.Validate(rule)
		case 2:
			if r.HealthCareCodeInformation2 != nil {
				err = r.HealthCareCodeInformation2.Validate(rule)
			}
		case 3:
			if r.HealthCareCodeInformation3 != nil {
				err = r.HealthCareCodeInformation3.Validate(rule)
			}
		case 4:
			if r.HealthCareCodeInformation4 != nil {
				err = r.HealthCareCodeInformation4.Validate(rule)
			}
		case 5:
			if r.HealthCareCodeInformation5 != nil {
				err = r.HealthCareCodeInformation5.Validate(rule)
			}
		case 6:
			if r.HealthCareCodeInformation6 != nil {
				err = r.HealthCareCodeInformation6.Validate(rule)
			}
		case 7:
			if r.HealthCareCodeInformation7 != nil {
				err = r.HealthCareCodeInformation7.Validate(rule)
			}
		case 8:
			if r.HealthCareCodeInformation8 != nil {
				err = r.HealthCareCodeInformation8.Validate(rule)
			}
		}

		if err != nil {
			return fmt.Errorf("hi's element (%s) has invalid value, %s", idx, err.Error())
		}

	}

	return nil
}

func (r *HI) Parse(data string, args ...string) (int, error) {

	var line string
	var err error
	var size int

	length := util.GetRecordSize(data)
	codeLen := len(r.Name())
	read := codeLen + 1

	if length < int64(read) {
		return 0, errors.New("hi segment has not enough input data")
	} else {
		line = data[:length]
	}

	if r.Name() != data[:codeLen] {
		return 0, errors.New("hi segment contains invalid code")
	}

	for i := 1; i <= r.fieldCount(); i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		rule := r.GetRule().Get(idx)

		if value, size, err = util.ReadField(line, read, rule, r.defaultMask(i)); err != nil {
			return 0, fmt.Errorf("unable to parse hi's element (%s), %s", idx, err.Error())
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

				if rules.IsMaskRequired(util.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse hi's element (%s), %s", idx, parseErr.Error())
				}
			case 2:
				if parseErr == nil {
					r.HealthCareCodeInformation2 = &composite
				}
			case 3:
				if parseErr == nil {
					r.HealthCareCodeInformation3 = &composite
				}
			case 4:
				if parseErr == nil {
					r.HealthCareCodeInformation4 = &composite
				}
			case 5:
				if parseErr == nil {
					r.HealthCareCodeInformation5 = &composite
				}
			case 6:
				if parseErr == nil {
					r.HealthCareCodeInformation6 = &composite
				}
			case 7:
				if parseErr == nil {
					r.HealthCareCodeInformation7 = &composite
				}
			case 8:
				if parseErr == nil {
					r.HealthCareCodeInformation8 = &composite
				}
			}

		}
	}

	return read, nil
}

func (r HI) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {

		var value any
		idx := fmt.Sprintf("%02d", i)

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
			buf = fmt.Sprintf("%v%s", value, util.SegmentTerminator)
		} else {
			buf = fmt.Sprintf("%v%s", value, util.DataElementSeparator) + buf
		}
	}

	if buf == "" {
		buf = fmt.Sprintf("%s%s", r.Name(), util.SegmentTerminator)
	} else {
		buf = fmt.Sprintf("%s%s", r.Name(), util.DataElementSeparator) + buf
	}

	return buf
}
