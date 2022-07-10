package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTom(t *testing.T) {
	tom := getTom()
	tom2 := getTom2()
	assert.Equal(t, tom, tom2)
}
