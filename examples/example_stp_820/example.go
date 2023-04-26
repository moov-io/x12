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
	. "github.com/moov-io/x12/rules/rule_stp_820"
)

func main() {

	reader, err := os.Open(path.Join("examples", "example_stp_820", "sample.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	copyRule := InterchangeRule

	eRule := copyRule.Group.GS.Elements["08"]
	eRule.AcceptValues = []string{"004010STP820", "004010"}
	copyRule.Group.GS.Elements["08"] = eRule

	segmentTerminator := "\\"

	newChange := file.NewFile(&InterchangeRule, segmentTerminator)

	if err = newChange.Parse(file.NewScanner(reader, segmentTerminator)); err != nil {
		log.Fatal(err.Error())
		return
	}

	if err = newChange.Validate(); err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("   REGENERATED FILE   ")
	fmt.Println(strings.ReplaceAll(newChange.String(), segmentTerminator, segmentTerminator+"\n"))

	newChange.Print(os.Stdout)
}
