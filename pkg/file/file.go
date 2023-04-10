// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package file

import (
	"bytes"
	"errors"
	"io"

	"github.com/moov-io/x12/pkg/rules"
)

func NewFile(rule *rules.InterChangeRule) *File {

	f := File{rule: rule}

	return &f
}

type File struct {
	Interchanges []Interchange

	rule *rules.InterChangeRule
}

func (f *File) Validate() error {

	if len(f.Interchanges) == 0 {
		return nil
	}

	if f.rule == nil {
		return errors.New("unable to find valid rule")
	}

	for _, change := range f.Interchanges {
		if err := change.Validate(f.rule); err != nil {
			return err
		}
	}

	return nil
}

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
		return errors.New("unable to find valid rule")
	}

	for line := scan.GetInterChange(); line != ""; line = scan.GetInterChange() {
		newChange := NewInterchange(f.rule)
		read, err := newChange.Parse(line)
		if err == nil && read > 0 {
			f.Interchanges = append(f.Interchanges, *newChange)
		}
	}

	return nil
}

func (f File) String() string {
	var buf bytes.Buffer

	for index := range f.Interchanges {
		buf.WriteString(f.Interchanges[index].String())
	}

	return buf.String()
}
