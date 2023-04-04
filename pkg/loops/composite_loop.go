// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package loops

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/moov-io/x12/pkg/rules"
)

func NewCompositeLoop(rule *rules.LoopRule) *CompositeLoop {

	newLoop := CompositeLoop{rule: rule}

	return &newLoop
}

type CompositeLoop struct {
	Loop     Loop
	SubLoops []CompositeLoop

	rule *rules.LoopRule
}

func (r CompositeLoop) Name() string {
	if r.rule != nil {
		return r.rule.Name
	}
	return "Composite loop"
}

func (r CompositeLoop) GetRule() *rules.LoopRule {
	return r.rule
}

func (r *CompositeLoop) SetRule(s *rules.LoopRule) {
	r.rule = s
}

func (r *CompositeLoop) Validate(loopRule *rules.LoopRule) error {

	if loopRule == nil && r.rule != nil {
		loopRule = r.rule
	}

	if loopRule == nil {
		return errors.New("please specify rules for this Composite loop")
	}

	err := r.Loop.Validate(&loopRule.Segments)
	if err != nil {
		return fmt.Errorf("loop(%s) is invalid, %s", r.rule.Name, err.Error())
	}

	segIndex := 0
	for index := 0; index < len(loopRule.Composite); index++ {
		rule := loopRule.Composite[index]
		for repeatCnt := 0; repeatCnt < rule.Repeat(); repeatCnt++ {

			if segIndex+1 > len(r.SubLoops) {
				if repeatCnt == 0 && rules.IsMaskRequired(rule.Mask) {
					return fmt.Errorf("please add new %s loop", strings.ToUpper(rule.Name))
				}
				continue
			}

			if r.SubLoops[segIndex].Name() != rule.Name {
				if rules.IsMaskRequired(rule.Mask) {
					return fmt.Errorf("loop(%02d)'s name is not equal with rule's name (%s)", segIndex, strings.ToLower(rule.Name))
				}
				continue
			}

			if err = r.SubLoops[segIndex].Validate(&rule); err != nil {
				if rules.IsMaskRequired(rule.Mask) {
					return fmt.Errorf("loop(%02d) should be valid %s loop, %s", segIndex, strings.ToLower(rule.Name), err.Error())
				}
				continue
			}

			segIndex++
		}
	}

	if len(r.SubLoops) > segIndex {
		if segIndex == len(r.SubLoops)-1 {
			return fmt.Errorf("unable to validate loop(%02d), rule is not specified", segIndex)
		} else {
			return fmt.Errorf("unable to validate loop(%02d~%02d), rule is not specified", segIndex, len(r.SubLoops)-1)
		}
	}

	return nil
}

func (r *CompositeLoop) Parse(data string, args ...string) (int, error) {

	if r.rule == nil {
		return 0, errors.New("please specify rules for this Composite loop")
	}

	r.Loop.SetRule(r.rule)
	size, err := r.Loop.Parse(data, args...)
	if err != nil {
		return 0, fmt.Errorf("unable to parse %s loop", strings.ToLower(r.rule.Name))
	}

	if len(r.rule.Composite) == 0 {
		return size, nil
	}

	read := size
	line := data[read:]

	for index := 0; index < len(r.rule.Composite); index++ {
		rule := r.rule.Composite[index]

		for repeatIdx := 0; repeatIdx < rule.Repeat(); repeatIdx++ {
			newChild := NewCompositeLoop(&rule)
			size, err = newChild.Parse(line, args...)
			if err == nil {
				read += size
				line = data[read:]
				r.SubLoops = append(r.SubLoops, *newChild)
			} else {
				if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
					return 0, fmt.Errorf("unable to parse %s loop (%s)", strings.ToLower(rule.Name), err.Error())
				}
				break
			}
		}

	}

	return read, nil
}

func (r CompositeLoop) String(args ...string) string {
	var buf bytes.Buffer

	buf.WriteString(r.Loop.String(args...))
	for index := range r.SubLoops {
		buf.WriteString(r.SubLoops[index].String(args...))
	}

	return buf.String()
}
