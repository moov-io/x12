// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package util

import (
	"errors"
)

const (
	LOOP_1000A  = "1000A"
	LOOP_1000B  = "1000B"
	LOOP_2000A  = "2000A"
	LOOP_2000B  = "2000B"
	LOOP_2010AA = "2010AA"
	LOOP_2010AB = "2010AB"
	LOOP_2010BA = "2010BA"
	LOOP_2010BB = "2010BB"
	LOOP_2310A  = "2310A"
	LOOP_2310B  = "2310B"
	LOOP_2310C  = "2310C"
	LOOP_2330B  = "2330B"
	LOOP_2330C  = "2330C"
	LOOP_2420A  = "2420A"
	LOOP_2100A  = "2100A"
	LOOP_2100B  = "2100B"
	LOOP_2100C  = "2100C"
	LOOP_2100D  = "2100D"
)

func ValidateLoopCode(code string) error {

	codes := []string{
		LOOP_1000A,
		LOOP_1000B,
		LOOP_2000A,
		LOOP_2010AA,
		LOOP_2010AB,
		LOOP_2010BA,
		LOOP_2010BB,
		LOOP_2310A,
		LOOP_2310B,
		LOOP_2310C,
		LOOP_2330B,
		LOOP_2330C,
		LOOP_2420A,
		LOOP_2100A,
		LOOP_2100B,
		LOOP_2100C,
		LOOP_2100D,
	}

	for _, l := range codes {
		if code == l {
			return nil
		}
	}

	return errors.New("invalid loop code")
}

func ValidateLoopCodeForNM1(code string) error {

	codes := []string{
		LOOP_1000A,
		LOOP_1000B,
		LOOP_2010AA,
		LOOP_2010AB,
		LOOP_2010BA,
		LOOP_2010BB,
		LOOP_2310A,
		LOOP_2310B,
		LOOP_2310C,
		LOOP_2330B,
		LOOP_2330C,
		LOOP_2420A,
		LOOP_2100A,
		LOOP_2100B,
		LOOP_2100C,
		LOOP_2100D,
	}

	for _, l := range codes {
		if code == l {
			return nil
		}
	}

	return errors.New("invalid loop code for nm1 segment")
}

func ValidateLoopCodeForPer(code string) error {

	codes := []string{
		LOOP_1000A,
	}

	for _, l := range codes {
		if code == l {
			return nil
		}
	}

	return errors.New("invalid loop code for per segment")
}

func ValidateLoopCodeForHl(code string) error {

	codes := []string{
		LOOP_2000A,
	}

	for _, l := range codes {
		if code == l {
			return nil
		}
	}

	return errors.New("invalid loop code for hl segment")
}

func ValidateLoopCodeForPrv(code string) error {

	codes := []string{
		LOOP_2000A,
	}

	for _, l := range codes {
		if code == l {
			return nil
		}
	}

	return errors.New("invalid loop code for prv segment")
}
