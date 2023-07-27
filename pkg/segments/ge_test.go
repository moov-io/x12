// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForGE(t *testing.T) {
	t.Run("parsing of ge segment", func(t *testing.T) {
		seg := NewGE(nil)

		in := "GE*3*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "GE*3*8*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "GE*3*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "GE*3~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse ge's element (02), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "GE"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "ge segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "ge segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of ge segment", func(t *testing.T) {
		seg := NewGE(nil)

		require.Equal(t, "GE**~", seg.String())

		in := "GE*3*~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of ge segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"5"}, Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewGE(&rule)

		in := "GE*5*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "GE*5~", seg.String())

		seg.SetFieldByIndex("01", "6")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "ge's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "GE*6*8~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse ge's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
