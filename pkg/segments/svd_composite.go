// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"fmt"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

type ProcedureIdentifier struct {
	ServiceIdQualifier string `index:"01" json:"01" xml:"01"`
	ServiceId          string `index:"02" json:"02" xml:"02"`
	Modifier1          string `index:"03" json:"03,omitempty" xml:"03,omitempty"`
	Modifier2          string `index:"04" json:"04,omitempty" xml:"04,omitempty"`
	Modifier3          string `index:"05" json:"05,omitempty" xml:"05,omitempty"`
	Modifier4          string `index:"06" json:"06,omitempty" xml:"06,omitempty"`
	Description        string `index:"07" json:"07,omitempty" xml:"07,omitempty"`

	Element
}

func (r *ProcedureIdentifier) defaultMask(index int) string {
	mask := rules.MASK_REQUIRED
	if index > 2 {
		mask = rules.MASK_OPTIONAL
	}
	return mask
}

func (r *ProcedureIdentifier) SetFieldByIndex(index string, data any) error {
	return util.SetFieldByIndex(r, index, data)
}

func (r ProcedureIdentifier) GetFieldByIndex(index string) any {
	return util.GetFieldByIndex(r, index)
}

func (r *ProcedureIdentifier) Validate(rule *rules.ElementSetRule) error {

	if rule == nil {
		rule = r.GetRule()
	}

	for i := 1; i <= 7; i++ {

		idx := fmt.Sprintf("%02d", i)

		if err := util.ValidateField(r.GetFieldByIndex(idx), rule.Get(idx), r.defaultMask(i)); err != nil {
			return fmt.Errorf("procedure identifier's element (%s) has invalid value, %s", idx, err.Error())
		}
	}

	return nil
}

func (r *ProcedureIdentifier) Parse(data string, args ...string) (int, error) {

	var err error
	var size, read int
	line := data

	for i := 1; i <= 7; i++ {

		var value string
		idx := fmt.Sprintf("%02d", i)

		if value, size, err = util.ReadCompositeField(line, read, r.GetRule().Get(idx), r.defaultMask(i), args...); err != nil {
			return 0, fmt.Errorf("unable to parse procedure identifier's element (%s), %s", idx, err.Error())
		} else {
			read += size
			r.SetFieldByIndex(idx, value)
		}
	}

	return read, nil
}

func (r *ProcedureIdentifier) String(args ...string) string {
	var buf string

	separator := util.SubElementSeparator
	if len(args) > 0 {
		separator = args[0]
	}

	for i := 7; i > 0; i-- {

		idx := fmt.Sprintf("%02d", i)
		value := r.GetFieldByIndex(idx)

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
			buf = fmt.Sprintf("%s", value)
		} else {
			buf = fmt.Sprintf("%v%s", value, separator) + buf
		}
	}

	return buf
}
