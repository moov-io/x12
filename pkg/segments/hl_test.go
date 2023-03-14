// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForHL(t *testing.T) {

	t.Run("parsing of hl segment", func(t *testing.T) {

		seg := NewHL(nil)

		in := "HL*HC*85-0858585**20130709~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "HL*HC*85-0858585**20130709*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "HL*HC*85-0858585**~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "HL*HC*85-0858585*~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse hl's element (04), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "HL"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "hl segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "hl segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of hl segment", func(t *testing.T) {

		seg := NewHL(nil)

		require.Equal(t, "HL****~", seg.String())

		in := "HL*HC*85-0858585**20130709~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of hl segment with specified rule", func(t *testing.T) {

		rule := rules.Elements{
			"01": {AcceptValues: []string{"HC"}},
			"02": {Mask: rules.MASK_OPTIONAL},
			"03": {Mask: rules.MASK_NOTUSED},
			"04": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewHL(&rule)

		in := "HL*HC*85-0858585**20130709~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "HL*HC*85-0858585~", seg.String())

		seg.SetFieldByIndex("01", "HT")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "hl's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "HL*HT*85-0858585**20130709~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse hl's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
