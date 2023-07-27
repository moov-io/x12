// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForSBR(t *testing.T) {
	t.Run("parsing of sbr segment", func(t *testing.T) {
		seg := NewSBR(nil)

		in := "SBR*P*18*******MC~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "SBR*P*18*******MC*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "SBR*P*18~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "SBR~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse sbr's element (01), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "SBR"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "sbr segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GTA~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "sbr segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of sbr segment", func(t *testing.T) {
		seg := NewSBR(nil)

		require.Equal(t, "SBR*~", seg.String())

		in := "SBR*P*18*******MC~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of sbr segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"P"}, Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewSBR(&rule)

		in := "SBR*P*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "SBR*P~", seg.String())

		seg.SetFieldByIndex("01", "B")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "sbr's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "SBR*B~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse sbr's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
