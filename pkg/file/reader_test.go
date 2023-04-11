// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {

	t.Run("testing scanner 1", func(t *testing.T) {

		raw := `ISA SECTION1 IEA*~`

		f := strings.NewReader(raw)
		scan := NewScanner(f)
		for line := scan.GetInterchange(); line != ""; line = scan.GetInterchange() {
			require.Equal(t, raw, line)
		}
	})

	t.Run("testing scanner 2", func(t *testing.T) {

		raw := `ISA SECTION1 
~ISA SECTION1*~`

		f := strings.NewReader(raw)
		scan := NewScanner(f)
		for line := scan.GetInterchange(); line != ""; line = scan.GetInterchange() {
			require.Equal(t, "ISA SECTION1 ~", line)
		}
	})

	t.Run("testing scanner 3", func(t *testing.T) {

		raw := `ISA SECTION1 
~IEA SECTION1*~`

		f := strings.NewReader(raw)
		scan := NewScanner(f)
		for line := scan.GetInterchange(); line != ""; line = scan.GetInterchange() {
			require.Equal(t, "ISA SECTION1 ~IEA SECTION1*~", line)
		}
	})

	t.Run("testing scanner 4", func(t *testing.T) {

		raw := `ERROR ISA SECTION1 
~IEA SECTION1*~`

		f := strings.NewReader(raw)
		scan := NewScanner(f)
		for line := scan.GetInterchange(); line != ""; line = scan.GetInterchange() {
			require.Equal(t, "ISA SECTION1 ~IEA SECTION1*~", line)
		}
	})
}
