package memviz

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  string
	output string
}{
	{"Hello world", "Hello world"},

	// double quotes are escaped
	{"\"Hello world\"", "\\\"Hello world\\\""},

	// brackets not escaped
	{"map[string]bool", "map[string]bool"},

	// braces escaped
	{"map[string]struct{}", "map[string]struct\\{\\}"},
}

func TestEscapeString(t *testing.T) {
	for _, tc := range cases {
		assert.Equal(t, tc.output, escapeString(tc.input))
	}
}

func TestEmptyStruct(t *testing.T) {
	set := make([]struct{}, 2)
	set[0] = struct{}{}
	set[1] = struct{}{}

	b := &bytes.Buffer{}
	Map(b, &set)
	fmt.Println(b.String())
	cupaloy.SnapshotT(t, b.Bytes())
}

func TestEmptyInterface(t *testing.T) {
	set := map[string]interface{}{}
	set["hello world"] = nil

	b := &bytes.Buffer{}
	Map(b, &set)
	fmt.Println(b.String())
	cupaloy.SnapshotT(t, b.Bytes())
}
