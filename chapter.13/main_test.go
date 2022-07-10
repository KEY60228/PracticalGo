package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestTom(t *testing.T) {
	tom := getTom()
	tom2 := getTom2()
	opt := []cmp.Option{
		cmp.AllowUnexported(User{}),
		cmpopts.IgnoreFields(User{}, "id"),
	}
	if diff := cmp.Diff(tom, tom2, opt...); diff != "" {
		t.Errorf("User value is mismatch (-tom +tom2):\n%s", diff)
	}
}
