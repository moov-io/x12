// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rule837d "github.com/moov-io/x12/rules/rule_5010_837d"
	rule837p "github.com/moov-io/x12/rules/rule_5010_837p"

	"github.com/moov-io/x12"
	"github.com/moov-io/x12/pkg/file"
	"github.com/moov-io/x12/pkg/rules"
)

var (
	programName     = filepath.Base(os.Args[0])
	describeFileCmd = "describe_file"
	describeRuleCmd = "describe_rule"
)

var availableRules = map[string]rules.InterchangeRule{
	"5010_837d": rule837d.InterchangeRule,
	"5010_837p": rule837p.InterchangeRule,
}

func main() {
	versionFlag := flag.Bool("version", false, "show version")
	describeFileCommand := flag.NewFlagSet(describeFileCmd, flag.ExitOnError)
	describeRuleCommand := flag.NewFlagSet(describeRuleCmd, flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Work seamlessly with EDI from the command line.\n\nUsage:\n  %s <command> [flags]\n\n", programName)
		fmt.Fprintf(os.Stdout, "Available commands:\n")

		fmt.Fprintf(os.Stdout, "  %s: display EDI file in a human-readable format\n", describeFileCmd)
		fmt.Fprintf(os.Stdout, "  %s: display Rule in a human-readable format\n", describeRuleCmd)
		fmt.Fprintf(os.Stdout, "\n")
	}

	describeFileCommand.Usage = func() {
		fmt.Fprintf(os.Stdout, "Display EDI in a human-readable format.\n\nUsage:\n  %s %s [flags] <files> \n\n", programName, describeFileCmd)
		fmt.Fprintf(os.Stdout, "Flags: \n")
		describeFileCommand.PrintDefaults()
		fmt.Fprintf(os.Stdout, "\n")
	}

	describeRuleCommand.Usage = func() {
		fmt.Fprintf(os.Stdout, "Display rule in a human-readable format.\n\nUsage:\n  %s %s [flags] <files> \n\n", programName, describeRuleCmd)
		fmt.Fprintf(os.Stdout, "Flags: \n")
		describeRuleCommand.PrintDefaults()
		fmt.Fprintf(os.Stdout, "\n")
	}

	var ruleNames []string
	for name := range availableRules {
		ruleNames = append(ruleNames, name)
	}
	availableRuleNames := strings.Join(ruleNames, ", ")

	flag.Parse()

	if *versionFlag {
		fmt.Fprintf(os.Stdout, "Version: %s\n\n", x12.Version)
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	command := flag.Arg(0)

	switch command {
	case describeFileCmd:

		ruleName := describeFileCommand.String("rule", "5010_837d", fmt.Sprintf("name of built-in rule: %s", availableRuleNames))
		validateFlag := describeFileCommand.Bool("validate", false, "running validation check")

		describeArgs := os.Args[2:]
		if len(describeArgs) == 0 {
			describeFileCommand.Usage()
			os.Exit(1)
		}

		describeFileCommand.Parse(os.Args[2:])

		var err error
		if rule, ok := availableRules[*ruleName]; ok {
			err = describeFile(describeFileCommand.Args(), rule, *validateFlag)
		} else {
			fmt.Fprintf(os.Stdout, "Unknown rule: %s\n\n", *ruleName)
			fmt.Fprintf(os.Stdout, "Supported rules: %s\n\n", availableRuleNames)
			os.Exit(1)
		}

		if err != nil {
			fmt.Fprintf(os.Stdout, "Error describing files: %s\n", err)
			os.Exit(1)
		}

	case describeRuleCmd:

		ruleName := describeRuleCommand.String("rule", "5010_837d", fmt.Sprintf("name of built-in rule: %s", availableRuleNames))
		isRequiredOnly := describeRuleCommand.Bool("required", false, "display required segments & rule only")

		describeArgs := os.Args[2:]
		if len(describeArgs) == 0 {
			describeRuleCommand.Usage()
			os.Exit(1)
		}

		describeRuleCommand.Parse(os.Args[2:])

		var err error
		if rule, ok := availableRules[*ruleName]; ok {
			err = describeRule(rule, *isRequiredOnly)
		} else {
			fmt.Fprintf(os.Stdout, "Unknown rule: %s\n\n", *ruleName)
			fmt.Fprintf(os.Stdout, "Supported rules: %s\n\n", availableRuleNames)
			os.Exit(1)
		}

		if err != nil {
			fmt.Fprintf(os.Stdout, "Error describing files: %s\n", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintf(os.Stdout, "Uknown command: %s\n\n", command)
		flag.Usage()
		os.Exit(1)
	}

	if describeFileCommand.Parsed() {
		files := describeFileCommand.Args()
		if len(files) == 0 {
			describeFileCommand.Usage()
			os.Exit(1)
		}
	}
}

func describeFile(paths []string, rule rules.InterchangeRule, validateFlag bool) error {

	for _, path := range paths {
		reader, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Failed to read edi from file(%s): %v\n", path, err)
			continue
		}

		f := file.NewFile(&rule)

		if err = f.Parse(file.NewScanner(reader)); err != nil {
			fmt.Fprintf(os.Stdout, "Failed to parse edi from file(%s): %v\n", path, err)
			continue
		}

		if validateFlag {
			if f.Validate() == nil {
				fmt.Fprintf(os.Stdout, "The edi file is valid(%s)\n", path)
			} else {
				fmt.Fprintf(os.Stdout, "Failed to validate edi(%s): %v\n", path, err)
			}
		}

		f.Print(os.Stdout)

		reader.Close()
	}

	return nil
}

func describeRule(rule rules.InterchangeRule, isRequiredOnly bool) error {

	rule.Print(os.Stdout, isRequiredOnly)

	return nil
}
