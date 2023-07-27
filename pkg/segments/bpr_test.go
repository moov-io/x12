// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForBPR(t *testing.T) {
	t.Run("parsing of bpr segment", func(t *testing.T) {
		seg := NewBPR(nil)

		in := "BPR*9999999*4*385.20*0*385.20*12*999999999999*13*2~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "BPR*9999999*4*385.20*0*385.20*12*999999999999*13*2*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "BPR*IC~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse bpr's element (02), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "SBR"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "bpr segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GTA~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "bpr segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of bpr segment", func(t *testing.T) {
		seg := NewBPR(nil)

		require.Equal(t, "BPR****~", seg.String())

		in := "BPR*9999999*4*385.20*0*385.20*12*999999999999*13*2~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of bpr segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"IC"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
			"07": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewBPR(&rule)

		in := "BPR*IC*BUSINESS OFFICE*TE*5052484349~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "BPR*IC*BUSINESS OFFICE*TE~", seg.String())

		seg.SetFieldByIndex("01", "B")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "bpr's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "BPR*B*BUSINESS OFFICE*TE*5052484349~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse bpr's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
