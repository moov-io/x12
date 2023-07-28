// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForN4(t *testing.T) {
	t.Run("parsing of n4 segment", func(t *testing.T) {
		seg := NewN4(nil)

		in := "N4*ALBUQUERQUE*NM*871201234~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "N4*ALBUQUERQUE*NM*871201234*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "N4*ALBUQUERQUE**~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "N4*ALBUQUERQUE*NM~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse n4's element (03), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "N4"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "n4 segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "N3~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "n4 segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of n4 segment", func(t *testing.T) {
		seg := NewN4(nil)

		require.Equal(t, "N4***~", seg.String())

		in := "N4*ALBUQUERQUE*NM*871201234~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of n4 segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"ALBUQUERQUE"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewN4(&rule)

		in := "N4*ALBUQUERQUE*NM*871201234~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "N4*ALBUQUERQUE*NM~", seg.String())

		seg.SetFieldByIndex("01", "ALBUQUERQU")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "n4's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "N4*ALBUQUERQU*NM*871201234~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse n4's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
