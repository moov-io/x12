// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForBHT(t *testing.T) {

	t.Run("parsing of bht segment", func(t *testing.T) {

		seg := NewBHT(nil)

		in := "BHT*0019*00*101654*20130709*1058*CH~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "BHT*0019*00*101654*20130709*1058*CH*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "BHT*0019*00*101654*20130709**~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "BHT*0019*00*101654*20130709*1058~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse bht's element (06), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "BH*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "bht segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "BHA~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "bht segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of bht segment", func(t *testing.T) {

		seg := NewBHT(nil)

		require.Equal(t, "BHT******~", seg.String())

		in := "BHT*0019*00*101654*20130709*1058*CH~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of bht segment with specified rule", func(t *testing.T) {

		rule := rules.Elements{
			"01": {AcceptValues: []string{"0019"}},
			"04": {Mask: rules.MASK_OPTIONAL},
			"05": {Mask: rules.MASK_NOTUSED},
			"06": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewBHT(&rule)

		in := "BHT*0019*00*101654*20130709*1058*CH~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "BHT*0019*00*101654*20130709~", seg.String())

		seg.SetFieldByIndex("01", "0018")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "bht's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "BHT*0018*00*101654*20130709*1058*CH~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse bht's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
