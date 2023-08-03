// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForOI(t *testing.T) {
	t.Run("parsing of oi segment", func(t *testing.T) {
		seg := NewOI(nil)

		in := "OI***0019***00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "OI***0019***00*~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "oi segment can't parse all input data", err.Error())
		require.Equal(t, 0, read)

		in = "OI*0019~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse oi's element (03), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "oi segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "oi segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of oi segment", func(t *testing.T) {
		seg := NewOI(nil)

		require.Equal(t, "OI******~", seg.String())

		in := "OI***0019***00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of oi segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"03": {AcceptValues: []string{"0019"}},
			"06": {AcceptValues: []string{"00"}},
		}

		seg := NewOI(&rule)

		in := "OI***0019***00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		seg.SetFieldByIndex("06", "01")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "oi's element (06) has invalid value, the element contains unexpected value", err.Error())

		seg.SetFieldByIndex("03", "0000")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "oi's element (03) has invalid value, the element contains unexpected value", err.Error())
	})
}
