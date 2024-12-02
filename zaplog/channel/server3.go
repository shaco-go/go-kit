package channel

import (
	"encoding/json"
	"fmt"
	"github.com/duke-git/lancet/v2/maputil"
	serverchan_sdk "github.com/easychen/serverchan-sdk-golang"
	"github.com/panjf2000/ants/v2"
	"github.com/shaco-go/go-kit/zaplog/core"
	"github.com/spf13/cast"
	"go.uber.org/zap/zapcore"
	"sort"
)

func NewServer3Channel(conf *core.Config) zapcore.Core {
	config := core.NewEncoderConfig(false)
	encoder := zapcore.NewJSONEncoder(config)
	level := conf.Level
	if conf.Level.Enabled(conf.Server3Conf.Level) {
		level = conf.Server3Conf.Level
	}
	return zapcore.NewCore(encoder, zapcore.AddSync(NewWriterServer3(conf.Server3Conf, config)), level)
}

func NewWriterServer3(conf core.Server3Config, encoder zapcore.EncoderConfig) *WriterServer3 {
	return &WriterServer3{
		conf:        conf,
		encoderConf: encoder,
	}
}

type WriterServer3 struct {
	conf        core.Server3Config
	encoderConf zapcore.EncoderConfig
}

func (l *WriterServer3) Write(p []byte) (int, error) {
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
