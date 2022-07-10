package main

import (
	"context"
	"fmt"
	"time"
)

type contextTimeKey string

const timeKey contextTimeKey = "timeKey"

func CurrentTime(ctx context.Context) time.Time {
	v := ctx.Value(timeKey)
	if t, ok := v.(time.Time); ok {
		return t
	}
	return time.Now()
}

func SetFixTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, timeKey, t)
}

func NextMonth(ctx context.Context) time.Month {
	now := CurrentTime(ctx)
	return now.AddDate(0, 1, 0).Month()
}

func main() {
	fmt.Println(NextMonth(context.Background()))
}
