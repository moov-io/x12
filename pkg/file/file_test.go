// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	rule "github.com/moov-io/x12/rules/rule_5010_837p"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFile(t *testing.T) {

	raw := `ISA*00*          *00*          *ZZ*133052274      *ZZ*311279999      *120419*2125*^*00501*000002120*0*P*<~
GS*HC*133052274*311279999*20120419*212549*2120*X*005010X222A1~
ST*837*000000533*005010X222A1~
BHT*0019*00*000017120*20120419*212550*CH~
NM1*41*2*CS*****46*133052274~
PER*IC*CUSTOMER SOLUTIONS*TE*8008456592~
NM1*40*2*RECEIVER*****46*311279999~
HL*1**20*1~
NM1*85*2*BILLING PROVIDER*****XX*1215954003~
N3*1540 SUNDAY DRIVE~
N4*RALEIGH*NC*276076000~
REF*EI*569999999~
NM1*87*2~
N3*P O BOX 63146~
N4*CHARLOTTE*NC*282633146~
HL*2*1*22*0~
SBR*P*18*******CI~
NM1*IL*1*SMITH*JOHN****MI*505XXXX~
N3*2345 NORTH RD~
N4*SPRINGFIELD*NY*11111~
DMG*D8*19950101*F~
NM1*PR*2*BLUE CROSS*****PI*XX123A~
N3*PO BOX 78765~
N4*SOUTH BEACH*FL*23466~
CLM*1029353-03*553.96***11<B<1*Y*A*Y*Y~
HI*BK<6268*BF<78930~
NM1*DN*1*SMITH*BOBBY****XX*123456789~
REF*1G*T32418~
NM1*82*2*SMITH*****XX*123456789~
REF*G2*00228498B~
NM1*77*2*SMITH LANEY*****XX*123456789~
N3*1111 DRICE PLACE~
N4*SOMEWHERE*NY*11119~
LX*1~
SV1*HC<87086*31.96*UN*1***1~
DTP*472*D8*20070118~
LX*2~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
LX*3~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
LX*4~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
HL*3*1*22*0~
SBR*P*18*******CI~
NM1*IL*1*SMITH*SARAH****MI*543XXXX~
N3*1111 NORTH LEG RD~
N4*ABC PLACE*NY*11223~
DMG*D8*19950101*F~
NM1*PR*2*AMERIGROUP/MDGA MANAGED CARE*****PI*12CCCA~
N3*PO BOX 44344~
N4*BRIGHTON*NY*23466~
CLM*1029353-03*174***11<B<1*Y*A*Y*Y~
HI*BK<5990~
NM1*DN*1*SMITH*WILLIAM****XX*7876666666~
REF*1G*T12345~
NM1*82*2*SMITH*****XX*111212222~
REF*G2*00228498B~
NM1*77*2*SMITH LAB*****XX*123456789~
N3*1111 DRICE PLACE~
N4*PLACE CITY*NY*11234~
LX*1~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
HL*4*1*22*0~
SBR*P*18*******CI~
NM1*IL*1*SMITH*TEST****MI*505XXXX~
N3*1111 A ROAD~
N4*PLACE CITY*NY*11234~
DMG*D8*19950101*F~
NM1*PR*2*ETNA*****PI*9876~
N3*PO BOX 61010~
N4*VIRGINIA BEACH*VA*23466~
CLM*1-1180*174***11<B<1*Y*A*Y*Y*P~
HI*BK<3829~
NM1*DN*1*SMITH*WILLIAM****XX*123456789~
REF*1G*T12345~
NM1*82*2*SMITH*****XX*123456789~
REF*G2*00227448X~
NM1*77*2*SMITH LAB*****XX*123456789~
N3*1527 NORTH LEG~
N4*PLACE CITY*NY*11234~
LX*1~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
HL*5*1*22*0~
SBR*P*18*******CI~
NM1*IL*1*LAB*TEST****MI*111XXXX~
N3*1111 A ROAD~
N4*PLACE CITY*NY*11234~
DMG*D8*19880101*M~
NM1*PR*2*PRONTO INSURANCE 101*****PI*12345~
N3*PO BOX 61010~
N4*VIRGINIA BEACH*VA*23466~
CLM*1-1181*174***11<B<1*Y*A*Y*Y*P~
HI*BK<05671~
NM1*DN*1*SMITH*WILLIAM****XX*123456789~
REF*1G*T32418~
NM1*82*2*SMITH*****XX*123456789~
REF*G2*00227448X~
NM1*77*2*SMITH LAB*****XX*123456789~
N3*1111 DRICE PLACE~
N4*PLACE CITY*NY*11234~
LX*1~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
HL*6*1*22*0~
SBR*P*18*******CI~
NM1*IL*1*LAB*TEST****MI*111XXXX~
N3*1111 A ROAD~
N4*PLACE CITY*NY*11234~
DMG*D8*19880101*M~
NM1*PR*2*FAST SERVICE INSURANCE*****PI*12345~
N3*PO BOX 61010~
N4*VIRGINIA BEACH*VA*23466~
CLM*1-1182*174***11<B<1*Y*A*Y*Y*P~
HI*BK<05671~
NM1*DN*1*SMITH*WILLIAM****XX*123456789~
REF*1G*T32418~
NM1*82*2*SMITH*****XX*123456789~
REF*G2*00227448X~
NM1*77*2*SMITH LAB*****XX*123456789~
N3*1111 DRICE PLACE~
N4*PLACE CITY*NY*11234~
LX*1~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
HL*7*1*22*0~
SBR*P*18*******CI~
NM1*IL*1*BORAT*LAB****MI*222XXXX~
N3*1402 WALTON WAY~
N4*AUGUSTA*GA*30901~
DMG*D8*19750101*M~
NM1*PR*2*AMERIGROUP/MDGA MANAGED CARE*****PI*12345~
N3*PO BOX 61010~
N4*VIRGINIA BEACH*VA*23466~
CLM*1-1183*174***11<B<1*Y*A*Y*Y*P~
HI*BK<20891~
NM1*DN*1*SMITH*WILLIAM****XX*123456789~
REF*1G*T32418~
NM1*82*2*SMITH*****XX*123456789~
REF*G2*00227448X~
NM1*77*2*SMITH LAB*****XX*123456789~
N3*1111 DRICE PLACE~
N4*PLACE CITY*NY*11234~
LX*1~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
HL*8*1*22*0~
SBR*P*18*******CI~
NM1*IL*1*BORAT*LAB****MI*222XXXX~
N3*1402 WALTON WAY~
N4*AUGUSTA*GA*30901~
DMG*D8*19750101*M~
NM1*PR*2*AMERIGROUP/MDGA MANAGED CARE*****PI*12345~
N3*PO BOX 61010~
N4*VIRGINIA BEACH*VA*23466~
CLM*1029353-03*174***11<B<1*Y*A*Y*Y~
HI*BK<20891~
NM1*DN*1*SMITH*WILLIAM****XX*123456789~
REF*1G*T32418~
NM1*82*2*SMITH*****XX*123456789~
REF*G2*00227448X~
NM1*77*2*SMITH LAB*****XX*123456789~
N3*1111 DRICE PLACE~
N4*PLACE CITY*NY*11234~
LX*1~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
HL*9*1*22*0~
SBR*P*18*******CI~
NM1*IL*1*BORAT*LAB****MI*222XXXX~
N3*1402 WALTON WAY~
N4*AUGUSTA*GA*30901~
DMG*D8*19750101*M~
NM1*PR*2*AMERIGROUP/MDGA MANAGED CARE*****PI*12345~
N3*PO BOX 61010~
N4*VIRGINIA BEACH*VA*23466~
CLM*1029353-03*174***11<B<1*Y*A*Y*Y~
HI*BK<20891~
NM1*DN*1*SMITH*WILLIAM****XX*123456789~
REF*1G*T32418~
NM1*82*2*SMITH*****XX*123456789~
REF*G2*00227448X~
NM1*77*2*SMITH LAB*****XX*123456789~
N3*1111 DRICE PLACE~
N4*PLACE CITY*NY*11234~
LX*1~
SV1*HC<97530*174*UN*3***1<2~
DTP*472*D8*20070118~
SE*191*000000533~
GE*1*2120~
IEA*1*000002120~`

	t.Run("testing file", func(t *testing.T) {

		reader := strings.NewReader(raw)
		scan := NewScanner(reader)

		f := NewFile(&rule.InterchangeRule)
		err := f.Parse(scan)
		require.NoError(t, err)

		err = f.Validate()
		require.NoError(t, err)

		out := f.String()
		stripRaw := strings.ReplaceAll(raw, "\n", "")

		require.Equal(t, stripRaw, out)

		buf := new(bytes.Buffer)
		f.Print(buf)
		require.NotEqual(t, 0, len(buf.String()))
	})

}
