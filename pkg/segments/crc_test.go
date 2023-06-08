// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForCRC(t *testing.T) {

	t.Run("parsing of crc segment", func(t *testing.T) {

		seg := NewCRC(nil)

		in := "CRC*19~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "CRC*19*******~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "CRC*19*****~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "CRC"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "crc segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "PAA~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "crc segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of crc segment", func(t *testing.T) {

		seg := NewCRC(nil)

		require.Equal(t, "CRC~", seg.String())

		in := "CRC*19********~"
		_, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, "CRC*19~", seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of crc segment with specified rule", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"19"}, Mask: rules.MASK_REQUIRED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_NOTUSED},
			"10": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewCRC(&rule)

		in := "CRC*19******~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "CRC*19~", seg.String())

		seg.SetFieldByIndex("01", "20")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "crc's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "CRC*20********~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse crc's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
