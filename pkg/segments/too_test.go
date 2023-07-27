// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestForTOO(t *testing.T) {
	t.Run("parsing of too segment", func(t *testing.T) {
		seg := NewTOO(nil)

		in := "TOO***11:B:1~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "TOO***11:B:1*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "TOO~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "too segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of too segment", func(t *testing.T) {
		seg := NewTOO(nil)

		require.Equal(t, "TOO~", seg.String())

		in := "TOO***11:B:1~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		read, err = seg.Parse(in, util.SegmentTerminator, "<")
		require.NoError(t, err)
		require.Equal(t, len(in), read)
	})

}
