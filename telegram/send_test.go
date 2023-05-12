package telegram

import (
	"testing"
)

func TestSendMsg(t *testing.T) {
	Default.
		SetChatId("-9003481").
		SetToken("6108764305:AAGw2BVSPYPjcc8l940bswQUTRUZIssS")

	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test send",
			args: args{
				msg: "test message",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Msg(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Msg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTelegram_SendMsg(t1 *testing.T) {
	tele := DefaultTelegram().
		SetChatId("-9003481").
		SetToken("6108764305:AAGw2BVSPYPjcc8l940bswQUTRUZIssS")

	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		fields  *Telegram
		args    args
		wantErr bool
	}{
		{
			name:   "test",
			fields: tele,
			args: args{
				msg: "test message",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if err := tt.fields.Msg(tt.args.msg); (err != nil) != tt.wantErr {
				t1.Errorf("Msg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
