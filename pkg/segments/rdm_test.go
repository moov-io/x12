// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForRDM(t *testing.T) {
	t.Run("parsing of rdm segment", func(t *testing.T) {
		seg := NewRDM(nil)

		in := "RDM*3*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "RDM*3*8*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "RDM*3*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "RDM*3~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse rdm's element (02), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "RDM"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "rdm segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GTa~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "rdm segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of RDM segment", func(t *testing.T) {
		seg := NewRDM(nil)

		require.Equal(t, "RDM**~", seg.String())

		in := "RDM*3*~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of rdm segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"5"}, Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewRDM(&rule)

		in := "RDM*5*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "RDM*5~", seg.String())

		seg.SetFieldByIndex("01", "6")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "rdm's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "RDM*6*8~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse rdm's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
