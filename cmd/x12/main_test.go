// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func sample(parts ...string) string {
	return filepath.Join(append([]string{"..", "..", "examples"}, parts...)...)
}

func runCLI(t *testing.T, args ...string) (int, string, string) {
	t.Helper()

	var stdout, stderr bytes.Buffer
	code := run("x12", args, &stdout, &stderr)
	return code, stdout.String(), stderr.String()
}

func TestRun_Usage(t *testing.T) {
	code, stdout, _ := runCLI(t)
	require.Equal(t, 1, code)
	require.Contains(t, stdout, "describe_file")
	require.Contains(t, stdout, "-version")
	require.Contains(t, stdout, "5010_835")

	code, stdout, stderr := runCLI(t, "bogus")
	require.Equal(t, 1, code)
	require.Contains(t, stderr, "Unknown command: bogus")
	require.Contains(t, stdout, "Available commands")

	code, _, _ = runCLI(t, "-h")
	require.Equal(t, 0, code)

	code, _, _ = runCLI(t, "-nope")
	require.Equal(t, 2, code)
}

func TestRun_Version(t *testing.T) {
	code, stdout, _ := runCLI(t, "-version")
	require.Equal(t, 0, code)
	require.Contains(t, stdout, "Version: ")
	require.NotEmpty(t, version())
}

func TestRun_DescribeRule(t *testing.T) {
	// No arguments uses the default rule
	code, stdout, _ := runCLI(t, "describe_rule")
	require.Equal(t, 0, code)
	require.Contains(t, stdout, "DUMP RULE 837D(005010X224A2)")

	code, stdout, _ = runCLI(t, "describe_rule", "-required", "-rule", "5010_835")
	require.Equal(t, 0, code)
	require.Contains(t, stdout, "DUMP RULE 835(005010X221A1)")

	// Rule names are not case sensitive
	code, stdout, _ = runCLI(t, "describe_rule", "-rule", "4010_837I")
	require.Equal(t, 0, code)
	require.Contains(t, stdout, "DUMP RULE 837I(004010X096A1)")

	code, _, stderr := runCLI(t, "describe_rule", "-rule", "nope")
	require.Equal(t, 1, code)
	require.Contains(t, stderr, "Unknown rule: nope")

	code, stdout, _ = runCLI(t, "describe_rule", "-h")
	require.Equal(t, 0, code)
	require.Contains(t, stdout, "describe_rule [flags]")
	require.NotContains(t, stdout, "<files>")

	code, _, _ = runCLI(t, "describe_rule", "-nope")
	require.Equal(t, 2, code)
}

func TestRun_DescribeFile(t *testing.T) {
	t.Run("usage", func(t *testing.T) {
		code, stdout, _ := runCLI(t, "describe_file")
		require.Equal(t, 1, code)
		require.Contains(t, stdout, "describe_file [flags] <files>")

		code, _, _ = runCLI(t, "describe_file", "-h")
		require.Equal(t, 0, code)

		code, _, _ = runCLI(t, "describe_file", "-nope")
		require.Equal(t, 2, code)
	})

	t.Run("detect rule", func(t *testing.T) {
		cases := map[string]string{
			"5010_835":  sample("example_5010_835", "sample2.txt"),
			"5010_837d": sample("example_5010_837d", "sample.txt"),
			"5010_837p": sample("example_5010_837p", "sample.txt"),
		}
		for ruleName, path := range cases {
			code, stdout, stderr := runCLI(t, "describe_file", "-validate", path)
			require.Equal(t, 0, code, stderr)
			require.Contains(t, stdout, "Using rule "+ruleName)
			require.Contains(t, stdout, "The edi file is valid")
		}
	})

	t.Run("explicit rule", func(t *testing.T) {
		code, stdout, stderr := runCLI(t, "describe_file", "-rule", "5010_835", "-validate", sample("example_5010_835", "sample1.txt"))
		require.Equal(t, 0, code, stderr)
		require.NotContains(t, stdout, "Using rule")
		require.Contains(t, stdout, "DUMP EDI FILE WITH 835(005010X221A1)")
	})

	t.Run("unknown rule", func(t *testing.T) {
		code, _, stderr := runCLI(t, "describe_file", "-rule", "nope", sample("example_5010_835", "sample1.txt"))
		require.Equal(t, 1, code)
		require.Contains(t, stderr, "Unknown rule: nope")
	})

	t.Run("missing file", func(t *testing.T) {
		code, _, stderr := runCLI(t, "describe_file", "nonexist.txt")
		require.Equal(t, 1, code)
		require.Contains(t, stderr, "Failed to read edi from file(nonexist.txt)")
	})

	t.Run("parse failure", func(t *testing.T) {
		code, _, stderr := runCLI(t, "describe_file", "-rule", "5010_837d", sample("example_5010_837p", "sample.txt"))
		require.Equal(t, 1, code)
		require.Contains(t, stderr, "Failed to parse edi")
	})

	t.Run("validate failure", func(t *testing.T) {
		bs, err := os.ReadFile(sample("example_5010_837p", "sample.txt"))
		require.NoError(t, err)

		// Remove the CLM segment, which the claim loop requires.
		var segments []string
		for _, seg := range strings.Split(string(bs), "~") {
			if !strings.HasPrefix(strings.TrimSpace(seg), "CLM*") {
				segments = append(segments, seg)
			}
		}
		path := filepath.Join(t.TempDir(), "invalid.txt")
		require.NoError(t, os.WriteFile(path, []byte(strings.Join(segments, "~")), 0600))

		code, _, stderr := runCLI(t, "describe_file", "-validate", path)
		require.Equal(t, 1, code)
		require.Contains(t, stderr, "Failed to")
	})

	t.Run("one bad file of many", func(t *testing.T) {
		code, stdout, stderr := runCLI(t, "describe_file", sample("example_5010_837d", "sample.txt"), "nonexist.txt")
		require.Equal(t, 1, code)
		require.Contains(t, stdout, "Using rule 5010_837d")
		require.Contains(t, stderr, "1 of 2 file(s) failed: nonexist.txt")
	})
}

func TestDetectRule(t *testing.T) {
	_, _, err := detectRule("not edi")
	require.ErrorContains(t, err, "ISA segment not found")

	_, _, err = detectRule("ISA*00*")
	require.ErrorContains(t, err, "too short")

	isa := "ISA*00*          *00*          *ZZ*SENDER         *ZZ*RECEIVER       *120419*2125*^*00501*000002120*0*P*:~"
	require.Len(t, isa, 106)

	_, _, err = detectRule(isa + "ST*837*0001~")
	require.ErrorContains(t, err, "GS segment not found")

	_, _, err = detectRule(isa + "GS*HC*A*B~")
	require.ErrorContains(t, err, "GS08 not found")

	_, _, err = detectRule(isa + "GS*HC*A*B*20120419*2125*1*X*999999~")
	require.ErrorContains(t, err, `no built-in rule accepts GS08 value "999999"`)

	name, _, err := detectRule(isa + "GS*HC*A*B*20120419*2125*1*X*005010X222A1~")
	require.NoError(t, err)
	require.Equal(t, "5010_837p", name)

	name, _, err = detectRule(isa + "\nGS*RA*A*B*20120419*2125*1*X*004010STP820~")
	require.NoError(t, err)
	require.Equal(t, "stp_820", name)

	// Only the 810 rule accepts GS08 with a regex
	name, _, err = detectRule(isa + "GS*IN*A*B*20120419*2125*1*X*004010~")
	require.NoError(t, err)
	require.Equal(t, "4010_810", name)
}
