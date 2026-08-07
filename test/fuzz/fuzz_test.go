// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzz

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/x12/pkg/file"
	"github.com/moov-io/x12/pkg/rules"
	rule4010810 "github.com/moov-io/x12/rules/rule_4010_810"
	rule5010835 "github.com/moov-io/x12/rules/rule_5010_835"
	rule5010837p "github.com/moov-io/x12/rules/rule_5010_837p"
	rulestp820 "github.com/moov-io/x12/rules/rule_stp_820"
)

func FuzzReaderWriterX12(f *testing.F) {
	populateCorpus(f)

	ruleSet := []*rules.InterchangeRule{
		&rule5010835.InterchangeRule,
		&rule5010837p.InterchangeRule,
		&rule4010810.InterchangeRule,
		&rulestp820.InterchangeRule,
	}

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}

		for _, rule := range ruleSet {
			fl := file.NewFile(rule)
			_ = fl.Parse(file.NewScanner(strings.NewReader(contents)))
			_ = fl.Validate()
			_ = fl.String()
		}
	})
}

func populateCorpus(f *testing.F) {
	f.Helper()

	f.Add("")
	f.Add("ISA*")
	f.Add("ISA*00*          *00*          *ZZ*SENDER         *ZZ*RECEIVER       *250101*1200*^*00501*000000001*0*P*:~")

	_ = filepath.Walk(filepath.Join("..", "..", "examples"), func(path string, info fs.FileInfo, err error) error {
		if err != nil || info == nil || info.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(path), ".txt") {
			bs, err := os.ReadFile(path)
			if err != nil || len(bs) > 512*1024 {
				return nil
			}
			f.Add(string(bs))
		}
		return nil
	})
}
