// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForPER(t *testing.T) {
	t.Run("parsing of per segment", func(t *testing.T) {
		seg := NewPER(nil)

		in := "PER*IC*BUSINESS OFFICE*TE*5052484349~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PER*IC*BUSINESS OFFICE*TE*5052484349*****~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "PER*IC*BUSINESS OFFICE*TE*5052484349~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PER*IC*BUSINESS OFFICE*TE~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse per's element (04), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "SBR"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "per segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GTA~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "per segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of per segment", func(t *testing.T) {
		seg := NewPER(nil)

		require.Equal(t, "PER****~", seg.String())

		in := "PER*IC*BUSINESS OFFICE*TE*5052484349~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of per segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"IC"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewPER(&rule)

		in := "PER*IC*BUSINESS OFFICE*TE*5052484349~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "PER*IC*BUSINESS OFFICE*TE~", seg.String())

		seg.SetFieldByIndex("01", "B")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "per's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "PER*B*BUSINESS OFFICE*TE*5052484349~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse per's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
