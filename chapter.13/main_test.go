package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTom(t *testing.T) {
	tom := getTom()
	tom2 := getTom2()
	if diff := cmp.Diff(tom, tom2); diff != "" {
		t.Errorf("User value is mismatch (-tom +tom2):\n%s", diff)
	}
}
