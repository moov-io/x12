// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestForSVC(t *testing.T) {
	t.Run("parsing of svc segment", func(t *testing.T) {
		seg := NewSVC(nil)

		in := "SVC*1-1180*174****11:B:1*~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "SVC*1-1180*174****11:B:1*Y*A*Y*Y*P*********~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, 27, read)

		in = "SVC*1-1180*174****11:B:1*Y*A*Y*Y~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, 27, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "svc segment contains invalid code", err.Error())
		require.Equal(t, 0, read)

	})

	t.Run("encoding of SVC segment", func(t *testing.T) {
		seg := NewSVC(nil)

		require.Equal(t, "SVC***~", seg.String())

		in := "SVC*11<B<1*****11<B<1~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		read, err = seg.Parse(in, util.SegmentTerminator, "<")
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		err = seg.Validate(nil)
		require.NoError(t, err)
	})

}
