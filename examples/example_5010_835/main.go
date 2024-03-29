// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/moov-io/x12/pkg/file"
	. "github.com/moov-io/x12/rules/rule_5010_835"
)

func main() {
	reader, err := os.Open(path.Join("examples", "example_5010_835", "sample2.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	newChange := file.NewFile(&InterchangeRule)

	if err = newChange.Parse(file.NewScanner(reader)); err != nil {
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
