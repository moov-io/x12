// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForREF(t *testing.T) {

	t.Run("parsing of ref segment", func(t *testing.T) {

		seg := NewREF(nil)

		in := "REF*3*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "REF*3*8*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "REF*3*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "REF*3~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse ref's element (02), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "REF"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "ref segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "ref segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of ref segment", func(t *testing.T) {

		seg := NewREF(nil)

		require.Equal(t, "REF**~", seg.String())

		in := "REF*3*~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of ref segment with specified rule", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"5"}, Mask: rules.MASK_OPTIONAL},
			"02": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewREF(&rule)

		in := "REF*5*8~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "REF*5~", seg.String())

		seg.SetFieldByIndex("01", "6")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "ref's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "REF*6*8~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse ref's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
