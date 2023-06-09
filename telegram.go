package telegram

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	SendMessage = "sendMessage"
)

var (
	json = jsoniter.ConfigFastest
)

type status int

const (
	StatusOK status = iota
	StatusErr
)

type Telegram struct {
	enabled bool
	apiUrl  string
	token   string
	chatId  string
	env     string
	status  status
	err     error
}

var Default *Telegram

func init() {
	Default = DefaultTelegram()
}

func DefaultTelegram() *Telegram {
	t := &Telegram{
		enabled: viper.GetBool("telegram.enabled"),
		apiUrl:  "https://api.telegram.org",
		env:     os.Getenv("env"),
		token:   viper.GetString("telegram.token"),
		chatId:  viper.GetString("telegram.chatId"),
	}

	return t
}

func NewTelegram(env string) {
	Default = DefaultTelegram().
		SetEnv(env)

	log.Info().
		Bool("enabled", Default.enabled).
		Str("chatId", Default.chatId).
		Msg("Init to Telegram")
}

func SetStatus(status status) *Telegram {
	return Default.SetStatus(status)
}

func (t *Telegram) SetStatus(status status) *Telegram {
	t.status = status
	return t
}

func SetToken(token string) *Telegram {
	return Default.SetToken(token)
}

func (t *Telegram) SetToken(token string) *Telegram {
	t.token = token
	return t
}

func SetEnv(env string) *Telegram {
	return Default.SetEnv(env)
}

func (t *Telegram) SetEnv(env string) *Telegram {
	t.env = env
	return t
}

func SetChatId(chatId string) *Telegram {
	return Default.SetChatId(chatId)
}

func (t *Telegram) SetChatId(chatId string) *Telegram {
	t.chatId = chatId
	return t
}

func (t *Telegram) action(method, msg string) error {
	if !t.enabled {
		return nil
	}

	if t.err != nil {
		msg = fmt.Sprintf("\n- Error: %v\n%s", t.err, msg)
		defer func(t *Telegram) {
			t.err = nil
		}(t)
	}

	if t.env != "" {
		msg = fmt.Sprintf("[%s] - %s", strings.ToUpper(t.env), msg)
	}

	switch t.status {
	case StatusOK:
		msg = fmt.Sprintf("🟢 %s", msg)
	case StatusErr:
		msg = fmt.Sprintf("🔴 %s", msg)
	}

	data := map[string]interface{}{
		"chat_id": t.chatId,
		"text":    msg,
	}

	// Chuyển requestBody thành JSON
	jsonValue, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Tạo request
	url := fmt.Sprintf("%s/bot%s/%s", t.apiUrl, t.token, method)

	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Err(err).Msg("Error create http.NewRequest")
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	// Gửi request
	client := http.DefaultClient
	client.Timeout = 10 * time.Second

	response, err := client.Do(request)
	if err != nil {
		log.Err(err).Msg("Error sending message")
		return err
	}
	defer response.Body.Close()

	return nil
}
