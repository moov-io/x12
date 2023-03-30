// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"github.com/moov-io/x12/pkg/rules"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForCLM(t *testing.T) {

	t.Run("parsing of clm segment", func(t *testing.T) {

		seg := NewCLM(nil)

		in := "CLM*1-1180*174***11:B:1*Y*A*Y*Y*P~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "CLM*1-1180*174***11:B:1*Y*A*Y*Y*P***********~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "CLM*1-1180*174***11:B:1*Y*A*Y*Y~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "CLM*1-1180*174***11:B:1*Y*A*Y~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse clm's element (09), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "clm segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "clm segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of clm segment", func(t *testing.T) {

		seg := NewCLM(nil)

		require.Equal(t, "CLM*********~", seg.String())

		in := "CLM*1-1180*174***11<B<1*Y*A*Y*Y*P~"
		read, err := seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse clm's element (05), unable to parse service location's element (02), doesn't enough input string", err.Error())

		read, err = seg.Parse(in, "<")
		require.NoError(t, err)
		require.Equal(t, len(in), read)
	})

	t.Run("parsing and encoding of clm segment with specified rule", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {
				Mask: rules.MASK_REQUIRED,
				Composite: rules.ElementSetRule{
					"02": {AcceptValues: []string{"2"}},
					"03": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
				},
			},
			"06": {AcceptValues: []string{"Y", "N"}},
			"07": {AcceptValues: []string{"A", "B", "C"}},
			"08": {AcceptValues: []string{"W"}},
			"09": {AcceptValues: []string{"I", "Y"}},
			"10": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"P"}},
			"11": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"AA", "EM", "OA"}},
					"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AA", "EM", "OA"}},
					"03": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"AA", "EM", "OA"}},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
				},
			},
			"12": {Mask: rules.MASK_OPTIONAL},
			"13": {Mask: rules.MASK_NOTUSED},
			"14": {Mask: rules.MASK_NOTUSED},
			"15": {Mask: rules.MASK_NOTUSED},
			"16": {Mask: rules.MASK_NOTUSED},
			"17": {Mask: rules.MASK_NOTUSED},
			"18": {Mask: rules.MASK_NOTUSED},
			"19": {Mask: rules.MASK_NOTUSED},
			"20": {Mask: rules.MASK_OPTIONAL},
		}

		seg := NewCLM(&rule)

		in := "CLM*1-1180*174***11:2:1*Y*A*W*Y*P~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "CLM*1-1180*174***11:2:1*Y*A*W*Y*P~", seg.String())

		seg.SetFieldByIndex("06", "W")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "clm's element (06) has invalid value, the element contains unexpected value", err.Error())

		in = "CLM*1-1180*174***11:2:1*Y*A*Y*Y*P~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse clm's element (08), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)

		in = "CLM*1-1180*174***11:A:1*Y*A*W*Y*P~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse clm's element (05), unable to parse service location's element (02), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)

		in = "CLM*1-1180*174***11:2:1*Y*A*W*Y*P*AA:PP~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "CLM*1-1180*174***11:2:1*Y*A*W*Y*P~", seg.String())
	})

	t.Run("parsing and encoding of clm segment with specified rule (837d)", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {
				Mask: rules.MASK_REQUIRED,
				Composite: rules.ElementSetRule{
					"02": {AcceptValues: []string{"B"}},
				},
			},
			"06": {AcceptValues: []string{"Y", "N"}},
			"07": {AcceptValues: []string{"A", "C"}},
			"08": {AcceptValues: []string{"N", "W", "Y"}},
			"09": {AcceptValues: []string{"I", "Y"}},
			"10": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"P"}},
			"11": {
				Mask: rules.MASK_OPTIONAL,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"AA", "EM", "OA"}},
					"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AA", "EM", "OA"}},
					"03": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"AA", "EM", "OA"}},
					"04": {Mask: rules.MASK_OPTIONAL},
					"05": {Mask: rules.MASK_OPTIONAL},
				},
			},
			"12": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"01", "02", "03", "05"}},
			"13": {Mask: rules.MASK_NOTUSED},
			"14": {Mask: rules.MASK_NOTUSED},
			"15": {Mask: rules.MASK_NOTUSED},
			"16": {Mask: rules.MASK_NOTUSED},
			"17": {Mask: rules.MASK_NOTUSED},
			"18": {Mask: rules.MASK_NOTUSED},
			"19": {Mask: rules.MASK_OPTIONAL},
			"20": {Mask: rules.MASK_OPTIONAL},
		}

		seg := NewCLM(&rule)

		in := "CLM*26403774*150***11:B:1*Y*A*Y*I~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "CLM*26403774*150***11:B:1*Y*A*Y*I~", seg.String())
	})

}
