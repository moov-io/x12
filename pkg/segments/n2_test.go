// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForN2(t *testing.T) {
	t.Run("parsing of n2 segment", func(t *testing.T) {
		seg := NewN2(nil)

		in := "N2*3*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "N2*3*8*~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "n2 segment can't parse all input data", err.Error())
		require.Equal(t, 0, read)

		in = "N2*3*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "N2~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse n2's element (01), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "N2"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "n2 segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "n2 segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of n2 segment", func(t *testing.T) {
		seg := NewN2(nil)

		require.Equal(t, "N2*~", seg.String())

		in := "N2*3~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of n2 segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"5"}, Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewN2(&rule)

		in := "N2*5*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "N2*5~", seg.String())

		seg.SetFieldByIndex("01", "6")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "n2's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "N2*6*8~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse n2's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
