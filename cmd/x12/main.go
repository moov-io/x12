// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"sort"
	"strings"

	rule810 "github.com/moov-io/x12/rules/rule_4010_810"
	rule4010_820 "github.com/moov-io/x12/rules/rule_4010_820"
	rule4010_835 "github.com/moov-io/x12/rules/rule_4010_835"
	rule837i "github.com/moov-io/x12/rules/rule_4010_837I"
	rule5010_835 "github.com/moov-io/x12/rules/rule_5010_835"
	rule837d "github.com/moov-io/x12/rules/rule_5010_837d"
	rule837p "github.com/moov-io/x12/rules/rule_5010_837p"
	ruleStp820 "github.com/moov-io/x12/rules/rule_stp_820"

	"github.com/moov-io/x12"
	"github.com/moov-io/x12/pkg/file"
	"github.com/moov-io/x12/pkg/rules"
)

var (
	describeFileCmd = "describe_file"
	describeRuleCmd = "describe_rule"

	defaultRuleName = "5010_837d"
)

var availableRules = map[string]rules.InterchangeRule{
	"4010_810":  rule810.InterchangeRule,
	"4010_820":  rule4010_820.InterchangeRule,
	"4010_835":  rule4010_835.InterchangeRule,
	"4010_837i": rule837i.InterchangeRule,
	"5010_835":  rule5010_835.InterchangeRule,
	"5010_837d": rule837d.InterchangeRule,
	"5010_837p": rule837p.InterchangeRule,
	"stp_820":   ruleStp820.InterchangeRule,
}

func main() {
	os.Exit(run(filepath.Base(os.Args[0]), os.Args[1:], os.Stdout, os.Stderr))
}

func ruleNames() []string {
	names := make([]string, 0, len(availableRules))
	for name := range availableRules {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func findRule(name string) (rules.InterchangeRule, bool) {
	rule, ok := availableRules[strings.ToLower(name)]
	return rule, ok
}

// version returns the version set at build time. When it is not set (for example
// after "go install github.com/moov-io/x12/cmd/x12@latest") it uses the module version.
func version() string {
	if x12.Version != "" && x12.Version != "v0.0.0" {
		return x12.Version
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if v := info.Main.Version; v != "" && v != "(devel)" {
			return v
		}
	}
	return x12.Version
}

// run executes the CLI and returns the exit code.
func run(programName string, args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet(programName, flag.ContinueOnError)
	fs.SetOutput(stderr)
	versionFlag := fs.Bool("version", false, "show version")

	availableRuleNames := strings.Join(ruleNames(), ", ")

	fs.Usage = func() {
		fmt.Fprintf(stdout, "Work seamlessly with EDI from the command line.\n\nUsage:\n  %s <command> [flags]\n  %s -version\n\n", programName, programName)
		fmt.Fprintf(stdout, "Available commands:\n")
		fmt.Fprintf(stdout, "  %s: display EDI files in a human-readable format\n", describeFileCmd)
		fmt.Fprintf(stdout, "  %s: display a rule in a human-readable format\n", describeRuleCmd)
		fmt.Fprintf(stdout, "\nFlags:\n")
		fmt.Fprintf(stdout, "  -version\n    \tshow version\n")
		fmt.Fprintf(stdout, "\nSupported rules: %s\n\n", availableRuleNames)
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}

	if *versionFlag {
		fmt.Fprintf(stdout, "Version: %s\n\n", version())
		return 0
	}

	if fs.NArg() < 1 {
		fs.Usage()
		return 1
	}

	command := fs.Arg(0)
	commandArgs := fs.Args()[1:]

	switch command {
	case describeFileCmd:
		cmd := flag.NewFlagSet(describeFileCmd, flag.ContinueOnError)
		cmd.SetOutput(stderr)
		ruleName := cmd.String("rule", "", fmt.Sprintf("name of built-in rule: %s (default: detect from GS08)", availableRuleNames))
		validateFlag := cmd.Bool("validate", false, "validate each file")
		cmd.Usage = func() {
			fmt.Fprintf(stdout, "Display EDI in a human-readable format.\n\nUsage:\n  %s %s [flags] <files>\n\n", programName, describeFileCmd)
			fmt.Fprintf(stdout, "Flags:\n")
			cmd.SetOutput(stdout)
			cmd.PrintDefaults()
			cmd.SetOutput(stderr)
			fmt.Fprintf(stdout, "\n")
		}
		if err := cmd.Parse(commandArgs); err != nil {
			if errors.Is(err, flag.ErrHelp) {
				return 0
			}
			return 2
		}
		if cmd.NArg() == 0 {
			cmd.Usage()
			return 1
		}

		var rule *rules.InterchangeRule
		if *ruleName != "" {
			r, ok := findRule(*ruleName)
			if !ok {
				fmt.Fprintf(stderr, "Unknown rule: %s\n\n", *ruleName)
				fmt.Fprintf(stderr, "Supported rules: %s\n\n", availableRuleNames)
				return 1
			}
			rule = &r
		}

		if err := describeFiles(stdout, stderr, cmd.Args(), rule, *validateFlag); err != nil {
			fmt.Fprintf(stderr, "Error describing files: %v\n", err)
			return 1
		}

	case describeRuleCmd:
		cmd := flag.NewFlagSet(describeRuleCmd, flag.ContinueOnError)
		cmd.SetOutput(stderr)
		ruleName := cmd.String("rule", defaultRuleName, fmt.Sprintf("name of built-in rule: %s", availableRuleNames))
		isRequiredOnly := cmd.Bool("required", false, "display required segments & rule only")
		cmd.Usage = func() {
			fmt.Fprintf(stdout, "Display a rule in a human-readable format.\n\nUsage:\n  %s %s [flags]\n\n", programName, describeRuleCmd)
			fmt.Fprintf(stdout, "Flags:\n")
			cmd.SetOutput(stdout)
			cmd.PrintDefaults()
			cmd.SetOutput(stderr)
			fmt.Fprintf(stdout, "\n")
		}
		if err := cmd.Parse(commandArgs); err != nil {
			if errors.Is(err, flag.ErrHelp) {
				return 0
			}
			return 2
		}

		rule, ok := findRule(*ruleName)
		if !ok {
			fmt.Fprintf(stderr, "Unknown rule: %s\n\n", *ruleName)
			fmt.Fprintf(stderr, "Supported rules: %s\n\n", availableRuleNames)
			return 1
		}
		rule.Print(stdout, *isRequiredOnly)

	default:
		fmt.Fprintf(stderr, "Unknown command: %s\n\n", command)
		fs.Usage()
		return 1
	}

	return 0
}

// describeFiles prints each file. It returns an error when one or more files
// cannot be read, parsed, or validated.
func describeFiles(stdout, stderr io.Writer, paths []string, rule *rules.InterchangeRule, validate bool) error {
	var failed []string

	for _, path := range paths {
		if err := describeFile(stdout, path, rule, validate); err != nil {
			fmt.Fprintf(stderr, "%v\n", err)
			failed = append(failed, path)
		}
	}

	if len(failed) > 0 {
		return fmt.Errorf("%d of %d file(s) failed: %s", len(failed), len(paths), strings.Join(failed, ", "))
	}
	return nil
}

func describeFile(w io.Writer, path string, rule *rules.InterchangeRule, validate bool) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Failed to read edi from file(%s): %w", path, err)
	}

	if rule == nil {
		name, r, err := detectRule(string(data))
		if err != nil {
			return fmt.Errorf("Failed to detect rule for file(%s): %w (use -rule)", path, err)
		}
		fmt.Fprintf(w, "Using rule %s for file(%s)\n", name, path)
		rule = &r
	}

	f := file.NewFile(rule)
	if err := f.Parse(file.NewScanner(strings.NewReader(string(data)))); err != nil {
		return fmt.Errorf("Failed to parse edi from file(%s): %w", path, err)
	}

	if validate {
		if err := f.Validate(); err != nil {
			return fmt.Errorf("Failed to validate edi(%s): %w", path, err)
		}
		fmt.Fprintf(w, "The edi file is valid(%s)\n", path)
	}

	f.Print(w)

	return nil
}

// readGS08 returns the version code (GS08) of the first functional group.
func readGS08(data string) (string, error) {
	data = strings.NewReplacer("\r", "", "\n", "").Replace(data)

	start := strings.Index(data, "ISA")
	if start < 0 {
		return "", errors.New("ISA segment not found")
	}
	isa := data[start:]
	// The ISA segment has a fixed length of 106 characters, including the segment terminator.
	if len(isa) < 106 {
		return "", errors.New("ISA segment is too short")
	}
	separator := string(isa[3])
	terminator := string(isa[105])

	for _, segment := range strings.Split(isa, terminator) {
		segment = strings.TrimSpace(segment)
		if !strings.HasPrefix(segment, "GS"+separator) {
			continue
		}
		elements := strings.Split(segment, separator)
		if len(elements) < 9 || elements[8] == "" {
			return "", errors.New("GS08 not found")
		}
		return elements[8], nil
	}
	return "", errors.New("GS segment not found")
}

func contains(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

// detectRule finds the built-in rule which accepts the GS08 value of the file.
// A rule that lists the exact value is used before a rule that matches it with a regex.
func detectRule(data string) (string, rules.InterchangeRule, error) {
	gs08, err := readGS08(data)
	if err != nil {
		return "", rules.InterchangeRule{}, err
	}

	var exact, regex []string
	for _, name := range ruleNames() {
		elm, ok := availableRules[name].Group.GS.Elements["08"]
		if !ok {
			continue
		}
		if contains(elm.AcceptValues, gs08) {
			exact = append(exact, name)
			continue
		}
		if elm.AcceptRegex != "" {
			if re, err := regexp.Compile(elm.AcceptRegex); err == nil && re.MatchString(gs08) {
				regex = append(regex, name)
			}
		}
	}

	matches := exact
	if len(matches) == 0 {
		matches = regex
	}
	switch len(matches) {
	case 0:
		return "", rules.InterchangeRule{}, fmt.Errorf("no built-in rule accepts GS08 value %q", gs08)
	case 1:
		return matches[0], availableRules[matches[0]], nil
	default:
		return "", rules.InterchangeRule{}, fmt.Errorf("more than one built-in rule accepts GS08 value %q: %s", gs08, strings.Join(matches, ", "))
	}
}
