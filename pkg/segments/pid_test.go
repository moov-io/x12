// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForPID(t *testing.T) {
	t.Run("parsing of pid segment", func(t *testing.T) {
		seg := NewPID(nil)

		in := "PID*0019*00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PID*0019*00*XX~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PID*0019*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "pid segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "DMT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "pid segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of pid segment", func(t *testing.T) {
		seg := NewPID(nil)

		require.Equal(t, "PID~", seg.String())

		in := "PID*0019*00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of pid segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"0019"}},
			"02": {Mask: rules.MASK_OPTIONAL},
		}

		seg := NewPID(&rule)

		in := "PID*0019*00~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "PID*0019*00~", seg.String())

		seg.SetFieldByIndex("01", "0018")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "pid's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "PID*0018*00~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse pid's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
