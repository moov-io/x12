// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForST(t *testing.T) {
	t.Run("parsing of st segment", func(t *testing.T) {
		seg := NewST(nil)

		in := "ST*837*0001*005010X222A1~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "ST*837*0001*005010X222A1*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "ST*3*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "ST~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse st's element (01), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "ST"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "st segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "st segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of st segment", func(t *testing.T) {
		seg := NewST(nil)

		require.Equal(t, "ST*~", seg.String())

		in := "ST*3~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of st segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"5"}, Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewST(&rule)

		in := "ST*5*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "ST*5~", seg.String())

		seg.SetFieldByIndex("01", "6")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "st's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "ST*6*8~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse st's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
