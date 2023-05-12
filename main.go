package main

import (
	"github.com/anhnmt/golang-telegram-simple/telegram"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Hello, world!")

	telegram.OK().
		SetEnv("DEV").
		SetToken("123").
		SetChatId("-123456789").
		Msg("Hello, world!")
}
