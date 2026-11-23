## Unreleased

IMPROVEMENTS

- cmd/x12: support every built-in rule (`4010_810`, `4010_820`, `4010_835`, `4010_837i`, `5010_835`, `5010_837d`, `5010_837p`, `stp_820`)
- cmd/x12: `describe_file` detects the rule from GS08 when `-rule` is not given
- cmd/x12: `describe_file` exits with a non-zero code when a file cannot be read, parsed, or validated
- cmd/x12: `describe_rule` runs without arguments, and the usage text lists `-version`
- cmd/x12: `-version` shows the release version for `go install` and `make dist` builds

BUG FIXES

- rules: correct the names of the `rule_5010_835` and `rule_4010_837I` rules

## v0.1.1 (Released 2026-09-25)

ADDITIONS

- Add rules for 4010 810, 4010 820, 4010 835, 4010 837I, 5010 835 and STP 820, with examples
- Add a control number validator

IMPROVEMENTS

- Transaction sets use a composite loop (`TransactionRule.Composite`) in place of the `Segments` and `Loops` fields
- Errors include the segment line and the stack of loops where the error occurred
- Reject non-numeric X12 trailer counts ([#45](https://github.com/moov-io/x12/pull/45))

BUG FIXES

- rule_5010_837p: fix the rule name and other fixes

BUILD

- Add fuzz tests seeded from the examples
- Update dependencies

## v0.1.0 (Released 2023-04-05)

INITIAL RELEASE

The v0.1.0 release is the first milestone for the x12 project. We have implemented a basic reader/writer for Files. Also included is support for Health Care Claim/Encounter: Dental (837D) and Health Care Claim Professional (837P) records. We encourage people to try out the library and any report features, bugs, issues, or other comments.
