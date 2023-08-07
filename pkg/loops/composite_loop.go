// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package loops

import (
	"bytes"
	"strings"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/moov-io/x12/pkg/util"
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

func (r CompositeLoop) GetSegments() []segments.SegmentInterface {
	segments := r.Loop.Segments
	for _, loop := range r.SubLoops {
		segments = append(segments, loop.GetSegments()...)
	}

	return segments
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
		return util.NewSpecifiedRuleError(r.Name())
	}

	err := r.Loop.Validate(&loopRule.Segments)
	if err != nil {
		return err
	}

	passIndex := 0
	for index := 0; index < len(loopRule.Composite); index++ {
		rule := loopRule.Composite[index]
		passLoopIndex := 0
		for repeatIdx := 0; repeatIdx < rule.Repeat(); repeatIdx++ {
			if passIndex+1 > len(r.SubLoops) {
				break
			}

			if r.SubLoops[passIndex].Name() != rule.Name {
				if rules.IsMaskRequired(rule.Mask) {
					return util.NewMismatchLoopError(r.SubLoops[passIndex].Name(), rule.Name)
				}
				break
			}

			if err = r.SubLoops[passIndex].Validate(&rule); err != nil {
				if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
					util.AppendErrorStack(err, r.SubLoops[passIndex].Name())
					return err
				}
				break
			}

			passLoopIndex++
			passIndex++
		}

		if passLoopIndex < rule.MinRepeat() {
			return util.NewMinLoopError(rule.Name)
		}
	}

	if len(r.SubLoops) > passIndex {
		return util.NewAllLoopError(r.Name())
	}

	return nil
}

func (r *CompositeLoop) Parse(data string, args ...string) (int, error) {
	if r.rule == nil {
		return 0, util.NewSpecifiedRuleError(r.Name())
	}

	r.Loop.SetRule(r.rule)
	size, err := r.Loop.Parse(data, args...)
	if err != nil {
		util.AppendErrorStack(err, strings.ToLower(r.rule.Name))
		return 0, err
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
			if err == nil && size > 0 {
				read += size
				line = data[read:]
				r.SubLoops = append(r.SubLoops, *newChild)
			} else {
				if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
					return 0, err
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

func (r CompositeLoop) DumpStructInfo(level int) []util.ElementInfo {
	var selfDumps []util.ElementInfo

	dumps := r.Loop.DumpStructInfo(level)
	selfDumps = append(selfDumps, dumps...)

	for _, s := range r.SubLoops {
		dumps = s.DumpStructInfo(level + 1)
		selfDumps = append(selfDumps, dumps...)
	}

	return selfDumps
}
