// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"io"

	"github.com/moov-io/x12/pkg/rules"
	"github.com/moov-io/x12/pkg/util"
)

func NewFile(rule *rules.InterchangeRule, args ...string) *File {
	f := File{rule: rule}

	if len(args) > 0 && args[0] != "" {
		f.terminator = args[0]
	}

	return &f
}

type File struct {
	Interchanges []Interchange

	terminator string
	rule       *rules.InterchangeRule
}

func (f File) getTerminator() string {
	if f.terminator != "" {
		return f.terminator
	}

	return util.SegmentTerminator
}

func (f *File) Validate() error {
	if len(f.Interchanges) == 0 {
		return util.NewFindElementError("interchange")
	}

	if f.rule == nil {
		return util.NewFindRuleError("valid")
	}

	for _, change := range f.Interchanges {
		if err := change.Validate(f.rule); err != nil {
			return err
		}
	}

	return nil
}

//  The function will print human-readable format of edi file
//  User should be fund segments and loops with level, indexes are normal spec indexes
//
//  	DUMP EDI FILE WITH 837D(005010X224A2)
//
//	INDEX:       | 00           | 01               | 02                     | 03        | 04    | 05             | 06 | 07             | 08        | 09  | 10 | 11   | 12       | 13 | 14 | 15 | 16 | 17 | 18 |
//	ISA          |00            |                  |00                      |           |ZZ     |133052274       |ZZ  |311279999       |120419     |2125 |^   |00501 |000002120 |0   |P   |:   |~   |    |    |
//	 GS          |HC            |133052274         |311279999               |20120419   |212549 |2120            |X   |005010X224A2    |~          |     |    |      |          |    |    |    |    |    |    |
//	  ST         |837           |3456              |005010X224A2            |~          |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |
//	   1000A     |              |                  |                        |           |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |
//	    NM1      |41            |2                 |PREMIER BILLING SERVICE |           |       |                |    |46              |TGJ23      |~    |    |      |          |    |    |    |    |    |    |
//	    PER      |IC            |JERRY             |TE                      |7176149999 |       |                |    |                |~          |     |    |      |          |    |    |    |    |    |    |
//   ... ...
//	  SE         |31            |3456              |~                       |           |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |
//	 GE          |1             |2120              |~                       |           |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |
//	IEA          |1             |000002120         |~                       |           |       |                |    |                |           |     |    |      |          |    |    |    |    |    |    |

func (f File) Print(w io.Writer) {
	if f.rule == nil {
		w.Write([]byte("unable to find valid rule"))
		return
	}

	w.Write([]byte("\n  DUMP EDI FILE WITH " + f.rule.Name + "\n\n"))

	for _, change := range f.Interchanges {
		change.Print(w)
		w.Write([]byte("\n"))
	}
}

func (f *File) Parse(scan Scanner) error {
	if f.rule == nil {
		return util.NewFindRuleError("valid")
	}

	for line := scan.GetInterchange(); line != ""; line = scan.GetInterchange() {
		newChange := NewInterchange(f.rule, f.getTerminator())
		read, err := newChange.Parse(line)
		if err == nil && read > 0 {
			f.Interchanges = append(f.Interchanges, *newChange)
		} else if err != nil {
			return err
		}
	}

	return nil
}

func (f File) String() string {
	var buf bytes.Buffer

	for index := range f.Interchanges {
		buf.WriteString(f.Interchanges[index].String(f.getTerminator()))
	}

	return buf.String()
}
