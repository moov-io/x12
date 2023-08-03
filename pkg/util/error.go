package util

import (
	"fmt"
	"strings"
)

const (
	ErrParseSegment     = 1
	ErrValidateSegment  = 2
	ErrMinLength        = 3
	ErrMaxLength        = 4
	ErrInvalidCode      = 5
	ErrUnsupportSegment = 6
	ErrFindSegment      = 7
	ErrFindElement      = 8
	ErrFindRule         = 9
	ErrSpecifiedRule    = 10
	ErrMismatchSegment  = 11
	ErrUnableParse      = 12
)

type ErrorObject struct {
	Code         int
	Reason       string
	Segment      string
	Originated   []string
	SegmentIndex int
}

func (e *ErrorObject) Error() string {
	return e.Reason
}

func NewParseSegmentError(name, idx, msg string) error {
	err := ErrorObject{
		Code:       ErrParseSegment,
		Reason:     fmt.Sprintf("unable to parse %s's element (%s), %s", strings.ToLower(name), idx, strings.ToLower(msg)),
		Originated: []string{name},
	}

	return &err
}

func NewValidateElementError(name, idx, msg string) error {
	err := ErrorObject{
		Code:       ErrValidateSegment,
		Reason:     fmt.Sprintf("%s's element (%s) has invalid value, %s", strings.ToLower(name), idx, strings.ToLower(msg)),
		Originated: []string{name},
	}

	return &err
}

func NewMinLengthErr(name string) error {
	err := ErrorObject{
		Code:       ErrMinLength,
		Reason:     fmt.Sprintf("%s segment has not enough input data", strings.ToLower(name)),
		Originated: []string{name},
	}

	return &err
}

func NewMaxLengthErr(name string) error {
	err := ErrorObject{
		Code:       ErrMaxLength,
		Reason:     fmt.Sprintf("%s segment can't parse all input data", strings.ToLower(name)),
		Originated: []string{name},
	}

	return &err
}

func NewInvalidCodeErr(name string) error {
	err := ErrorObject{
		Code:       ErrInvalidCode,
		Reason:     fmt.Sprintf("%s segment contains invalid code", strings.ToLower(name)),
		Originated: []string{name},
	}

	return &err
}

func NewUnsupportSegmentError(name string) error {
	err := ErrorObject{
		Code:       ErrUnsupportSegment,
		Reason:     fmt.Sprintf("unsupported segment name(%s)", strings.ToLower(name)),
		Originated: []string{name},
	}

	return &err
}

func NewFindSegmentError(name string) error {
	err := ErrorObject{
		Code:       ErrFindSegment,
		Reason:     fmt.Sprintf("unable to find %s segment", strings.ToLower(name)),
		Originated: []string{name},
	}

	return &err
}

func NewFindRuleError(name string) error {
	err := ErrorObject{
		Code:       ErrFindRule,
		Reason:     fmt.Sprintf("unable to find %s rule", strings.ToLower(name)),
		Originated: []string{name},
	}

	return &err
}

func NewFindElementError(name string) error {
	err := ErrorObject{
		Code:       ErrFindElement,
		Reason:     fmt.Sprintf("unable to find %s", strings.ToLower(name)),
		Originated: []string{name},
	}

	return &err
}

func NewSpecifiedRuleError(name string) error {
	err := ErrorObject{
		Code:   ErrSpecifiedRule,
		Reason: fmt.Sprintf("please specify rules for loop(%s)", strings.ToLower(name)),
	}

	return &err
}

func NewMismatchSegmentError(segName, ruleName string) error {
	err := ErrorObject{
		Code:   ErrMismatchSegment,
		Reason: fmt.Sprintf("segment(%s)'s name is not equal with rule's name (%s)", strings.ToLower(segName), strings.ToLower(ruleName)),
	}

	return &err
}

func NewUnableParseError(name string) error {
	err := ErrorObject{
		Code:       ErrUnableParse,
		Reason:     fmt.Sprintf("unable to parse %s", strings.ToLower(name)),
		Originated: []string{name},
	}

	return &err
}

func AppendErrorSegmentLine(err error, data string, args ...string) {
	if err == nil {
		return
	}

	obj, ok := err.(*ErrorObject)
	if !ok {
		return
	}

	if size := GetRecordSize(data, args...); size > 0 {
		msg := data[:size]
		msg = strings.ReplaceAll(msg, "\n", "\\n")
		msg = strings.ReplaceAll(msg, "\r", "\\r")
		obj.Segment = msg
	}

	return
}

func AppendErrorStack(err error, name string) {
	if err == nil {
		return
	}

	obj, ok := err.(*ErrorObject)
	if !ok {
		return
	}

	obj.Originated = append(obj.Originated, name)
	return
}

func UpdateErrorReason(err error) error {
	buildStack := func(stack []string) string {
		msg := ""
		for _, item := range stack {
			if msg == "" {
				msg = item
			} else {
				msg = item + "->" + msg
			}

		}
		return msg
	}

	if err == nil {
		return err
	}

	obj, ok := err.(*ErrorObject)
	if !ok {
		return err
	}

	if len(obj.Originated) > 1 {
		obj.Reason = fmt.Sprintf("%s, error stack: '%s', error segment line: '%s'", obj.Reason, buildStack(obj.Originated), obj.Segment)
	} else {
		obj.Reason = fmt.Sprintf("%s, error segment line: '%s'", obj.Reason, obj.Segment)
	}

	return err
}
