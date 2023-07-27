// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForSV5(t *testing.T) {
	t.Run("parsing of sv5 segment", func(t *testing.T) {
		seg := NewSV5(nil)

		in := "SV5*85*2*INDIAN HEALTH HOSPITAL****~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "SV5*85*2*INDIAN HEALTH HOSPITAL*****~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "SV5*85*2*INDIAN HEALTH HOSPITAL**~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse sv5's element (06), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "SV5"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "sv5 segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "NMN~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "sv5 segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of sv5 segment", func(t *testing.T) {
		seg := NewSV5(nil)

		require.Equal(t, "SV5******~", seg.String())

		in := "SV5*85*2*INDIAN HEALTH HOSPITAL***~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of sv5 segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"85"}},
		}

		seg := NewSV5(&rule)

		in := "SV5*85*2*INDIAN HEALTH HOSPITAL*****XX*7745613100~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, 36, read)
		require.Equal(t, "SV5*85*2*INDIAN HEALTH HOSPITAL***~", seg.String())

		seg.SetFieldByIndex("01", "86")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "sv5's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "SV5*86*2*INDIAN HEALTH HOSPITAL***~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse sv5's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
