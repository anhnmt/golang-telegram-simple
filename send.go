package telegram

import "fmt"

func OK() *Telegram {
	return Default().OK()
}

func (t *Telegram) OK() *Telegram {
	return t.SetStatus(StatusOK)
}

func Err(err error) *Telegram {
	return Default().Err(err)
}

func (t *Telegram) Err(err error) *Telegram {
	t.err = err
	return t.SetStatus(StatusErr)
}

func Msg(msg string, a ...any) error {
	return Default().Msg(msg, a...)
}

func (t *Telegram) Msg(msg string, a ...any) error {
	return t.action(SendMessage, fmt.Sprintf(msg, a...))
}
