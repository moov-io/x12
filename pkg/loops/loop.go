// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package loops

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
)

func NewLoop(rule *rules.LoopRule) *Loop {

	newLoop := Loop{
		rule: rule,
	}

	return &newLoop
}

type Loop struct {
	Segments []segments.SegmentInterface

	rule *rules.LoopRule
}

func (r Loop) GetRule() *rules.LoopRule {
	return r.rule
}

func (r *Loop) SetRule(s *rules.LoopRule) {
	r.rule = s
}

func (r Loop) Name() string {
	if r.rule != nil {
		return r.rule.Name
	}
	return "loop"
}

func (r *Loop) Validate(loopRule *rules.SegmentSetRule) error {

	if loopRule == nil && r.rule != nil {
		loopRule = &r.rule.Segments
	}

	if loopRule == nil {
		return errors.New("please specify rules for this loop")
	}

	index := 0
	segIndex := 0
	for rule := loopRule.Get(index); rule != nil; rule = loopRule.Get(index) {

		for repeatIdx := 0; repeatIdx < rule.Repeat(); repeatIdx++ {

			if segIndex+1 > len(r.Segments) {
				if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
					return fmt.Errorf("please add new %s segment", strings.ToLower(rule.Name))
				}
				continue
			}

			if r.Segments[segIndex].Name() != rule.Name {
				if rules.IsMaskRequired(rule.Mask) {
					return fmt.Errorf("segment(%02d)'s name is not equal with rule's name (%s)", segIndex, strings.ToLower(rule.Name))
				}
				continue
			}

			if err := r.Segments[segIndex].Validate(&rule.Elements); err != nil {
				if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
					return fmt.Errorf("segment(%02d) should be valid %s segment", segIndex, strings.ToLower(rule.Name))
				}
				continue
			}

			segIndex++
		}

		index++
	}

	if len(r.Segments) > segIndex {
		if segIndex == len(r.Segments)-1 {
			return fmt.Errorf("unable to validate segment(%02d), rule is not specified", segIndex)
		} else {
			return fmt.Errorf("unable to validate segment(%02d~%02d), rule is not specified", segIndex, len(r.Segments)-1)
		}
	}

	return nil
}

func (r *Loop) Parse(data string, args ...string) (int, error) {

	if r.rule == nil {
		return 0, errors.New("please specify rules for this loop")
	}

	var newSegments []segments.SegmentInterface
	var read int

	segRules := r.rule.Segments
	index := 0

	for rule := segRules.Get(index); rule != nil; rule = segRules.Get(index) {
		for repeatIdx := 0; repeatIdx < rule.Repeat(); repeatIdx++ {
			segment, err := segments.CreateSegment(rule.Name, rule)
			if err != nil {
				log.Println(err)
				return 0, fmt.Errorf("unable to parse %s segment", strings.ToLower(rule.Name))
			}

			size, err := segment.Parse(data[read:], args...)
			if err != nil {
				if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
					return 0, fmt.Errorf("unable to parse %s segment", strings.ToLower(rule.Name))
				}
				continue
			} else {
				read += size
				newSegments = append(newSegments, segment)
			}
		}

		index++
	}

	r.Segments = newSegments

	return read, nil
}

func (r Loop) String(args ...string) string {
	var buf bytes.Buffer

	for index := range r.Segments {
		buf.WriteString(r.Segments[index].String(args...))
	}

	return buf.String()
}
