// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForCreateSegment(t *testing.T) {

	seg, err := CreateSegment("invalid", nil)
	require.Error(t, err)
	require.Equal(t, "unsupported segment name", err.Error())
	require.Nil(t, seg)

	seg, err = CreateSegment("AMT", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "AMT", seg.Name())

	seg, err = CreateSegment("BHT", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "BHT", seg.Name())

	seg, err = CreateSegment("CAS", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CAS", seg.Name())

	seg, err = CreateSegment("CLM", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CLM", seg.Name())

	seg, err = CreateSegment("CR1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CR1", seg.Name())

	seg, err = CreateSegment("DMG", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "DMG", seg.Name())

	seg, err = CreateSegment("DTP", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "DTP", seg.Name())

	seg, err = CreateSegment("GE", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "GE", seg.Name())

	seg, err = CreateSegment("GS", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "GS", seg.Name())

	seg, err = CreateSegment("HI", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "HI", seg.Name())

	seg, err = CreateSegment("HL", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "HL", seg.Name())

	seg, err = CreateSegment("IEA", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "IEA", seg.Name())

	seg, err = CreateSegment("ISA", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "ISA", seg.Name())

	seg, err = CreateSegment("MOA", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "MOA", seg.Name())

	seg, err = CreateSegment("N1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "N1", seg.Name())

	seg, err = CreateSegment("N2", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "N2", seg.Name())

	seg, err = CreateSegment("N3", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "N3", seg.Name())

	seg, err = CreateSegment("N4", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "N4", seg.Name())

	seg, err = CreateSegment("NM1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "NM1", seg.Name())

	seg, err = CreateSegment("NTE", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "NTE", seg.Name())

	seg, err = CreateSegment("OI", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "OI", seg.Name())

	seg, err = CreateSegment("PAT", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PAT", seg.Name())

	seg, err = CreateSegment("PER", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PER", seg.Name())

	seg, err = CreateSegment("PRV", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PRV", seg.Name())

	seg, err = CreateSegment("PWK", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PWK", seg.Name())

	seg, err = CreateSegment("REF", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "REF", seg.Name())

	seg, err = CreateSegment("RMR", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "RMR", seg.Name())

	seg, err = CreateSegment("SBR", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SBR", seg.Name())

	seg, err = CreateSegment("SE", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SE", seg.Name())

	seg, err = CreateSegment("ST", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "ST", seg.Name())

	seg, err = CreateSegment("LX", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "LX", seg.Name())

	seg, err = CreateSegment("SV1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SV1", seg.Name())

	seg, err = CreateSegment("SV5", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SV5", seg.Name())

	seg, err = CreateSegment("SVD", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SVD", seg.Name())

	seg, err = CreateSegment("CN1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CN1", seg.Name())

	seg, err = CreateSegment("CUR", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CUR", seg.Name())

	seg, err = CreateSegment("DN1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "DN1", seg.Name())

	seg, err = CreateSegment("DN2", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "DN2", seg.Name())

	seg, err = CreateSegment("SV3", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SV3", seg.Name())

	seg, err = CreateSegment("TOO", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "TOO", seg.Name())

	seg, err = CreateSegment("TRN", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "TRN", seg.Name())
}
