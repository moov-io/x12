// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForNM1(t *testing.T) {
	t.Run("parsing of nm1 segment", func(t *testing.T) {
		seg := NewNM1(nil)

		in := "NM1*85*2*INDIAN HEALTH HOSPITAL*****XX*7745613100~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "NM1*85*2*INDIAN HEALTH HOSPITAL*****XX*7745613100*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "NM1*85*2*INDIAN HEALTH HOSPITAL******~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "NM1*85*2*INDIAN HEALTH HOSPITAL*****XX~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse nm1's element (09), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "NM1"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "nm1 segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "NMN~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "nm1 segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of nm1 segment", func(t *testing.T) {
		seg := NewNM1(nil)

		require.Equal(t, "NM1*********~", seg.String())

		in := "NM1*85*2*INDIAN HEALTH HOSPITAL*****XX*7745613100~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of nm1 segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"85"}},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewNM1(&rule)

		in := "NM1*85*2*INDIAN HEALTH HOSPITAL*****XX*7745613100~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "NM1*85*2*INDIAN HEALTH HOSPITAL*****XX~", seg.String())

		seg.SetFieldByIndex("01", "86")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "nm1's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "NM1*86*2*INDIAN HEALTH HOSPITAL*****XX*7745613100~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse nm1's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
