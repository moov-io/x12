// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForGS(t *testing.T) {
	t.Run("parsing of gs segment", func(t *testing.T) {
		seg := NewGS(nil)

		in := "GS*HC*85-0858585**20130709*1058*101654*X*005010X222A1~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "GS*HC*85-0858585**20130709*1058*101654*X*005010X222A1*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "GS*HC*85-0858585**20130709*1058*101654**~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "GS*HC*85-0858585**20130709*1058*101654*X~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse gs's element (08), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "GS"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "gs segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "gs segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of gs segment", func(t *testing.T) {
		seg := NewGS(nil)

		require.Equal(t, "GS********~", seg.String())

		in := "GS*HC*85-0858585**20130709*1058*101654*X*005010X222A1~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of gs segment with specified rule", func(t *testing.T) {
		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"HC"}},
			"06": {Mask: rules.MASK_OPTIONAL},
			"07": {Mask: rules.MASK_NOTUSED},
			"08": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewGS(&rule)

		in := "GS*HC*85-0858585**20130709*1058*101654*X*005010X222A1~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "GS*HC*85-0858585**20130709*1058*101654~", seg.String())

		seg.SetFieldByIndex("01", "HT")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "gs's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "GS*HT*85-0858585**20130709*1058*101654*X*005010X222A1~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse gs's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
