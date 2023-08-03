// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForPLB(t *testing.T) {
	t.Run("parsing of PLB segment", func(t *testing.T) {
		seg := NewPLB(nil)

		in := "PLB*NPI Number*20111231*50>111*6173.4*50~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PLB*NPI Number*20111231*50>111*6173.4*50**********~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "plb segment can't parse all input data", err.Error())
		require.Equal(t, 0, read)

		in = "PLB*NPI Number*20111231*50>111*6173.4*50~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PLB*NPI Number*20111231*50>111~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse plb's element (04), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "plb segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "plb segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of PLB segment", func(t *testing.T) {
		seg := NewPLB(nil)

		require.Equal(t, "PLB****~", seg.String())

		in := "PLB*NPI Number*20111231*50<111*6173.4*50~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
	})

	t.Run("parsing and encoding of PLB segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {Mask: rules.MASK_REQUIRED},
			"02": {Mask: rules.MASK_REQUIRED},
			"03": {
				Mask: rules.MASK_REQUIRED,
				Composite: rules.ElementSetRule{
					"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"2"}},
					"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
				},
			},
			"04": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"6173.4"}},
		}

		seg := NewPLB(&rule)

		in := "PLB*NPI Number*20111231*2:1*6173.4~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "PLB*NPI Number*20111231*2:1*6173.4~", seg.String())

		seg.SetFieldByIndex("04", "W")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "plb's element (04) has invalid value, the element contains unexpected value", err.Error())
	})
}
