// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"github.com/moov-io/x12/pkg/rules"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForHI(t *testing.T) {

	t.Run("parsing of hi segment", func(t *testing.T) {

		seg := NewHI(nil)

		in := "HI*composite1:AA~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "HI*composite1:AA:::::::~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "HI*composite1~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse hi's element (01), unable to parse health care code's element (02), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "hi segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "DM~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "hi segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of hi segment", func(t *testing.T) {

		seg := NewHI(nil)

		require.Equal(t, "HI*~", seg.String())

		in := "HI*composit<AA~"
		read, err := seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse hi's element (01), unable to parse health care code's element (02), doesn't enough input string", err.Error())

		read, err = seg.Parse(in, "<")
		require.NoError(t, err)
		require.Equal(t, len(in), read)
	})

	t.Run("parsing and encoding of hi segment with specified rule", func(t *testing.T) {

		rule := rules.Elements{
			"01": {
				Mask:           rules.MASK_REQUIRED,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"composite1"}},
					"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
			"02": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"composite2"}},
					"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
			"03": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"composite3"}},
					"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
			"04": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"composite4"}},
					"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
			"05": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"composite5"}},
					"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
			"06": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"composite6"}},
					"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
			"07": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"composite7"}},
					"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
			"08": {
				Mask:           rules.MASK_OPTIONAL,
				HasSubElements: true,
				SubRule: map[string]rules.ElementRule{
					"01": {AcceptValues: []string{"composite8"}},
					"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
					"03": {Mask: rules.MASK_NOTUSED},
					"04": {Mask: rules.MASK_NOTUSED},
					"05": {Mask: rules.MASK_NOTUSED},
					"06": {Mask: rules.MASK_NOTUSED},
					"07": {Mask: rules.MASK_NOTUSED},
					"08": {Mask: rules.MASK_NOTUSED},
					"09": {Mask: rules.MASK_NOTUSED},
				},
			},
		}

		seg := NewHI(&rule)

		in := "HI*composite1:1*composite2:1*composite3:1*composite4:1*composite5:1*composite6:1*composite7:7*composite8:8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.NoError(t, seg.Validate(nil))
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		err = seg.SetFieldByIndex("08", "")
		require.Error(t, err)
		require.Equal(t, "doesn't match setting type", err.Error())

		err = seg.SetFieldByIndex("08", nil)
		in = "HI*composite1:1*composite2:1*composite3:1*composite4:1*composite5:1*composite6:1*composite7:7*composite8~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, "HI*composite1:1*composite2:1*composite3:1*composite4:1*composite5:1*composite6:1*composite7:7~", seg.String())

		err = seg.SetFieldByIndex("07", &HealthCareCode{})
		in = "HI*composite1:1*composite2:1*composite3:1*composite4:1*composite5:1*composite6:1~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, "HI*composite1:1*composite2:1*composite3:1*composite4:1*composite5:1*composite6:1~", seg.String())

		in = "CLM*1-1180*174***11:A:1*Y*A*W*Y*P~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "hi segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

}
