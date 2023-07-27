// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/moov-io/x12/pkg/rules"
)

const (
	DataElementSeparator = "*"
	SubElementSeparator  = ":"
	RepetitionSeparator  = "^"
	SegmentTerminator    = "~"
)

type ElementInfo struct {
	Name  string
	Level int
	Items []string
}

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

func DumpStructInfo(r any, level int) ElementInfo {
	info := ElementInfo{
		Level: level,
	}

	var dataStruct reflect.Value
	if reflect.ValueOf(r).Kind() == reflect.Ptr {
		dataStruct = reflect.ValueOf(r).Elem()
		info.Name = reflect.TypeOf(r).Elem().Name()
	} else if reflect.ValueOf(r).Kind() == reflect.Struct {
		dataStruct = reflect.ValueOf(r)
		info.Name = reflect.TypeOf(r).Name()
	}

	for i := 0; i < dataStruct.NumField(); i++ {
		field := dataStruct.Type().Field(i)

		if len(field.Tag.Get("index")) > 0 {
			elm := dataStruct.Field(i)
			if elm.Kind() == reflect.String {
				inf := elm.Interface()
				if v, ok := inf.(string); ok {
					info.Items = append(info.Items, v)
				}
			} else if elm.Kind() == reflect.Struct {
				method := elm.MethodByName("String")
				if method.IsValid() {
					response := method.Call(nil)
					info.Items = append(info.Items, response[0].String())
				}
			} else if elm.Kind() == reflect.Ptr && !elm.IsNil() {
				method := reflect.ValueOf(elm.Interface()).MethodByName("String")
				if method.IsValid() {
					response := method.Call(nil)
					info.Items = append(info.Items, response[0].String())
				}
			}
		}
	}

	return info
}

func getIndex(input string, args ...string) int {
	idx1 := strings.Index(input, DataElementSeparator)
	idx2 := strings.Index(input, GetSegmentTerminator(args...))

	if idx1 == -1 {
		return idx2
	}

	if idx2 > -1 && idx2 < idx1 {
		return idx2
	}

	return idx1
}

func GetRecordSize(line string, args ...string) int {
	size := strings.Index(line, GetSegmentTerminator(args...))
	if size >= 0 {
		return size + 1
	}

	return size
}

func ValidateField(data any, spec rules.ElementRule, mask string) error {
	mask = rules.GetMask(spec.Mask, mask)
	dataStr := strings.TrimSpace(fmt.Sprintf("%v", data))

	switch mask {
	/*
		case rules.MASK_REQUIRED:
			if dataStr == "" {
				// TODO
				return errors.New("the element is require")
			}
	*/

	case rules.MASK_NOTUSED:
		// TODO
		return nil

	case rules.MASK_OPTIONAL:
		if dataStr == "" {
			return nil
		}
	}

	if len(spec.AcceptRegex) > 0 {
		re, _ := regexp.Compile(spec.AcceptRegex)
		if !re.MatchString(dataStr) {
			return errors.New("the element contains unexpected value")
		}
	}

	if len(spec.AcceptValues) > 0 {
		for _, value := range spec.AcceptValues {
			if dataStr == strings.TrimSpace(value) {
				return nil
			}
		}
		return errors.New("the element contains unexpected value")
	}

	return nil
}

func ReadCompositeField(input string, start int, spec rules.ElementRule, mask string, args ...string) (string, int, error) {
	data := ""
	separator := GetElementSeparator(args...)

	if start < len(input) {
		data = input[start:]
	}

	if data == "" {
		if rules.GetMask(spec.Mask, mask) == rules.MASK_NOTUSED || rules.GetMask(spec.Mask, mask) == rules.MASK_OPTIONAL {
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

func ReadField(input string, start int, spec rules.ElementRule, defaultMask string, args ...string) (string, int, error) {
	data := ""
	mask := rules.GetMask(spec.Mask, defaultMask)

	if start < len(input) {
		data = input[start:]
	}

	if data == "" {
		if mask == rules.MASK_NOTUSED || mask == rules.MASK_OPTIONAL {
			return "", 0, nil
		}
		return "", 0, fmt.Errorf("doesn't enough input string")
	}

	idx := getIndex(data, args...)
	if idx == -1 {
		return "", 0, fmt.Errorf("doesn't have valid delimiter")
	}

	value := data[:idx]
	if err := ValidateField(value, spec, mask); err != nil {
		return "", 0, err
	}

	return value, idx + 1, nil
}

func GetSegmentTerminator(args ...string) string {
	if len(args) < 1 {
		return SegmentTerminator
	}

	terminator := args[0]
	if terminator == "" {
		return SegmentTerminator
	}

	return terminator
}

func GetElementSeparator(args ...string) string {
	if len(args) < 2 {
		return SubElementSeparator
	}

	separator := args[1]
	if len(separator) != 1 {
		return SubElementSeparator
	}

	return separator
}

func GetDuplicateControlNumber(strSlice []string) (bool, string) {
	keys := make(map[string]bool)
	for _, item := range strSlice {
		if _, value := keys[item]; !value {
			keys[item] = true
		} else {
			return true, item
		}
	}
	return false, ""
}
