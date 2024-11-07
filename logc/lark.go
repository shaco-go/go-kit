package logc

import (
	"github.com/shaco-go/go-kit/notify/feishu/client"
)

func NewLarkWriter(conf Lark) *LarkWriter {
	return &LarkWriter{
		conf: conf,
	}
}

type LarkWriter struct {
	conf Lark
}

func (l *LarkWriter) Write(p []byte) (int, error) {
	nt := client.NewFeiShuClient(l.conf.Webhook)
	err := nt.SendTextMessage(string(p))
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
