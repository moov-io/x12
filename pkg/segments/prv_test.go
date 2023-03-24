// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForPRV(t *testing.T) {

	t.Run("parsing of prv segment", func(t *testing.T) {

		seg := NewPRV(nil)

		in := "PRV*ALBUQUERQUE*NM*871201234~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PRV*ALBUQUERQUE*NM*871201234*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "PRV*ALBUQUERQUE**~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PRV*ALBUQUERQUE*NM~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse prv's element (03), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "PRV"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "prv segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "N3~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "prv segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of prv segment", func(t *testing.T) {

		seg := NewPRV(nil)

		require.Equal(t, "PRV***~", seg.String())

		in := "PRV*ALBUQUERQUE*NM*871201234~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of prv segment with specified rule", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"ALBUQUERQUE"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewPRV(&rule)

		in := "PRV*ALBUQUERQUE*NM*871201234~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "PRV*ALBUQUERQUE*NM~", seg.String())

		seg.SetFieldByIndex("01", "ALBUQUERQU")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "prv's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "PRV*ALBUQUERQU*NM*871201234~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse prv's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
