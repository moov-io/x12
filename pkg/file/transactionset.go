// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/moov-io/x12/pkg/loops"
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
)

func NewTransactionSet(rule *rules.TransactionRule) *TransactionSet {

	newTransaction := TransactionSet{rule: rule}

	return &newTransaction
}

type TransactionSet struct {
	ST    segments.ST         `json:"ST" xml:"ST"`
	BHT   *segments.BHT       `json:"BHT,omitempty" xml:"BHT,omitempty"`
	Loops []loops.UnifiedLoop `json:"loops" xml:"loops"`
	SE    *segments.SE        `json:"SE,omitempty" xml:"SE,omitempty"`

	rule *rules.TransactionRule
}

func (r *TransactionSet) Validate(transRule *rules.TransactionRule) error {

	if transRule == nil && r.rule != nil {
		transRule = r.rule
	}

	var err error

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

	// Validating BHT Segment
	bhtRule := transRule.BHT
	if rules.IsMaskRequired(bhtRule.Mask) && r.BHT == nil {
		return errors.New("bht segment is required segment")
	}

	if r.BHT != nil {
		if bhtRule.Name != "BHT" {
			return errors.New("invalid st rule")
		}

		err = r.BHT.Validate(&bhtRule.Elements)
		if err != nil && rules.IsMaskRequired(bhtRule.Mask) {
			return errors.New("unable to validate bht segment")
		}
	}

	// Validating loops
	{
		segIndex := 0
		for index := 0; index < len(transRule.Loops); index++ {
			rule := transRule.Loops[index]

			for repeatCnt := 0; repeatCnt < rule.Repeat(); repeatCnt++ {

				if segIndex+1 > len(r.Loops) {
					if repeatCnt == 0 && rules.IsMaskRequired(rule.Mask) {
						return fmt.Errorf("please add new %s loop", strings.ToLower(rule.Name))
					}
					continue
				}

				if r.Loops[segIndex].Name() != rule.Name {
					if rules.IsMaskRequired(rule.Mask) {
						return fmt.Errorf("loop(%02d)'s name is not equal with rule's name (%s)", segIndex, strings.ToLower(rule.Name))
					}
					continue
				}

				if err = r.Loops[segIndex].Validate(&rule); err != nil {
					if rules.IsMaskRequired(rule.Mask) {
						log.Println(err)
						return fmt.Errorf("loop(%02d) should have valid %s loop, %s", segIndex, strings.ToLower(rule.Name), err.Error())
					}
					continue
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

	return nil
}

func (r *TransactionSet) Parse(data string, args ...string) (int, error) {

	if r.rule == nil {
		return 0, errors.New("please specify rules for this transaction set")
	}

	if len(r.rule.Loops) == 0 {
		return 0, errors.New("missing loops rules")
	}

	var size, read int
	var err error

	line := data[read:]

	// Parsing ST Segment
	stRule := r.rule.ST
	{
		if stRule.Name != "ST" {
			return 0, errors.New("invalid st rule")
		}

		r.ST.SetRule(&stRule.Elements)
		size, err = r.ST.Parse(line, args...)
		if err != nil {
			return 0, errors.New("unable to parse st segment")
		} else {
			read += size
			line = data[read:]
		}
	}

	// Parsing BHT Segment
	bhtRule := r.rule.BHT
	if bhtRule.Name == "BHT" {
		newBHT := segments.NewBHT(&bhtRule.Elements)
		size, err = newBHT.Parse(line, args...)
		if err != nil && rules.IsMaskRequired(bhtRule.Mask) {
			return 0, errors.New("unable to parse bht segment")
		} else if err == nil {
			read += size
			line = data[read:]
			r.BHT = newBHT.(*segments.BHT)
		}
	}

	// Parsing LOOPS
	for index := 0; index < len(r.rule.Loops); index++ {
		loopRule := r.rule.Loops[index]

		for repeatIdx := 0; repeatIdx < loopRule.Repeat(); repeatIdx++ {
			newLoop := loops.NewUnifiedLoop(&loopRule)
			size, err = newLoop.Parse(line, args...)
			if err == nil {
				read += size
				line = data[read:]
				r.Loops = append(r.Loops, *newLoop)
			} else {
				if repeatIdx == 0 && rules.IsMaskRequired(loopRule.Mask) {
					return 0, fmt.Errorf("unable to parse %s loop", strings.ToLower(loopRule.Name))
				}
				break
			}
		}

	}

	// Parsing SE Segment
	seRule := r.rule.SE
	if seRule.Name == "SE" {
		newSE := segments.NewSE(&seRule.Elements)
		size, err = newSE.Parse(line, args...)
		if err != nil && rules.IsMaskRequired(seRule.Mask) {
			return 0, errors.New("unable to parse se segment")
		} else if err == nil {
			read += size
			line = data[read:]
			r.SE = newSE.(*segments.SE)
		}
	}

	return read, nil
}

func (r TransactionSet) String(args ...string) string {
	var buf bytes.Buffer

	buf.WriteString(r.ST.String(args...))
	if r.BHT != nil {
		buf.WriteString(r.BHT.String(args...))
	}
	for index := range r.Loops {
		buf.WriteString(r.Loops[index].String(args...))
	}
	if r.SE != nil {
		buf.WriteString(r.SE.String(args...))
	}

	return buf.String()
}
