package telegram

import (
	"fmt"
	"strings"
	"sync/atomic"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

type Telegram struct {
	bot     *tgbotapi.BotAPI
	enabled bool
	chatId  int64
	env     string
	token   string
}

var defaultTelegram atomic.Value

func Default() *Telegram {
	return defaultTelegram.Load().(*Telegram)
}

func SetDefault(t *Telegram) {
	defaultTelegram.Store(t)
}

func New() (*Telegram, error) {
	t := &Telegram{
		env:     viper.GetString(Env),
		token:   viper.GetString(Token),
		enabled: viper.GetBool(Enabled),
		chatId:  viper.GetInt64(ChatId),
	}

	bot, err := tgbotapi.NewBotAPI(t.token)
	if err != nil {
		return nil, err
	}

	t.bot = bot

	return t, nil
}

func SetEnabled(enabled bool) *Telegram {
	return Default().SetEnabled(enabled)
}

func (t *Telegram) SetEnabled(enabled bool) *Telegram {
	t.enabled = enabled
	return t
}

func SetEnv(env string) *Telegram {
	return Default().SetEnv(env)
}

func (t *Telegram) SetEnv(env string) *Telegram {
	t.env = env
	return t
}

func SetToken(token string) *Telegram {
	return Default().SetToken(token)
}

func (t *Telegram) SetToken(token string) *Telegram {
	t.token = token

	bot, err := tgbotapi.NewBotAPI(t.token)
	if err == nil {
		t.bot = bot
	}

	return t
}

func SetChatId(chatId int64) *Telegram {
	return Default().SetChatId(chatId)
}

func (t *Telegram) SetChatId(chatId int64) *Telegram {
	t.chatId = chatId
	return t
}

func OK(text string, a ...any) error {
	return Default().OK(text, a...)
}

func (t *Telegram) OK(text string, a ...any) error {
	return t.msg(nil, text, a...)
}

func Err(err error, text string, a ...any) error {
	return Default().Err(err, text, a...)
}

func (t *Telegram) Err(err error, text string, a ...any) error {
	return t.msg(err, text, a...)
}

func (t *Telegram) msg(err error, text string, a ...any) error {
	if !t.enabled || t.chatId == 0 {
		return nil
	}

	if len(a) > 0 {
		text = fmt.Sprintf(text, a...)
	}

	if err != nil {
		text = fmt.Sprintf("\n- Error: %v\n%s", err, text)
	}

	if t.env != "" {
		text = fmt.Sprintf("[%s] - %s", strings.ToUpper(t.env), text)
	}

	if err != nil {
		text = fmt.Sprintf("🔴 %s", text)
	} else {
		text = fmt.Sprintf("🟢 %s", text)
	}

	msg := tgbotapi.NewMessage(t.chatId, text)
	_, sendErr := t.bot.Send(msg)
	if sendErr != nil {
		return sendErr
	}

	return nil
}
