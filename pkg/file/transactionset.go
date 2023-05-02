// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/moov-io/x12/pkg/loops"
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/moov-io/x12/pkg/util"
)

func NewTransactionSet(rule *rules.TransactionRule) *TransactionSet {

	newTransaction := TransactionSet{rule: rule}

	return &newTransaction
}

type TransactionSet struct {
	ST    segments.ST           `json:"ST" xml:"ST"`
	Loops []loops.CompositeLoop `json:"loops" xml:"loops"`
	SE    *segments.SE          `json:"SE,omitempty" xml:"SE,omitempty"`

	Segments []segments.SegmentInterface

	rule *rules.TransactionRule
}

func (r *TransactionSet) Validate(transRule *rules.TransactionRule) error {

	if transRule == nil && r.rule != nil {
		transRule = r.rule
	}

	var err error
	var segIndex, index int

	// Validating ST Segment
	stRule := transRule.ST
	{
		if stRule.Name != "ST" {
			return errors.New("invalid st rule")
		}

		err = r.ST.Validate(&stRule.Elements)
		if err != nil {
			return errors.New("unable to validate st segment")
		}
	}

	// Validating segments (BHT, BPR, TRN, CUR, REF, ...)
	{
		index = 0
		segIndex = 0
		segRules := transRule.Segments
		for rule := segRules.Get(index); rule != nil; rule = segRules.Get(index) {

			for repeatIdx := 0; repeatIdx < rule.Repeat(); repeatIdx++ {

				if segIndex+1 > len(r.Segments) {
					if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
						return fmt.Errorf("please add new %s segment", strings.ToUpper(rule.Name))
					}
					break
				}

				if r.Segments[segIndex].Name() != rule.Name {
					if rules.IsMaskRequired(rule.Mask) {
						return fmt.Errorf("segment(%02d)'s name is not equal with rule's name (%s)", segIndex, strings.ToLower(rule.Name))
					}
					break
				}

				if validateErr := r.Segments[segIndex].Validate(&rule.Elements); validateErr != nil {
					if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
						return fmt.Errorf("segment(%02d) should be valid %s segment", segIndex, strings.ToUpper(rule.Name))
					}
					break
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
	}

	// Validating loops
	{
		segIndex = 0
		for index = 0; index < len(transRule.Loops); index++ {
			rule := transRule.Loops[index]

			for repeatIdx := 0; repeatIdx < rule.Repeat(); repeatIdx++ {

				if segIndex+1 > len(r.Loops) {
					if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
						return fmt.Errorf("please add new %s loop", strings.ToUpper(rule.Name))
					}
					break
				}

				if r.Loops[segIndex].Name() != rule.Name {
					if rules.IsMaskRequired(rule.Mask) {
						return fmt.Errorf("loop(%02d)'s name is not equal with rule's name (%s)", segIndex, strings.ToLower(rule.Name))
					}
					break
				}

				if err = r.Loops[segIndex].Validate(&rule); err != nil {
					if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
						return fmt.Errorf("loop(%02d) should have valid %s loop, %s", segIndex, strings.ToLower(rule.Name), err.Error())
					}
					break
				}

				segIndex++
			}
		}

		if len(r.Loops) > segIndex {
			if segIndex == len(r.Loops)-1 {
				return fmt.Errorf("unable to validate loop(%02d), rule is not specified", segIndex)
			} else {
				return fmt.Errorf("unable to validate loop(%02d~%02d), rule is not specified", segIndex, len(r.Loops)-1)
			}
		}

	}

	// Validating SE Segment
	seRule := transRule.SE
	if rules.IsMaskRequired(seRule.Mask) && r.SE == nil {
		return errors.New("se segment is required segment")
	}

	if r.SE != nil {
		if seRule.Name != "SE" {
			return errors.New("invalid se rule")
		}

		err = r.SE.Validate(&seRule.Elements)
		if err != nil && rules.IsMaskRequired(seRule.Mask) {
			return errors.New("unable to validate se segment")
		}
	}

	// Validating Transaction Set
	if r.SE != nil {

		// compare control set number
		if r.SE.TransactionSetControlNumber != r.ST.TransactionSetControlNumber {
			return errors.New("has invalid transaction set control number")
		}

		// getting segments count
		segmentCnt := 2

		// segments
		segmentCnt += len(r.Segments)

		for _, loop := range r.Loops {
			segs := loop.GetSegments()
			segmentCnt += len(segs)
		}

		// compare number of segments
		if v, conErr := strconv.ParseInt(r.SE.NumberOfSegments, 10, 32); conErr == nil {
			if v != int64(segmentCnt) {
				return errors.New("has invalid number of segments")
			}
		}
	}

	return nil
}

func (r *TransactionSet) Parse(data string, args ...string) (int, error) {

	if r.rule == nil {
		return 0, errors.New("please specify rules for this transaction set")
	}

	if len(r.rule.Loops) == 0 {
		return 0, errors.New("missing loops rules")
	}

	var size, read, index int
	var err error

	// Parsing ST Segment
	stRule := r.rule.ST
	{
		if stRule.Name != "ST" {
			return 0, errors.New("invalid st rule")
		}

		r.ST.SetRule(&stRule.Elements)
		size, err = r.ST.Parse(data[read:], args...)
		if err != nil {
			return 0, errors.New("unable to parse st segment, (" + err.Error() + ")")
		} else {
			read += size
		}
	}

	// Parsing Segments
	{
		segRules := r.rule.Segments
		index = 0
		for rule := segRules.Get(index); rule != nil; rule = segRules.Get(index) {
			for repeatIdx := 0; repeatIdx < rule.Repeat(); repeatIdx++ {
				newSegment, createErr := segments.CreateSegment(rule.Name, rule)
				if createErr != nil {
					if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
						return 0, fmt.Errorf("unable to parse %s segment", strings.ToLower(rule.Name))
					}
					break
				}

				readSize, parseErr := newSegment.Parse(data[read:], args...)
				if parseErr != nil {
					if repeatIdx == 0 && rules.IsMaskRequired(rule.Mask) {
						return 0, fmt.Errorf("unable to parse %s segment (%s)", strings.ToLower(rule.Name), parseErr.Error())
					}
					break
				} else {
					read += readSize
					r.Segments = append(r.Segments, newSegment)
				}
			}

			index++
		}
	}

	// Parsing LOOPS
	{
		for index = 0; index < len(r.rule.Loops); index++ {
			loopRule := r.rule.Loops[index]

			for repeatIdx := 0; repeatIdx < loopRule.Repeat(); repeatIdx++ {
				newLoop := loops.NewCompositeLoop(&loopRule)
				size, err = newLoop.Parse(data[read:], args...)
				if err == nil {
					read += size
					r.Loops = append(r.Loops, *newLoop)
				} else {
					if repeatIdx == 0 && rules.IsMaskRequired(loopRule.Mask) {
						return 0, fmt.Errorf("unable to parse %s loop", strings.ToLower(loopRule.Name))
					}
					break
				}
			}

		}
	}

	// Parsing SE Segment
	seRule := r.rule.SE
	if seRule.Name == "SE" {
		newSE := segments.NewSE(&seRule.Elements)
		size, err = newSE.Parse(data[read:], args...)
		if err != nil && rules.IsMaskRequired(seRule.Mask) {
			return 0, errors.New("unable to parse se segment")
		} else if err == nil {
			read += size
			if s, ok := newSE.(*segments.SE); ok {
				r.SE = s
			}
		}
	}

	return read, nil
}

func (r TransactionSet) String(args ...string) string {
	var buf bytes.Buffer

	buf.WriteString(r.ST.String(args...))

	for index := range r.Segments {
		buf.WriteString(r.Segments[index].String(args...))
	}
	for index := range r.Loops {
		buf.WriteString(r.Loops[index].String(args...))
	}
	if r.SE != nil {
		buf.WriteString(r.SE.String(args...))
	}

	return buf.String()
}

func (r TransactionSet) DumpStructInfo(level int) []util.ElementInfo {
	var selfDumps []util.ElementInfo

	dump := util.DumpStructInfo(r.ST, level)
	selfDumps = append(selfDumps, dump)

	for _, l := range r.Loops {
		dumps := l.DumpStructInfo(level + 1)
		selfDumps = append(selfDumps, dumps...)
	}

	if r.SE != nil {
		dump = util.DumpStructInfo(r.SE, level)
		selfDumps = append(selfDumps, dump)
	}

	return selfDumps
}
