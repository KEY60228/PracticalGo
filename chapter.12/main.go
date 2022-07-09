package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	logger := log.With().Int("user_id", 1024).Str("path", "/api/user").Str("method", "post").Logger()
	ctx := logger.WithContext(context.Background())

	newLogger := zerolog.Ctx(ctx)
	newLogger.Print("debug message")
}
