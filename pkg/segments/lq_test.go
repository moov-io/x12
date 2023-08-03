// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestForLQ(t *testing.T) {
	t.Run("parsing of LQ segment", func(t *testing.T) {
		seg := NewLQ(nil)

		in := "LQ**~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "LQ***~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "lq segment can't parse all input data", err.Error())
		require.Equal(t, 0, read)

		in = "DT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "lq segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of LQ segment", func(t *testing.T) {
		seg := NewLQ(nil)

		require.Equal(t, "LQ**~", seg.String())

		in := "LQ*11:B:1*~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		read, err = seg.Parse(in, util.SegmentTerminator, "<")
		require.NoError(t, err)
		require.Equal(t, len(in), read)
	})

}
