# lintcomparestrings

A Go linter which checks whether the more efficient
[`strings.Compare`](https://pkg.go.dev/strings#Compare) three-way-compare
function is used for strings rather than [`cmp.Compare`](https://pkg.go.dev/cmp#Compare).

Using cmp.Compare to compare strings is less performant than strings.Compare
since Go 1.23, especially for large strings. See https://go.dev/issues/61725 and
https://go.dev/cl/532195 for details.`

## Installation

    go install github.com/tklauser/lintcomparestrings@latest
