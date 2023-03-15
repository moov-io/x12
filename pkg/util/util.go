// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"fmt"
	"github.com/moov-io/x12/pkg/rules"
	"reflect"
	"strconv"
	"strings"
)

const (
	DataElementSeparator = "*"
	SubElementSeparator  = ":"
	RepetitionSeparator  = "^"
	SegmentTerminator    = "~"
)

func SetFieldByIndex(r any, index string, data any) error {

	var dataStruct reflect.Value
	if reflect.ValueOf(r).Kind() == reflect.Ptr {
		dataStruct = reflect.ValueOf(r).Elem()
	} else {
		return errors.New("could not set data")
	}

	for i := 0; i < dataStruct.NumField(); i++ {
		field := dataStruct.Type().Field(i)

		if field.Tag.Get("index") == index {

			org := dataStruct.Field(i)
			value := reflect.ValueOf(data)

			if !value.IsValid() {
				value = reflect.Zero(org.Type())
			} else if org.Type() != value.Type() {
				return errors.New("doesn't match setting type")
			}

			org.Set(value)

			return nil
		}
	}

	return errors.New("unable to find matched index")
}

func GetFieldByIndex(r any, index string) any {

	var dataStruct reflect.Value
	if reflect.ValueOf(r).Kind() == reflect.Ptr {
		dataStruct = reflect.ValueOf(r).Elem()
	} else if reflect.ValueOf(r).Kind() == reflect.Struct {
		dataStruct = reflect.ValueOf(r)
	}

	for i := 0; i < dataStruct.NumField(); i++ {
		field := dataStruct.Type().Field(i)

		if field.Tag.Get("index") == index {
			return dataStruct.Field(i).Interface()
		}
	}

	return nil
}

func getIndex(input string) int {

	idx1 := strings.Index(input, DataElementSeparator)
	idx2 := strings.Index(input, SegmentTerminator)

	if idx1 == -1 {
		return idx2
	}

	if idx2 > -1 && idx2 < idx1 {
		return idx2
	}

	return idx1
}

func GetRecordSize(line string) int64 {

	size := strings.Index(line, SegmentTerminator)
	if size >= 0 {
		return int64(size + 1)
	}

	return int64(size)
}

func GetMask(mask, defaultMask string) string {
	if mask == rules.MASK_NONE {
		mask = defaultMask
	}
	return mask
}

func ValidateField(data any, spec rules.ElementRule, mask string) error {

	mask = GetMask(spec.Mask, mask)
	dataStr := strings.TrimSpace(fmt.Sprintf("%v", data))

	switch mask {
	case rules.MASK_REQUIRED:
		if dataStr == "" {
			// TODO
			fmt.Errorf("the element is require")
		}
		break
	case rules.MASK_NOTUSED:
		if dataStr != "" {
			// TODO
			fmt.Errorf("the element is not used")
		}
		break
	case rules.MASK_OPTIONAL:
		if dataStr == "" {
			return nil
		}
		break
	}

	if len(spec.AcceptValues) > 0 {
		for _, value := range spec.AcceptValues {
			if dataStr == strings.TrimSpace(value) {
				return nil
			}
		}
		return fmt.Errorf("the element contains unexpected value")
	}

	return nil
}

func ReadCompositeField(input string, start int, spec rules.ElementRule, mask string, args ...string) (string, int, error) {

	data := ""
	separator := SubElementSeparator
	if len(args) > 0 {
		separator = args[0]
	}

	if start < len(input) {
		data = input[start:]
	}

	if data == "" {
		if GetMask(spec.Mask, mask) == rules.MASK_NOTUSED || GetMask(spec.Mask, mask) == rules.MASK_OPTIONAL {
			return "", 0, nil
		}
		return "", 0, fmt.Errorf("doesn't enough input string")
	}

	isLast := false
	idx := strings.Index(data, separator)
	if idx == -1 {
		idx = len(data)
		isLast = true
	}

	value := data[:idx]
	if err := ValidateField(value, spec, mask); err != nil {
		return "", 0, err
	}

	if isLast {
		return value, idx, nil
	}

	return value, idx + 1, nil
}

func ReadField(input string, start int, spec rules.ElementRule, mask string) (string, int, error) {

	data := ""

	if start < len(input) {
		data = input[start:]
	}

	if data == "" {
		if GetMask(spec.Mask, mask) == rules.MASK_NOTUSED || GetMask(spec.Mask, mask) == rules.MASK_OPTIONAL {
			return "", 0, nil
		}
		return "", 0, fmt.Errorf("doesn't enough input string")
	}

	idx := getIndex(data)
	if idx == -1 {
		return "", 0, fmt.Errorf("doesn't have valid delimiter")
	}

	value := data[:idx]
	if err := ValidateField(value, spec, mask); err != nil {
		return "", 0, err
	}

	return value, idx + 1, nil
}

func ReadFieldAsInt(input string, start int, spec rules.ElementRule, mask string) (int64, int, error) {

	data := ""

	if start < len(input) {
		data = input[start:]
	}

	if data == "" {
		if GetMask(spec.Mask, mask) == rules.MASK_NOTUSED {
			return 0, 0, nil
		}
		return 0, 0, fmt.Errorf("doesn't enough input string")
	}

	idx := getIndex(data)
	if idx == -1 {
		return 0, 0, fmt.Errorf("doesn't have valid delimiter")
	}

	if data[:idx] == "" {
		return 0, 1, nil
	}

	value, err := strconv.ParseInt(data[:idx], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("doesn't have valid value")
	}

	if err = ValidateField(value, spec, mask); err != nil {
		return 0, 0, err
	}

	return value, idx + 1, nil
}
