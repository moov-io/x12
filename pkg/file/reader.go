// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bufio"
	"io"
	"strings"

	"github.com/moov-io/x12/pkg/util"
)

type Scanner struct {
	scan *bufio.Scanner
}

func (b *Scanner) GetInterChange() string {
	if !b.scan.Scan() {
		return ""
	}

	raw := b.scan.Text()

	// Removing all of new lines
	raw = strings.ReplaceAll(raw, "\r\n", "")
	raw = strings.ReplaceAll(raw, "\r", "")
	raw = strings.ReplaceAll(raw, "\n", "")

	startPos := strings.Index(raw, "ISA")
	if startPos > 0 {
		raw = raw[startPos:]
	}

	return strings.TrimSpace(raw)
}

func NewScanner(fd io.Reader) Scanner {
	scan := bufio.NewScanner(fd)
	scan.Split(scanInterChange)
	return Scanner{scan: scan}
}

func getInterChangeTerminatorPosition(input string) int {

	startPos1 := strings.Index(input, "ISA")
	startPos2 := strings.LastIndex(input, "ISA")

	if startPos1 < 0 && startPos2 < 0 {
		return 0
	}

	if (startPos1 != -1 && startPos2 != -1) && (startPos1 != startPos2) {
		return startPos2
	}

	endPos := strings.Index(input, "IEA")
	if endPos == -1 {
		return 0
	}

	tPos := strings.Index(input[endPos:], util.SegmentTerminator)
	if tPos == -1 {
		return 0
	}

	return tPos + endPos + 1
}

func scanInterChange(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	line := string(data)
	pos := getInterChangeTerminatorPosition(line)
	if getInterChangeTerminatorPosition(line) < 0 || !atEOF {
		// need more data
		return 0, nil, nil
	}

	return pos, data[:pos], nil
}
