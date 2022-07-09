package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Debug().Msgf("debug") // 出力されない
	log.Info().Msg("info")
	log.Error().Msg("error")
}
