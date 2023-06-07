// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForCR6(t *testing.T) {

	t.Run("parsing of cr6 segment", func(t *testing.T) {

		seg := NewCR6(nil)

		in := "CR6*19~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "CR6*19**********~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "CR6*19********~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "CR6"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "cr6 segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "PAA~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "cr6 segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of cr6 segment", func(t *testing.T) {

		seg := NewCR6(nil)

		require.Equal(t, "CR6~", seg.String())

		in := "CR6*19********~"
		_, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, "CR6*19~", seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of cr6 segment with specified rule", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"19"}, Mask: rules.MASK_REQUIRED},
			"08": {Mask: rules.MASK_OPTIONAL},
			"09": {Mask: rules.MASK_NOTUSED},
			"10": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewCR6(&rule)

		in := "CR6*19********~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "CR6*19~", seg.String())

		seg.SetFieldByIndex("01", "20")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "cr6's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "CR6*20********~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse cr6's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
