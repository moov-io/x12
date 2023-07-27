// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForCTT(t *testing.T) {
	t.Run("parsing of ctt segment", func(t *testing.T) {
		seg := NewCTT(nil)

		in := "CTT*0019*00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "CTT*0019*00*XX~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "CTT*0019*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "ctt segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "ctt segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of ctt segment", func(t *testing.T) {
		seg := NewCTT(nil)

		require.Equal(t, "CTT~", seg.String())

		in := "CTT*0019*00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of ctt segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"0019"}},
			"02": {Mask: rules.MASK_OPTIONAL},
		}

		seg := NewCTT(&rule)

		in := "CTT*0019*00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "CTT*0019*00~", seg.String())

		seg.SetFieldByIndex("01", "0018")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "ctt's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "CTT*0018*00~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse ctt's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
