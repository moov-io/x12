// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/moov-io/x12/pkg/file"
	. "github.com/moov-io/x12/rules/rule_4010_810"
)

func main() {
	reader, err := os.Open(samplePath("sample.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	newChange := file.NewFile(&InterchangeRule, "~")

	if err = newChange.Parse(file.NewScanner(reader, "~")); err != nil {
		log.Fatal(err.Error())
		return
	}

	if err = newChange.Validate(); err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("   REGENERATED FILE   ")
	fmt.Println(strings.ReplaceAll(newChange.String(), "~", "~\n"))

	newChange.Print(os.Stdout)

	InterchangeRule.Print(os.Stdout, true)
}

// samplePath returns the path of a sample file. The example can run from
// the repository root (go run ./examples/example_4010_810) or from its own folder.
func samplePath(name string) string {
	path := filepath.Join("examples", "example_4010_810", name)
	if _, err := os.Stat(path); err == nil {
		return path
	}
	return name
}
