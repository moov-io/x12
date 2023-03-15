// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForProcedureIdentifier(t *testing.T) {

	rule := rules.Elements{
		"02": {AcceptValues: []string{"2"}},
		"03": {AcceptValues: []string{"1", "2", "3", "4", "5", "7", "8"}},
	}

	t.Run("parsing of service location", func(t *testing.T) {

		composite := ProcedureIdentifier{}
		composite.SetRule(&rule)

		in := "composite:2:1"
		read, err := composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, composite.String())

		in = "composite:2:1:::::"
		read, err = composite.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "composite:2:1", composite.String())

		in = "composite<2<1"
		read, err = composite.Parse(in, "<")
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "composite<2<1", composite.String("<"))

		in = "composite:2:9"
		read, err = composite.Parse(in)
		require.Error(t, err)
		require.Equal(t, 0, read)
		require.Equal(t, "unable to parse procedure identifier's element (03), the element contains unexpected value", err.Error())
	})

}
