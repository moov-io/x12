// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForAMT(t *testing.T) {
	t.Run("parsing of amt segment", func(t *testing.T) {
		seg := NewAMT(nil)

		in := "AMT*0019*00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "AMT*0019*00*XX~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-3, read)

		in = "AMT*0019*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "AMT*0019~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse amt's element (02), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "amt segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "amt segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of amt segment", func(t *testing.T) {
		seg := NewAMT(nil)

		require.Equal(t, "AMT**~", seg.String())

		in := "AMT*0019*00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of amt segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"0019"}},
			"02": {Mask: rules.MASK_OPTIONAL},
		}

		seg := NewAMT(&rule)

		in := "AMT*0019*00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "AMT*0019*00~", seg.String())

		seg.SetFieldByIndex("01", "0018")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "amt's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "AMT*0018*00~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse amt's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
