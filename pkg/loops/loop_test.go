// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package loops

import (
	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/segments"
	"github.com/stretchr/testify/require"
	"testing"
)

var testSegRule = rules.SegmentSetRule{
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

var testRule = rules.LoopRule{
	Name:     "1000A",
	Segments: testSegRule,
}

func TestNewLoop(t *testing.T) {
	t.Run("testing empty  loop", func(t *testing.T) {

		loop := &Loop{}
		require.Error(t, loop.Validate(nil))
		require.Equal(t, "please specify rules for this loop", loop.Validate(nil).Error())
		require.Equal(t, "", loop.String())
		require.Equal(t, "loop", loop.Name())

		read, err := loop.Parse("")
		require.Error(t, err)
		require.Equal(t, "please specify rules for this loop", err.Error())
		require.Equal(t, 0, read)

		loop = NewLoop(nil)
		require.Error(t, loop.Validate(nil))
		require.Equal(t, "please specify rules for this loop", loop.Validate(nil).Error())
		require.Equal(t, "", loop.String())
		require.Equal(t, "loop", loop.Name())

		read, err = loop.Parse("")
		require.Error(t, err)
		require.Equal(t, "please specify rules for this loop", err.Error())
		require.Equal(t, 0, read)

		loop = NewLoop(&testRule)
		require.NotNil(t, loop.GetRule())

		loop.SetRule(nil)
		require.Nil(t, loop.GetRule())

		require.Equal(t, "loop", loop.Name())
	})
}

func TestLoop100A(t *testing.T) {
	t.Run("testing loop 1000a", func(t *testing.T) {

		loop := NewLoop(&testRule)

		require.Error(t, loop.Validate(nil))
		require.Equal(t, "", loop.String())

		loop.Segments = append(loop.Segments, segments.NewNM1(nil), segments.NewPER(nil))
		require.Error(t, loop.Validate(nil))
		require.Equal(t, "segment(00) should be valid NM1 segment", loop.Validate(nil).Error())
		require.Equal(t, "NM1*********~PER****~", loop.String())

	})

	t.Run("parsing loop 1000a", func(t *testing.T) {

		in := "NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~PER*IC*JERRY*TE*7176149999~"
		loop := NewLoop(&testRule)

		read, err := loop.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, loop.String())
		require.NoError(t, loop.Validate(nil))

		in = "NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~PER*IC*JERRY*TE*7176149999~NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~"
		read, err = loop.Parse(in)
		require.NoError(t, err)
		require.Equal(t, 73, read)
		require.Equal(t, "NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~PER*IC*JERRY*TE*7176149999~", loop.String())
		require.NoError(t, loop.Validate(nil))

		in = "PER*IC*JERRY*TE*7176149999~NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~"
		read, err = loop.Parse(in)
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse nm1 segment (nm1 segment contains invalid code)", err.Error())

		in = "NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~PRV*IC*JERRY*TE*7176149999~"
		read, err = loop.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse per segment (per segment contains invalid code)", err.Error())
		require.Equal(t, 0, read)

		in = "NM1*43*2*PREMIER BILLING SERVICE*****46*TGJ23~PRV*IC*JERRY*TE*7176149999~"
		read, err = loop.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse nm1 segment (unable to parse nm1's element (01), the element contains unexpected value)", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("validating loop 1000a", func(t *testing.T) {

		loop := NewLoop(&testRule)

		nm1 := segments.NM1{}
		nm1.SetFieldByIndex("01", "41")
		nm1.SetFieldByIndex("02", "2")
		nm1.SetFieldByIndex("03", "PREMIER BILLING SERVICE")
		nm1.SetFieldByIndex("08", "46")
		nm1.SetFieldByIndex("09", "TGJ23")

		per1 := segments.PER{}
		per1.SetFieldByIndex("01", "IC")
		per1.SetFieldByIndex("02", "JERRY")
		per1.SetFieldByIndex("03", "TE")
		per1.SetFieldByIndex("04", "7176149999")

		loop.Segments = []segments.SegmentInterface{&nm1, &per1}

		err := loop.Validate(nil)
		require.NoError(t, err)

		loop.Segments = []segments.SegmentInterface{&nm1, &nm1}

		err = loop.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "segment(01)'s name is not equal with rule's name (per)", err.Error())

		loop.Segments = []segments.SegmentInterface{&per1, &per1}

		err = loop.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "segment(00)'s name is not equal with rule's name (nm1)", err.Error())

		loop.Segments = []segments.SegmentInterface{&nm1, &per1, &per1}

		err = loop.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "unable to validate segment(02), rule is not specified", err.Error())

		loop.Segments = []segments.SegmentInterface{&nm1, &per1, &nm1, &per1}

		err = loop.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "unable to validate segment(02~03), rule is not specified", err.Error())

		per1.SetFieldByIndex("01", "XX")

		loop.Segments = []segments.SegmentInterface{&nm1, &per1}

		err = loop.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "segment(01) should be valid PER segment", err.Error())

		loop.Segments = []segments.SegmentInterface{&nm1}

		err = loop.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "please add new PER segment", err.Error())
	})

	t.Run("testing loop 1000a with repeated segments", func(t *testing.T) {

		var rule = rules.LoopRule{
			Name: "1000A",
			Segments: rules.SegmentSetRule{
				0: rules.SegmentRule{
					Name:        "NM1",
					Description: "SUBMITTER NAME-1000A",
					RepeatCount: 2,
					Mask:        rules.MASK_OPTIONAL,
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
			},
		}

		loop := NewLoop(&rule)

		nm1 := segments.NM1{}
		nm1.SetFieldByIndex("01", "41")
		nm1.SetFieldByIndex("02", "2")
		nm1.SetFieldByIndex("03", "PREMIER BILLING SERVICE")
		nm1.SetFieldByIndex("08", "46")
		nm1.SetFieldByIndex("09", "TGJ23")

		per1 := segments.PER{}
		per1.SetFieldByIndex("01", "IC")
		per1.SetFieldByIndex("02", "JERRY")
		per1.SetFieldByIndex("03", "TE")
		per1.SetFieldByIndex("04", "7176149999")

		loop.Segments = []segments.SegmentInterface{&nm1, &nm1, &per1}

		err := loop.Validate(nil)
		require.NoError(t, err)

		loop.Segments = []segments.SegmentInterface{&nm1, &per1}

		err = loop.Validate(nil)
		require.NoError(t, err)

		loop.Segments = []segments.SegmentInterface{&per1}

		err = loop.Validate(nil)
		require.NoError(t, err)

		in := "NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~PER*IC*JERRY*TE*7176149999~"

		read, err := loop.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, loop.String())
		require.NoError(t, loop.Validate(nil))

		in = "NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~PER*IC*JERRY*TE*7176149999~"

		read, err = loop.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, loop.String())
		require.NoError(t, loop.Validate(nil))

		in = "PER*IC*JERRY*TE*7176149999~"

		read, err = loop.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, loop.String())
		require.NoError(t, loop.Validate(nil))

	})

	t.Run("testing loop 1000a with repeated segments", func(t *testing.T) {

		var rule = rules.LoopRule{
			Name: "1000A",
			Segments: rules.SegmentSetRule{
				0: rules.SegmentRule{
					Name:        "NM1",
					Description: "SUBMITTER NAME-1000A",
					RepeatCount: 2,
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
			},
		}

		loop := NewLoop(&rule)

		nm1 := segments.NM1{}
		nm1.SetFieldByIndex("01", "41")
		nm1.SetFieldByIndex("02", "2")
		nm1.SetFieldByIndex("03", "PREMIER BILLING SERVICE")
		nm1.SetFieldByIndex("08", "46")
		nm1.SetFieldByIndex("09", "TGJ23")

		per1 := segments.PER{}
		per1.SetFieldByIndex("01", "IC")
		per1.SetFieldByIndex("02", "JERRY")
		per1.SetFieldByIndex("03", "TE")
		per1.SetFieldByIndex("04", "7176149999")

		loop.Segments = []segments.SegmentInterface{&nm1, &nm1, &per1}

		err := loop.Validate(nil)
		require.NoError(t, err)

		loop.Segments = []segments.SegmentInterface{&nm1, &nm1}

		err = loop.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "please add new PER segment", err.Error())

		loop.Segments = []segments.SegmentInterface{&nm1}

		err = loop.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "please add new PER segment", err.Error())

		in := "NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~NM1*41*2*PREMIER BILLING SERVICE*****46*TGJ23~PER*IC*JERRY*TE*7176149999~"

		read, err := loop.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, loop.String())
		require.NoError(t, loop.Validate(nil))

	})
}
