package channel

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/panjf2000/ants/v2"
	"github.com/shaco-go/go-kit/notify/lark"
	"github.com/shaco-go/go-kit/zaplog/core"
	"github.com/spf13/cast"
	"go.uber.org/zap/zapcore"
)

func NewLarkChannel(conf *core.Config) zapcore.Core {
	config := core.NewEncoderConfig(false)
	encoder := zapcore.NewJSONEncoder(config)
	return zapcore.NewCore(encoder, zapcore.AddSync(NewWriterLark(conf.LarkConf, config)), conf.LarkConf.Level)
}

func NewWriterLark(conf core.LarkConfig, encoder zapcore.EncoderConfig) *WriterLark {
	return &WriterLark{
		conf:        conf,
		encoderConf: encoder,
	}
}

type WriterLark struct {
	conf        core.LarkConfig
	encoderConf zapcore.EncoderConfig
}

func (l *WriterLark) Write(p []byte) (int, error) {
	err := ants.Submit(func() {
		var data map[string]any
		_ = json.Unmarshal(p, &data)
		msg := &lark.LarkMsg{
			Title:       l.conf.Name,
			Markdown:    data,
			HeaderColor: lark.ColorDefault,
		}

		if val, ok := data[l.encoderConf.LevelKey]; ok {
			// 颜色级别
			level, err := zapcore.ParseLevel(cast.ToString(val))
			if err == nil {
				// 颜色更改
				if level >= zapcore.ErrorLevel {
					msg.HeaderColor = lark.ColorRed
				} else if level >= zapcore.WarnLevel {
					msg.HeaderColor = lark.ColorYellow
				} else if level >= zapcore.InfoLevel {
					msg.HeaderColor = lark.ColorBlue
				}
			}
		}

		if !l.conf.Detailed {
			// 是否完整
			msg.Markdown = maputil.FilterByKeys(data, []string{l.encoderConf.LevelKey, l.encoderConf.MessageKey,
				l.encoderConf.TimeKey, "err", "error"})
		}

		err := lark.SendLarkMsg(l.conf.Webhook, msg)
		if err != nil {
			fmt.Println(err)
		}
	})
	if err != nil {
		fmt.Println(err)
	}
	return len(p), nil
}
