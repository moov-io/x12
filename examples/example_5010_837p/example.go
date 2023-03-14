// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/moov-io/x12/pkg/file"
	"github.com/moov-io/x12/rule_5010_837p"
)

func main() {

	f, err := os.Open(path.Join("examples", "example_5010_837p", "sample.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var raw string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		raw += strings.TrimSpace(scanner.Text())
	}

	newChange := file.NewInterchange(&rule_5010_837p.InterchangeRule)

	_, err = newChange.Parse(raw, "<")
	if err = newChange.Validate(nil); err != nil {
		fmt.Println(err)
		return
	}

	err = newChange.Validate(&rule_5010_837p.InterchangeRule)
	if err = newChange.Validate(nil); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("   REGENERATED FILE   ")
	fmt.Println(strings.ReplaceAll(newChange.String("<"), "~", "~\n"))
}
