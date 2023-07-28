package main

import (
	"github.com/rs/zerolog/log"

	telegram "github.com/anhnmt/golang-telegram-simple"
)

func main() {
	log.Info().Msg("Hello, world!")

	telegram.OK().
		SetEnabled(true).
		SetEnv("DEV").
		SetToken("123").
		SetChatId("-123456789").
		Msg("Hello, world!")
}
