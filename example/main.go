package main

import (
	"fmt"

	telegram "github.com/anhnmt/golang-telegram-simple"
)

func main() {
	fmt.Println("Hello, world!")

	telegram.
		SetEnabled(true).
		SetEnv("DEV").
		SetToken("abc:123456789").
		SetChatId(-123456789).
		OK("Hello, world!")
}
