// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForHealthCareServiceLocation(t *testing.T) {

	rule := rules.ElementSetRule{
		"02": {AcceptValues: []string{"2"}},
		"03": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
	}

	t.Run("parsing of service location", func(t *testing.T) {

		composite := HealthCareServiceLocation{}
		composite.SetRule(&rule)

		in := "composite:2:1"
		read, err := composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, composite.String())

		in = "composite:2:9"
		read, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse service location's element (03), the element contains unexpected value", err.Error())
	})

}

func TestForRelatedCausesInformation(t *testing.T) {

	rule := rules.ElementSetRule{
		"01": {AcceptValues: []string{"AA", "EM", "OA"}},
		"02": {Mask: rules.MASK_OPTIONAL, AcceptValues: []string{"AA", "EM", "OA"}},
		"03": {Mask: rules.MASK_NOTUSED, AcceptValues: []string{"AA", "EM", "OA"}},
		"04": {Mask: rules.MASK_OPTIONAL},
		"05": {Mask: rules.MASK_OPTIONAL},
	}

	t.Run("parsing of causes information", func(t *testing.T) {

		composite := HealthCareServiceLocation{}
		composite.SetRule(&rule)

		in := "AA"
		read, err := composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, composite.String())

		in = "AA::OA"
		read, err = composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "AA", composite.String())

		in = "CC"
		_, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse service location's element (01), the element contains unexpected value", err.Error())

		in = "AA:CC"
		_, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse service location's element (02), the element contains unexpected value", err.Error())

		in = "AA:AA:CC"
		_, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse service location's element (03), the element contains unexpected value", err.Error())
	})

}
