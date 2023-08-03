// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package segments

import (
	"testing"

	"github.com/moov-io/x12/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestForCreateSegment(t *testing.T) {
	seg, err := CreateSegment("invalid", nil)
	require.Error(t, err)
	require.Equal(t, "unsupported segment name(invalid)", err.Error())
	require.Nil(t, seg)

	seg, err = CreateSegment("AMT", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "AMT", util.GetStructName(seg))

	seg, err = CreateSegment("BHT", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "BHT", util.GetStructName(seg))

	seg, err = CreateSegment("BIG", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "BIG", util.GetStructName(seg))

	seg, err = CreateSegment("BPR", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "BPR", util.GetStructName(seg))

	seg, err = CreateSegment("CAS", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CAS", util.GetStructName(seg))

	seg, err = CreateSegment("CLM", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CLM", util.GetStructName(seg))

	seg, err = CreateSegment("CLP", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CLP", util.GetStructName(seg))

	seg, err = CreateSegment("CR1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CR1", util.GetStructName(seg))

	seg, err = CreateSegment("CR6", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CR6", util.GetStructName(seg))

	seg, err = CreateSegment("CRC", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CRC", util.GetStructName(seg))

	seg, err = CreateSegment("DMG", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "DMG", util.GetStructName(seg))

	seg, err = CreateSegment("DTM", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "DTM", util.GetStructName(seg))

	seg, err = CreateSegment("DTP", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "DTP", util.GetStructName(seg))

	seg, err = CreateSegment("FA1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "FA1", util.GetStructName(seg))

	seg, err = CreateSegment("FA2", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "FA2", util.GetStructName(seg))

	seg, err = CreateSegment("GE", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "GE", util.GetStructName(seg))

	seg, err = CreateSegment("GS", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "GS", util.GetStructName(seg))

	seg, err = CreateSegment("HI", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "HI", util.GetStructName(seg))

	seg, err = CreateSegment("HL", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "HL", util.GetStructName(seg))

	seg, err = CreateSegment("IEA", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "IEA", util.GetStructName(seg))

	seg, err = CreateSegment("ISA", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "ISA", util.GetStructName(seg))

	seg, err = CreateSegment("IT1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "IT1", util.GetStructName(seg))

	seg, err = CreateSegment("ITD", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "ITD", util.GetStructName(seg))

	seg, err = CreateSegment("LQ", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "LQ", util.GetStructName(seg))

	seg, err = CreateSegment("LX", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "LX", util.GetStructName(seg))

	seg, err = CreateSegment("MIA", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "MIA", util.GetStructName(seg))

	seg, err = CreateSegment("MOA", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "MOA", util.GetStructName(seg))

	seg, err = CreateSegment("MSG", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "MSG", util.GetStructName(seg))

	seg, err = CreateSegment("N1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "N1", util.GetStructName(seg))

	seg, err = CreateSegment("N2", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "N2", util.GetStructName(seg))

	seg, err = CreateSegment("N3", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "N3", util.GetStructName(seg))

	seg, err = CreateSegment("N4", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "N4", util.GetStructName(seg))

	seg, err = CreateSegment("N9", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "N9", util.GetStructName(seg))

	seg, err = CreateSegment("NM1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "NM1", util.GetStructName(seg))

	seg, err = CreateSegment("NTE", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "NTE", util.GetStructName(seg))

	seg, err = CreateSegment("OI", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "OI", util.GetStructName(seg))

	seg, err = CreateSegment("PAT", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PAT", util.GetStructName(seg))

	seg, err = CreateSegment("PER", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PER", util.GetStructName(seg))

	seg, err = CreateSegment("PID", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PID", util.GetStructName(seg))

	seg, err = CreateSegment("PLB", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PLB", util.GetStructName(seg))

	seg, err = CreateSegment("PRV", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PRV", util.GetStructName(seg))

	seg, err = CreateSegment("PWK", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "PWK", util.GetStructName(seg))

	seg, err = CreateSegment("QTY", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "QTY", util.GetStructName(seg))

	seg, err = CreateSegment("RDM", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "RDM", util.GetStructName(seg))

	seg, err = CreateSegment("REF", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "REF", util.GetStructName(seg))

	seg, err = CreateSegment("RMR", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "RMR", util.GetStructName(seg))

	seg, err = CreateSegment("SAC", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SAC", util.GetStructName(seg))

	seg, err = CreateSegment("SBR", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SBR", util.GetStructName(seg))

	seg, err = CreateSegment("SE", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SE", util.GetStructName(seg))

	seg, err = CreateSegment("ST", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "ST", util.GetStructName(seg))

	seg, err = CreateSegment("LX", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "LX", util.GetStructName(seg))

	seg, err = CreateSegment("SV1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SV1", util.GetStructName(seg))

	seg, err = CreateSegment("SV2", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SV2", util.GetStructName(seg))

	seg, err = CreateSegment("SV3", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SV3", util.GetStructName(seg))

	seg, err = CreateSegment("SV5", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SV5", util.GetStructName(seg))

	seg, err = CreateSegment("SVC", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SVC", util.GetStructName(seg))

	seg, err = CreateSegment("SVD", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "SVD", util.GetStructName(seg))

	seg, err = CreateSegment("CN1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CN1", util.GetStructName(seg))

	seg, err = CreateSegment("CTP", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CTP", util.GetStructName(seg))

	seg, err = CreateSegment("CTT", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CTT", util.GetStructName(seg))

	seg, err = CreateSegment("CUR", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "CUR", util.GetStructName(seg))

	seg, err = CreateSegment("DN1", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "DN1", util.GetStructName(seg))

	seg, err = CreateSegment("DN2", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "DN2", util.GetStructName(seg))

	seg, err = CreateSegment("TDS", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "TDS", util.GetStructName(seg))

	seg, err = CreateSegment("TOO", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "TOO", util.GetStructName(seg))

	seg, err = CreateSegment("TRN", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "TRN", util.GetStructName(seg))

	seg, err = CreateSegment("TS2", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "TS2", util.GetStructName(seg))

	seg, err = CreateSegment("TS3", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "TS3", util.GetStructName(seg))

	seg, err = CreateSegment("TXI", nil)
	require.NoError(t, err)
	require.NotNil(t, seg)
	require.Equal(t, "TXI", util.GetStructName(seg))
}
