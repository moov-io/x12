// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForDTP(t *testing.T) {

	t.Run("parsing of dtp segment", func(t *testing.T) {

		seg := NewDTP(nil)

		in := "DTP*0019*00*101654~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "DTP*0019*00*101654*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "DTP*0019*00*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "DTP*0019*00~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse dtp's element (03), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "dtp segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "dtp segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of dtp segment", func(t *testing.T) {

		seg := NewDTP(nil)

		require.Equal(t, "DTP***~", seg.String())

		in := "DTP*0019*00*101654~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of dtp segment with specified rule", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"0019"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewDTP(&rule)

		in := "DTP*0019*00*101654~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "DTP*0019*00~", seg.String())

		seg.SetFieldByIndex("01", "0018")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "dtp's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "DTP*0018*00*101654~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse dtp's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
