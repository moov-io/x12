// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForPWK(t *testing.T) {

	t.Run("parsing of pwk segment", func(t *testing.T) {

		seg := NewPWK(nil)

		in := "PWK*IC*BUSINESS OFFICE*TE*5052484349~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PWK*IC*BUSINESS OFFICE*TE*5052484349***~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "PWK*IC*BUSINESS OFFICE*TE*5052484349~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "PWK*IC~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse pwk's element (02), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "SBR"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "pwk segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "GT~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "pwk segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of pwk segment", func(t *testing.T) {

		seg := NewPWK(nil)

		require.Equal(t, "PWK**~", seg.String())

		in := "PWK*IC*BUSINESS OFFICE*TE*5052484349~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of pwk segment with specified rule", func(t *testing.T) {

		rule := rules.Elements{
			"01": {AcceptValues: []string{"IC"}},
			"03": {Mask: rules.MASK_OPTIONAL},
			"04": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewPWK(&rule)

		in := "PWK*IC*BUSINESS OFFICE*TE*5052484349~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "PWK*IC*BUSINESS OFFICE*TE~", seg.String())

		seg.SetFieldByIndex("01", "B")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "pwk's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "PWK*B*BUSINESS OFFICE*TE*5052484349~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse pwk's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
