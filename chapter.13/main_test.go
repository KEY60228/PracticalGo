package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNextMonth(t *testing.T) {
	ctx := SetFixTime(context.Background(), time.Date(1980, time.December, 1, 0, 0, 0, 0, time.Local))
	assert.Equal(t, time.January, NextMonth(ctx))
}
