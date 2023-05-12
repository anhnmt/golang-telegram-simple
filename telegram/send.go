package telegram

func OK() *Telegram {
	return Default.SetStatus(StatusOK)
}

func (t *Telegram) OK() *Telegram {
	return t.SetStatus(StatusOK)
}

func Err() *Telegram {
	return Default.SetStatus(StatusErr)
}

func (t *Telegram) Err() *Telegram {
	return t.SetStatus(StatusErr)
}

func Msg(msg string) error {
	return Default.Msg(msg)
}

func (t *Telegram) Msg(msg string) error {
	return t.action(SendMessage, msg)
}
