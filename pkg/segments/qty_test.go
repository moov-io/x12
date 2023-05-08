// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestForQTY(t *testing.T) {

	t.Run("parsing of QTY segment", func(t *testing.T) {

		seg := NewQTY(nil)

		in := "QTY**~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "QTY***~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "qty segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of QTY segment", func(t *testing.T) {

		seg := NewQTY(nil)

		require.Equal(t, "QTY**~", seg.String())

		in := "QTY*11:B:1*~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		read, err = seg.Parse(in, util.SegmentTerminator, "<")
		require.NoError(t, err)
		require.Equal(t, len(in), read)
	})

}
