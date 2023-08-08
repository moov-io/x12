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
	ErrMinSegment       = 13
	ErrAllSegment       = 14
	ErrMismatchLoop     = 15
	ErrAllLoop          = 16
	ErrMinLoop          = 17
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
		Reason:     strings.ToLower(fmt.Sprintf("unable to parse %s's element (%s), %s", strings.ToLower(name), idx, msg)),
		Originated: []string{name},
	}

	return &err
}

func NewValidateElementError(name, idx, msg string) error {
	err := ErrorObject{
		Code:       ErrValidateSegment,
		Reason:     strings.ToLower(fmt.Sprintf("%s's element (%s) has invalid value, %s", name, idx, msg)),
		Originated: []string{name},
	}

	return &err
}

func NewMinLengthErr(name string) error {
	err := ErrorObject{
		Code:       ErrMinLength,
		Reason:     strings.ToLower(fmt.Sprintf("%s segment has not enough input data", name)),
		Originated: []string{name},
	}

	return &err
}

func NewMaxLengthErr(name string) error {
	err := ErrorObject{
		Code:       ErrMaxLength,
		Reason:     strings.ToLower(fmt.Sprintf("%s segment can't parse all input data", name)),
		Originated: []string{name},
	}

	return &err
}

func NewInvalidCodeErr(name string) error {
	err := ErrorObject{
		Code:       ErrInvalidCode,
		Reason:     strings.ToLower(fmt.Sprintf("%s segment contains invalid code", name)),
		Originated: []string{name},
	}

	return &err
}

func NewUnsupportSegmentError(name string) error {
	err := ErrorObject{
		Code:       ErrUnsupportSegment,
		Reason:     strings.ToLower(fmt.Sprintf("unsupported segment name(%s)", name)),
		Originated: []string{name},
	}

	return &err
}

func NewFindSegmentError(name string) error {
	err := ErrorObject{
		Code:       ErrFindSegment,
		Reason:     strings.ToLower(fmt.Sprintf("unable to find %s segment", name)),
		Originated: []string{name},
	}

	return &err
}

func NewFindRuleError(name string) error {
	err := ErrorObject{
		Code:       ErrFindRule,
		Reason:     strings.ToLower(fmt.Sprintf("unable to find %s rule", name)),
		Originated: []string{name},
	}

	return &err
}

func NewFindElementError(name string) error {
	err := ErrorObject{
		Code:       ErrFindElement,
		Reason:     strings.ToLower(fmt.Sprintf("unable to find %s", name)),
		Originated: []string{name},
	}

	return &err
}

func NewSpecifiedRuleError(name string) error {
	err := ErrorObject{
		Code:       ErrSpecifiedRule,
		Reason:     strings.ToLower(fmt.Sprintf("element(%s) rule is not defined", name)),
		Originated: []string{name},
	}

	return &err
}

func NewMismatchSegmentError(segName, ruleName string) error {
	err := ErrorObject{
		Code:       ErrMismatchSegment,
		Reason:     strings.ToLower(fmt.Sprintf("segment(%s) don't accept specified rule(%s), please verify segment orders or has dirty segments as previous segment", segName, ruleName)),
		Originated: []string{segName},
	}

	return &err
}

func NewUnableParseError(name string) error {
	err := ErrorObject{
		Code:       ErrUnableParse,
		Reason:     strings.ToLower(fmt.Sprintf("unable to parse %s", name)),
		Originated: []string{name},
	}

	return &err
}

func NewMinSegmentError(name string) error {
	err := ErrorObject{
		Code:       ErrMinSegment,
		Reason:     strings.ToLower(fmt.Sprintf("segment(%s) does not repeat as specified times", name)),
		Originated: []string{name},
	}

	return &err
}

func NewAllSegmentError(name string) error {
	err := ErrorObject{
		Code:       ErrAllSegment,
		Reason:     strings.ToLower(fmt.Sprintf("all segments of %s don't to validate using specified rule, have dirty segments", name)),
		Originated: []string{name},
	}

	return &err
}

func NewMismatchLoopError(segName, ruleName string) error {
	err := ErrorObject{
		Code:       ErrMismatchLoop,
		Reason:     strings.ToLower(fmt.Sprintf("loop(%s) don't accept specified rule(%s), please verify loop orders or has dirty loops as previous loop", segName, ruleName)),
		Originated: []string{segName},
	}

	return &err
}

func NewAllLoopError(name string) error {
	err := ErrorObject{
		Code:       ErrAllLoop,
		Reason:     strings.ToLower(fmt.Sprintf("all loops of %s don't to validate using specified rule, have dirty rules", name)),
		Originated: []string{name},
	}

	return &err
}

func NewMinLoopError(name string) error {
	err := ErrorObject{
		Code:       ErrMinLoop,
		Reason:     strings.ToLower(fmt.Sprintf("loop(%s) does not repeat as specified times", name)),
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
}

func UpdateErrorReason(err error) error {
	buildStack := func(stack []string) string {
		msg := ""
		for _, item := range stack {
			item = strings.ToLower(item)
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
		if len(obj.Segment) > 0 {
			obj.Reason = fmt.Sprintf("%s, error segment line: '%s', error stack: '%s'", obj.Reason, buildStack(obj.Originated), obj.Segment)
		} else {
			obj.Reason = fmt.Sprintf("%s, error stack: '%s'", obj.Reason, buildStack(obj.Originated))
		}
	} else {
		obj.Reason = fmt.Sprintf("%s, error segment line: '%s'", obj.Reason, obj.Segment)
	}

	return err
}
