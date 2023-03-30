// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForSV3(t *testing.T) {

	t.Run("parsing of sv3 segment", func(t *testing.T) {

		seg := NewSV3(nil)

		in := "SV3*1-1180*174***11:B:1*Y*A*Y*Y*P~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "SV3*1-1180*174***11:B:1*Y*A*Y*Y*P*********~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "SV3*1-1180*174***11:B:1*Y*A*Y*Y~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "sv3 segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of SV3 segment", func(t *testing.T) {

		seg := NewSV3(nil)

		require.Equal(t, "SV3~", seg.String())

		in := "SV3*11<B<1****11<B<1******11<B<1~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		read, err = seg.Parse(in, "<")
		require.NoError(t, err)
		require.Equal(t, len(in), read)
	})

}
