package server3

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/maputil"
	serverchan_sdk "github.com/easychen/serverchan-sdk-golang"
	"github.com/panjf2000/ants/v2"
	"github.com/shaco-go/go-kit/log/channel"
	"github.com/spf13/cast"
	"go.uber.org/zap/zapcore"
	"sort"
)

func New(conf Config) zapcore.Core {
	config := channel.NewEncoderConfig(false)
	encoder := zapcore.NewJSONEncoder(config)
	return zapcore.NewCore(encoder, zapcore.AddSync(NewWriter(conf, config)), conf.Level)
}

func NewWriter(conf Config, encoder zapcore.EncoderConfig) *Writer {
	return &Writer{
		conf:        conf,
		encoderConf: encoder,
	}
}

type Writer struct {
	conf        Config
	encoderConf zapcore.EncoderConfig
}

func (l *Writer) Write(p []byte) (int, error) {
	err := ants.Submit(func() {
		var data map[string]any
		_ = json.Unmarshal(p, &data)

		if !l.conf.Detailed {
			// 是否完整
			data = maputil.FilterByKeys(data, []string{l.encoderConf.LevelKey, l.encoderConf.MessageKey,
				l.encoderConf.TimeKey, "err", "error"})
		}

		msg := cast.ToString(data[l.encoderConf.MessageKey])
		level := cast.ToString(data[l.encoderConf.LevelKey])
		mk := maputil.Keys(data)
		sort.Strings(mk)
		var md string
		for _, key := range mk {
			if val, ok := data[key]; ok {
				md += fmt.Sprintf("**%s**：%s\n", key, val)
			}
		}
		_, err := serverchan_sdk.ScSend(l.conf.SendKey, l.conf.Name, md, &serverchan_sdk.ScSendOptions{
			Tags:  fmt.Sprintf("%s|%s", level, l.conf.Name),
			Short: msg,
		})
		if err != nil {
			fmt.Println(err)
		}
	})
	if err != nil {
		fmt.Println(err)
	}
	return len(p), nil
}
