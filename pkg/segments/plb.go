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

func NewPLB(rule *rules.ElementSetRule) SegmentInterface {
	newSegment := PLB{}

	if rule == nil {
		newRule := make(rules.ElementSetRule)
		newSegment.SetRule(&newRule)
	} else {
		newSegment.SetRule(rule)
	}

	return &newSegment
}

type PLB struct {
	ReferenceIdentification string                  `index:"01" json:"01" xml:"01"`
	Date                    string                  `index:"02" json:"02" xml:"02"`
	MonetaryIdentification1 MonetaryIdentification  `index:"03" json:"03" xml:"03"`
	MonetaryAmount1         string                  `index:"04" json:"04" xml:"04"`
	MonetaryIdentification2 *MonetaryIdentification `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	MonetaryAmount2         string                  `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	MonetaryIdentification3 *MonetaryIdentification `index:"07" json:"07,omitempty" xml:"07,omitempty"`
	MonetaryAmount3         string                  `index:"08" json:"08,omitempty" xml:"08,omitempty"`
	MonetaryIdentification4 *MonetaryIdentification `index:"09" json:"09,omitempty" xml:"09,omitempty"`
	MonetaryAmount4         string                  `index:"10" json:"10,omitempty" xml:"10,omitempty"`
	MonetaryIdentification5 *MonetaryIdentification `index:"11" json:"11,omitempty" xml:"11,omitempty"`
	MonetaryAmount5         string                  `index:"12" json:"12,omitempty" xml:"12,omitempty"`
	MonetaryIdentification6 *MonetaryIdentification `index:"13" json:"13,omitempty" xml:"13,omitempty"`
	MonetaryAmount6         string                  `index:"14" json:"14,omitempty" xml:"14,omitempty"`

	Element
}

func (r PLB) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index > 4 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r PLB) fieldCount() int {
	return 14
}

func (r PLB) Name() string {
	return "PLB"
}

func (r *PLB) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r PLB) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *PLB) Validate(rule *rules.ElementSetRule) error {
	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= r.fieldCount(); i++ {
		var err error
		idx := fmt.Sprintf("%02d", i)

		if i == 3 {
			cRule := rule.Get(idx).Composite
			err = r.MonetaryIdentification1.Validate(&cRule)
		} else if i == 5 {
			if r.MonetaryIdentification2 != nil {
				cRule := rule.Get(idx).Composite
				err = r.MonetaryIdentification2.Validate(&cRule)
			}
		} else if i == 7 {
			if r.MonetaryIdentification3 != nil {
				cRule := rule.Get(idx).Composite
				err = r.MonetaryIdentification3.Validate(&cRule)
			}
		} else if i == 9 {
			if r.MonetaryIdentification4 != nil {
				cRule := rule.Get(idx).Composite
				err = r.MonetaryIdentification4.Validate(&cRule)
			}
		} else if i == 11 {
			if r.MonetaryIdentification5 != nil {
				cRule := rule.Get(idx).Composite
				err = r.MonetaryIdentification5.Validate(&cRule)
			}
		} else if i == 13 {
			if r.MonetaryIdentification6 != nil {
				cRule := rule.Get(idx).Composite
				err = r.MonetaryIdentification6.Validate(&cRule)
			}
		} else {
			err = util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i))
		}

		if err != nil {
			return fmt.Errorf("plb's element (%s) has invalid value, %s", idx, err.Error())
		}

	}

	return nil
}

func (r *PLB) Parse(data string, args ...string) (int, error) {
	var size int
	name := strings.ToLower(r.Name())
	read, line, err := r.VerifyCode(data, name, args...)
	if err != nil {
		return 0, err
	}

	for i := 1; i <= r.fieldCount(); i++ {
		var value string
		idx := fmt.Sprintf("%02d", i)
		rule := r.GetRule().Get(idx)

		if value, size, err = util.ReadField(line, read, rule, r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, err.Error())
		} else {
			read += size
			compositeRule := rule.Composite

			if i == 3 {
				var composite MonetaryIdentification
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.MonetaryIdentification1 = composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}
			} else if i == 5 {
				var composite MonetaryIdentification
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.MonetaryIdentification2 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}
			} else if i == 7 {
				var composite MonetaryIdentification
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.MonetaryIdentification3 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}
			} else if i == 9 {
				var composite MonetaryIdentification
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.MonetaryIdentification4 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}
			} else if i == 11 {
				var composite MonetaryIdentification
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.MonetaryIdentification5 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}
			} else if i == 13 {
				var composite MonetaryIdentification
				if compositeRule != nil {
					composite.SetRule(&compositeRule)
				}

				_, parseErr := composite.Parse(value, args...)
				if parseErr == nil {
					r.MonetaryIdentification6 = &composite
				}

				if rules.IsMaskRequired(rules.GetMask(rule.Mask, r.defaultMask(i))) && parseErr != nil {
					return 0, fmt.Errorf("unable to parse %s's element (%s), %s", name, idx, parseErr.Error())
				}
			} else {
				r.SetFieldByIndex(idx, value)
			}
		}
	}

	return read, nil
}

func (r PLB) String(args ...string) string {
	var buf string

	for i := r.fieldCount(); i > 0; i-- {
		var value any
		idx := fmt.Sprintf("%02d", i)
		mask := r.GetRule().GetMask(idx, r.defaultMask(i))

		if i == 3 {
			value = r.MonetaryIdentification1.String(args...)
		} else if i == 5 {
			if r.MonetaryIdentification2 != nil {
				value = r.MonetaryIdentification2.String(args...)
			}
		} else if i == 7 {
			if r.MonetaryIdentification3 != nil {
				value = r.MonetaryIdentification3.String(args...)
			}
		} else if i == 9 {
			if r.MonetaryIdentification4 != nil {
				value = r.MonetaryIdentification4.String(args...)
			}
		} else if i == 11 {
			if r.MonetaryIdentification5 != nil {
				value = r.MonetaryIdentification5.String(args...)
			}
		} else if i == 13 {
			if r.MonetaryIdentification6 != nil {
				value = r.MonetaryIdentification6.String(args...)
			}
		} else {
			value = r.GetFieldByIndex(idx)
		}

		buf = r.CompositeString(buf, mask, util.DataElementSeparator, util.GetSegmentTerminator(args...), value)
	}

	return r.TerminateString(buf, r.Name())
}
