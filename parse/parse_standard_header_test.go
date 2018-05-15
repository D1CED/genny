package parse

import (
	"regexp"
	"testing"
)

// "^// Code generated .* DO NOT EDIT.$" // official
// "^// Code generated .*DO NOT EDIT.?$" // modified
// (?m) enables multiline mode. needed for ^ and $
// see regexp/syntax documentation
const headerMatch = "(?m)^// Code generated .*DO NOT EDIT.?$"

// See github.com/golang/go/issues/13560
// auto-generated files must contain matching expression before package declaration
func TestStandardConformingHeader(t *testing.T) {
	regex := regexp.MustCompile(headerMatch)
	if regex.Find(header) == nil {
		t.Errorf("Header %s does not match required regex.\n", header)
	}
}
