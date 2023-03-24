// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"github.com/moov-io/x12/pkg/rules"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForSVD(t *testing.T) {

	t.Run("parsing of svd segment", func(t *testing.T) {

		seg := NewSVD(nil)

		in := "SVD*CODE*AMOUNT*AA:B**A*~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "SVD*CODE*AMOUNT*AA:B**A**~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "SVD*CODE*AMOUNT*AA:B**A~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "SVD*CODE*AMOUNT*AA:B*~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse svd's element (05), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "svd segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "svd segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of SVD segment", func(t *testing.T) {

		seg := NewSVD(nil)

		require.Equal(t, "SVD*****~", seg.String())

		in := "SVD*****~"
		read, err := seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse svd's element (03), unable to parse procedure identifier's element (01), doesn't enough input string", err.Error())

		read, err = seg.Parse(in, "<")
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse svd's element (03), unable to parse procedure identifier's element (01), doesn't enough input string", err.Error())
	})

	t.Run("parsing and encoding of svd segment with specified rule", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"03": {
				Mask: rules.MASK_REQUIRED,
				Composite: rules.ElementSetRule{
					"01": {AcceptValues: []string{"A"}},
					"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
				},
			},
			"04": {Mask: rules.MASK_NOTUSED},
			"05": {AcceptValues: []string{"Y", "N"}},
			"06": {Mask: rules.MASK_OPTIONAL},
		}

		seg := NewSVD(&rule)

		in := "SVD*CODE*AMOUNT*A:8**Y~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		seg.SetFieldByIndex("05", "W")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "svd's element (05) has invalid value, the element contains unexpected value", err.Error())

		in = "SVD*CODE*AMOUNT*A:A**Y~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse svd's element (03), unable to parse procedure identifier's element (02), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})

}
