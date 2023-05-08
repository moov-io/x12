// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForServiceProcedure(t *testing.T) {

	rule := rules.ElementSetRule{
		"01": {AcceptValues: []string{"composite"}},
		"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
	}

	t.Run("parsing of dental service procedure", func(t *testing.T) {

		composite := ServiceProcedure{}
		composite.SetRule(&rule)

		in := "composite:1"
		read, err := composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, composite.String())

		in = "composite:9"
		read, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse dental procedure's element (02), the element contains unexpected value", err.Error())
	})

}

func TestForDentalServiceCode(t *testing.T) {

	rule := rules.ElementSetRule{
		"01": {AcceptValues: []string{"composite"}},
		"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
	}

	t.Run("parsing of dental service code", func(t *testing.T) {

		composite := DentalServiceCode{}
		composite.SetRule(&rule)

		in := "composite:1"
		read, err := composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, composite.String())

		in = "composite:9"
		read, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse dental code's element (02), the element contains unexpected value", err.Error())
	})

}

func TestForDentalServiceCodePointer(t *testing.T) {

	rule := rules.ElementSetRule{
		"01": {AcceptValues: []string{"composite"}},
		"02": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
	}

	t.Run("parsing of dental service code pointer", func(t *testing.T) {

		composite := DentalServiceCodePointer{}
		composite.SetRule(&rule)

		in := "composite:1"
		read, err := composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, composite.String())

		in = "composite:9"
		read, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse dental code pointer's element (02), the element contains unexpected value", err.Error())
	})

}
