# golang-telegram-simple

```go
package main

func main() {
	log.Info().Msg("Hello, world!")

	telegram.OK().
		SetEnv("DEV").
		SetToken("123").
		SetChatId("-123456789").
		Msg("Hello, world!")
}
```