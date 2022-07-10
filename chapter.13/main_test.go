package main

import (
	"bytes"
	"testing"
)

func TestConsoleOut(t *testing.T) {
	var b bytes.Buffer
	DumpUserTo(&b, &User{Name: "KEY"})
	if b.String() != "KEY(住所不定)" {
		t.Errorf("error (expected: 'KEY(住所不定)', actual = '%s'", b.String())
	}
}
