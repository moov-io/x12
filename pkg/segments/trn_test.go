// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForTRN(t *testing.T) {
	t.Run("parsing of trn segment", func(t *testing.T) {
		seg := NewTRN(nil)

		in := "TRN*85*2*INDIAN HEALTH HOSPITAL*~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "TRN*85*2*INDIAN HEALTH HOSPITAL**~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "trn segment can't parse all input data", err.Error())
		require.Equal(t, 0, read)

		in = "TRN*85*2*INDIAN HEALTH HOSPITAL~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "TRN"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "trn segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "NMN~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "trn segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of trn segment", func(t *testing.T) {
		seg := NewTRN(nil)

		require.Equal(t, "TRN**~", seg.String())

		in := "TRN*85*2~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of trn segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"85"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewTRN(&rule)

		in := "TRN*85*2*INDIAN HEALTH HOSPITAL*~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "TRN*85*2*INDIAN HEALTH HOSPITAL~", seg.String())

		seg.SetFieldByIndex("01", "86")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "trn's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "TRN*86*2*INDIAN HEALTH HOSPITAL~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse trn's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
