// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/stretchr/testify/require"
)

func TestForISA(t *testing.T) {

	t.Run("parsing of isa segment", func(t *testing.T) {

		seg := NewISA(nil)

		in := "ISA*00* *00* *ZZ*85-0858585 *ZZ* *130709*1058*^*00501*000101654*1*P*:~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "ISA*00* *00* *ZZ*85-0858585 *ZZ* *130709*1058*^*00501*000101654*1*P*:*~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in)-1, read)

		in = "ISA*00* *00* *ZZ*85-0858585 *ZZ* *130709*1058*^*00501*000101654*1**~"
		read, err = seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)

		in = "ISA*00* *00* *ZZ*85-0858585 *ZZ* *130709*1058*^*00501*000101654*1*P~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse isa's element (16), doesn't enough input string", err.Error())
		require.Equal(t, 0, read)

		in = "IEA*"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "isa segment has not enough input data", err.Error())
		require.Equal(t, 0, read)

		in = "IEA~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "isa segment contains invalid code", err.Error())
		require.Equal(t, 0, read)
	})

	t.Run("encoding of isa segment", func(t *testing.T) {

		seg := NewISA(nil)

		require.Equal(t, "ISA****************~", seg.String())

		in := "ISA*00* *00* *ZZ*85-0858585 *ZZ* *130709*1058*^*00501*000101654*1*P*:~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, in, seg.String())

		require.NoError(t, seg.Validate(nil))
	})

	t.Run("parsing and encoding of isa segment with specified rule", func(t *testing.T) {

		rule := rules.ElementSetRule{
			"01": {AcceptValues: []string{"00"}},
			"14": {Mask: rules.MASK_OPTIONAL},
			"15": {Mask: rules.MASK_NOTUSED},
			"16": {Mask: rules.MASK_NOTUSED},
		}

		seg := NewISA(&rule)

		in := "ISA*00* *00* *ZZ*85-0858585 *ZZ* *130709*1058*^*00501*000101654*1*P*:~"
		read, err := seg.Parse(in)
		require.NoError(t, err)
		require.Equal(t, len(in), read)
		require.Equal(t, "ISA*00* *00* *ZZ*85-0858585 *ZZ* *130709*1058*^*00501*000101654*1~", seg.String())

		seg.SetFieldByIndex("01", "01")
		err = seg.Validate(nil)
		require.Error(t, err)
		require.Equal(t, "isa's element (01) has invalid value, the element contains unexpected value", err.Error())

		in = "ISA*01* *00* *ZZ*85-0858585 *ZZ* *130709*1058*^*00501*000101654*1*P*:~"
		read, err = seg.Parse(in)
		require.Error(t, err)
		require.Equal(t, "unable to parse isa's element (01), the element contains unexpected value", err.Error())
		require.Equal(t, 0, read)
	})
}
