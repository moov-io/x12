// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestForSV2(t *testing.T) {

	t.Run("parsing of sv2 segment", func(t *testing.T) {

		seg := NewSV2(nil)

		in := "SV2*1-1180*174***11:B:1*Y*A*Y*Y*P~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "SV2*a*1-1180*174***11:B:1*Y*A*Y*Y*P*****~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, 34, read)

		in = "SV2*1-1180*174***11:B:1*Y*A*Y*Y~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "sv2 segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of SV2 segment", func(t *testing.T) {

		seg := NewSV2(nil)

		require.Equal(t, "SV2~", seg.String())

		in := "SV2*a*11<B<1********~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		read, err = seg.Parse(in, util.SegmentTerminator, "<")
		require.NoError(t, err)
		require.Equal(t, len(in), read)
	})

}
