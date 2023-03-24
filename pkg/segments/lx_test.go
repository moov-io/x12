// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForLX(t *testing.T) {

	t.Run("parsing of lx segment", func(t *testing.T) {

		seg := NewLX(nil)

		in := "LX*3~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "LX*3*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "LX*3~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "LX~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse lx's element (01), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "LX"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "lx segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "lx segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of lx segment", func(t *testing.T) {

		seg := NewLX(nil)

		require.Equal(t, "LX*~", seg.String())

		in := "LX*3~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of lx segment with specified rule", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"5"}, Mask: rules.MASK_OPTIONAL},
		}

		seg := NewLX(&rule)

		in := "LX*5~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "LX*5~", seg.String())

		seg.SetFieldByIndex("01", "6")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "lx's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "LX*6*8~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse lx's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
