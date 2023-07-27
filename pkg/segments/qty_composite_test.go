// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestFoQtyComposite(t *testing.T) {
	rule := rules.ElementSetRule{
		"01": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"AA", "EM", "OA"}},
		"02": {Mask: rules.MASK_REQUIRED, AcceptValues: []string{"AA", "EM", "OA"}},
	}

	t.Run("parsing of qty composite", func(t *testing.T) {
		composite := QtyComposite{}
		composite.SetRule(&rule)

		in := "AA:AA"
		read, err := composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, composite.String())

		in = "AA:OA"
		read, err = composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "AA:OA", composite.String())

		in = "CC"
		_, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse qty composite's element (01), the element contains unexpected value", err.Error())

		in = "AA:CC"
		_, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse qty composite's element (02), the element contains unexpected value", err.Error())

		in = "AA:AA"
		_, err = composite.Parse(in)
		require.NoError(t, err)
	})

}
