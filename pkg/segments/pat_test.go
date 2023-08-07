// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForPAT(t *testing.T) {
	t.Run("parsing of pat segment", func(t *testing.T) {
		seg := NewPAT(nil)

		in := "PAT*19~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PAT*19*********~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "pat segment can't parse all input data", err.Error())
		require.Equal(t, 0, read)

		in = "PAT*19********~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PAT"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "pat segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "PAA~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "pat segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of pat segment", func(t *testing.T) {
		seg := NewPAT(nil)

		require.Equal(t, "PAT~", seg.String())

		in := "PAT*19********~"
		_, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, "PAT*19~", seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of pat segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"19"}, Mask: rules.MASK_REQUIRED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewPAT(&rule)

		in := "PAT*19********~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "PAT*19~", seg.String())

		seg.SetFieldByIndex("01", "20")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "pat's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "PAT*20********~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse pat's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
