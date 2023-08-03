// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/moov-io/x12/pkg/util"
)

type Scanner struct {
	scan          *bufio.Scanner
	terminator    string
	removeNewLine bool
}

func (b *Scanner) GetInterchange() string {
	if !b.scan.Scan() {
		return ""
	}

	raw := b.scan.Text()

	if b.removeNewLine {
		// Removing all of new lines
		raw = strings.ReplaceAll(raw, "\n", "")
		raw = strings.ReplaceAll(raw, "\r", "")
	}

	startPos := strings.Index(raw, "ISA")
	if startPos > 0 {
		raw = raw[startPos:]
	}

	return strings.TrimSpace(raw)
}

// ARGS
//
//	first: terminator
//	second: don't allow remove newline
func NewScanner(fd io.Reader, args ...string) Scanner { // init scan
	scan := bufio.NewScanner(fd)

	// init object
	scanner := Scanner{scan: scan, removeNewLine: true}
	if len(args) > 0 {
		scanner.terminator = args[0]
	} else {
		scanner.terminator = util.SegmentTerminator
	}

	if len(args) > 1 {
		dontAllow, _ := strconv.ParseBool(args[1])
		scanner.removeNewLine = !dontAllow
	}

	// set split function
	scan.Split(scanner.scanInterChange)

	return scanner
}

func (b *Scanner) getInterchangeTerminatorPosition(input string, atEOF bool) int {
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

	tPos := strings.Index(input[endPos:], b.terminator)
	if tPos == -1 {
		if atEOF {
			return len(input)
		}
		return 0
	}

	return tPos + endPos + 1
}

func (b *Scanner) scanInterChange(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	line := string(data)
	pos := b.getInterchangeTerminatorPosition(line, atEOF)
	if b.getInterchangeTerminatorPosition(line, atEOF) < 0 || !atEOF {
		// need more data
		return 0, nil, nil
	}

	return pos, data[:pos], nil
}
