// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package loops

import (
	"bytes"
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/moov-io/x12/pkg/util"
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

func (r *Loop) Validate(setRules *rules.SegmentSetRule) error {
	if setRules == nil && r.rule != nil {
		setRules = &r.rule.Segments
	}

	if setRules == nil {
		return util.NewSpecifiedRuleError(r.Name())
	}

	ruleIndex := 0
	passIndex := 0
	for rule := setRules.Get(ruleIndex); rule != nil; rule = setRules.Get(ruleIndex) {
		passSegIndex := 0
		for repeatIdx := 0; repeatIdx < rule.Repeat(); repeatIdx++ {
			if passIndex+1 > len(r.Segments) {
				break
			}

			if util.GetStructName(r.Segments[passIndex]) != rule.Name {
				if rules.IsMaskRequired(rule.Mask) {
					return util.NewMismatchSegmentError(util.GetStructName(r.Segments[passIndex]), rule.Name)
				}
				break
			}

			if err := r.Segments[passIndex].Validate(&rule.Elements); err != nil {
				if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
					return err
				}
				break
			}

			passSegIndex++
			passIndex++
		}

		if passSegIndex < rule.MinRepeat() {
			return util.NewMinSegmentError(rule.Name)
		}

		ruleIndex++
	}

	if len(r.Segments) > passIndex {
		return util.NewAllSegmentError(r.Name())
	}

	return nil
}

func (r *Loop) Parse(data string, args ...string) (int, error) {
	if r.rule == nil {
		return 0, util.NewSpecifiedRuleError(r.Name())
	}

	var newSegments []segments.SegmentInterface
	var read int

	segRules := r.rule.Segments
	index := 0

	for rule := segRules.Get(index); rule != nil; rule = segRules.Get(index) {
		for repeatIdx := 0; repeatIdx < rule.Repeat(); repeatIdx++ {
			segment, err := segments.CreateSegment(rule.Name, rule)
			if err != nil {
				if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
					util.AppendErrorSegmentLine(err, data[read:], args...)
					return 0, err
				}
				break
			}

			size, err := segment.Parse(data[read:], args...)
			if err != nil {
				if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
					util.AppendErrorSegmentLine(err, data[read:], args...)
					return 0, err
				}
				break
			} else {
				read += size
				newSegments = append(newSegments, segment)
			}
		}

		index++
	}

	if len(newSegments) == 0 {
		return 0, util.NewUnableParseError(r.Name())
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

func (r Loop) DumpStructInfo(level int) []util.ElementInfo {
	var selfDumps []util.ElementInfo

	selfDumps = append(selfDumps, util.ElementInfo{Name: r.Name(), Level: level})

	for _, s := range r.Segments {
		dumps := util.DumpStructInfo(s, level+1)
		selfDumps = append(selfDumps, dumps)
	}

	return selfDumps
}
