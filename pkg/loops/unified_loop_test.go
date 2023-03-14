// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package loops

import (
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/moov-io/x12/rule_5010_837p"
	"testing"

	"github.com/stretchr/testify/require"
)

var testSegRule1 = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUBMITTER NAME-1000A",
		RepeatCount: 1,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"41"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {AcceptValues: []string{"46"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "PER",
		Description: "SUBMITTER EDI CONTACT INFORMATION-1000A",
		RepeatCount: 1,
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"IC"}},
			"03": {AcceptValues: []string{"TE"}},
			"05": {AcceptValues: []string{"EM"}, Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var testSegRule2 = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUBMITTER NAME-1001A",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"42"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {AcceptValues: []string{"46"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "PER",
		Description: "SUBMITTER EDI CONTACT INFORMATION-1000A",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"IC"}},
			"03": {AcceptValues: []string{"TE"}},
			"05": {AcceptValues: []string{"EM"}, Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var testSegRule3 = rules.Segments{
	0: rules.SegmentRule{
		Name:        "NM1",
		Description: "SUBMITTER NAME-1002A",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"43"}},
			"02": {AcceptValues: []string{"1", "2"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {AcceptValues: []string{"46"}},
		},
	},
	1: rules.SegmentRule{
		Name:        "PER",
		Description: "SUBMITTER EDI CONTACT INFORMATION-1000A",
		Elements: map[string]rules.ElementRule{
			"01": {AcceptValues: []string{"IC"}},
			"03": {AcceptValues: []string{"TE"}},
			"05": {AcceptValues: []string{"EM"}, Mask: rules.MASK_OPTIONAL},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
		},
	},
}

var testComplexRule = rules.LoopRule{
	Name:     "1000A",
	Segments: testSegRule1,
	SubLoopRule: map[int]rules.LoopRule{
		0: {
			Segments: testSegRule2,
			Mask:     rules.MASK_REQUIRED,
			Name:     "1001A",
		},
		1: {
			Segments:    testSegRule3,
			Mask:        rules.MASK_REQUIRED,
			RepeatCount: 2,
			Name:        "1002A",
		},
	},
}

func FillUnifiedLoopWithRule(loop *UnifiedLoop) {

	seg11 := segments.NewNM1(&testSegRule1.Get(0).Elements)
	seg11.SetFieldByIndex("01", "41")
	seg11.SetFieldByIndex("02", "2")
	seg11.SetFieldByIndex("08", "46")
	seg12 := segments.NewPER(&testSegRule1.Get(1).Elements)
	seg12.SetFieldByIndex("01", "IC")
	seg12.SetFieldByIndex("03", "TE")
	seg12.SetFieldByIndex("05", "EM")

	l1000a := NewLoop(&testComplexRule)
	l1000a.Segments = []segments.SegmentInterface{seg11, seg12}
	loop.Loop = *l1000a

	seg21 := segments.NewNM1(&testSegRule2.Get(0).Elements)
	seg21.SetFieldByIndex("01", "42")
	seg21.SetFieldByIndex("02", "2")
	seg21.SetFieldByIndex("08", "46")
	seg22 := segments.NewPER(&testSegRule2.Get(1).Elements)
	seg22.SetFieldByIndex("01", "IC")
	seg22.SetFieldByIndex("03", "TE")
	seg22.SetFieldByIndex("05", "EM")

	l1001a := NewLoop(&testComplexRule)
	l1001a.Segments = []segments.SegmentInterface{seg21, seg22}

	subRule1 := testComplexRule.SubLoopRule[0]
	sub1 := NewUnifiedLoop(&subRule1)
	sub1.Loop = *l1001a

	seg31 := segments.NewNM1(&testSegRule3.Get(0).Elements)
	seg31.SetFieldByIndex("01", "43")
	seg31.SetFieldByIndex("02", "2")
	seg31.SetFieldByIndex("08", "46")
	seg32 := segments.NewPER(&testSegRule3.Get(1).Elements)
	seg32.SetFieldByIndex("01", "IC")
	seg32.SetFieldByIndex("03", "TE")
	seg32.SetFieldByIndex("05", "EM")

	l1002a := NewLoop(&testComplexRule)
	l1002a.Segments = []segments.SegmentInterface{seg31, seg32}

	subRule2 := testComplexRule.SubLoopRule[1]
	sub2 := NewUnifiedLoop(&subRule2)
	sub2.Loop = *l1002a

	loop.SubLoops = []UnifiedLoop{*sub1, *sub2}

	return
}

func TestUnifiedLoop(t *testing.T) {

	t.Run("testing empty unified loop", func(t *testing.T) {

		loop := &UnifiedLoop{}
		require.Error(t, loop.Validate(nil))
		require.Equal(t, "please specify rules for this unified loop", loop.Validate(nil).Error())
		require.Equal(t, "", loop.String())
		require.Equal(t, "unified loop", loop.Name())

		read, err := loop.Parse("")
		require.Error(t, err)
		require.Equal(t, "please specify rules for this unified loop", err.Error())
		require.Equal(t, 0, read)

		loop = NewUnifiedLoop(nil)
		require.Error(t, loop.Validate(nil))
		require.Equal(t, "please specify rules for this unified loop", loop.Validate(nil).Error())
		require.Equal(t, "", loop.String())
		require.Equal(t, "unified loop", loop.Name())

		read, err = loop.Parse("")
		require.Error(t, err)
		require.Equal(t, "please specify rules for this unified loop", err.Error())
		require.Equal(t, 0, read)

		loop = NewUnifiedLoop(&testRule)
		require.NotNil(t, loop.GetRule())

		loop.SetRule(nil)
		require.Nil(t, loop.GetRule())

		require.Equal(t, "unified loop", loop.Name())
	})

	t.Run("testing unified loop with specified rule", func(t *testing.T) {

		loop := NewUnifiedLoop(&testComplexRule)
		require.Error(t, loop.Validate(nil))
		require.Equal(t, "loop(1000A) is invalid, please add new nm1 segment", loop.Validate(nil).Error())
		require.Equal(t, "", loop.String())
		require.Equal(t, "1000A", loop.Name())

		FillUnifiedLoopWithRule(loop)
		require.Equal(t, 2, len(loop.SubLoops))
		require.NoError(t, loop.Validate(nil))
		require.NoError(t, loop.Validate(&testComplexRule))

		subLoops := loop.SubLoops

		loop.SubLoops = append(subLoops, subLoops[1], subLoops[1])
		require.Equal(t, 4, len(loop.SubLoops))
		err := loop.Validate(nil)
		require.Error(t, err)

		loop.SubLoops = loop.SubLoops[0:1]
		err = loop.Validate(nil)
		require.Equal(t, 1, len(loop.SubLoops))
		require.Error(t, err)
		require.Equal(t, "please add new 1002a loop", err.Error())

	})

	t.Run("testing unified loop with 1000A", func(t *testing.T) {

		rule := rules.LoopRule{
			Segments: rule_5010_837p.L1000ARule,
			Mask:     rules.MASK_REQUIRED,
			Name:     "1000A",
		}

		loop := NewUnifiedLoop(&rule)

		read, err := loop.Parse("NM1*41*2*CS*****46*133052274~PER*IC*CUSTOMER SOLUTIONS*TE*8008456592~")
		require.NoError(t, err)
		require.Equal(t, 69, read)
		require.Equal(t, nil, loop.Validate(nil))
		require.Equal(t, nil, loop.Validate(&rule))
	})

}
