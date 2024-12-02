package zaplog

import (
	"github.com/shaco-go/go-kit/zaplog/channel"
	"github.com/shaco-go/go-kit/zaplog/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New 创建新的日志记录器
func New(config *Config) *zap.Logger {
	conf := toCoreConfig(config)
	var cores []zapcore.Core
	for _, channelItem := range conf.Channel {
		switch channelItem {
		case core.FileChannel:
			cores = append(cores, channel.NewFileChannel(conf))
		case core.ConsoleChannel:
			cores = append(cores, channel.NewConsoleChannel(conf))
		case core.LarkChannel:
			cores = append(cores, channel.NewLarkChannel(conf))
		case core.Server3Channel:
			cores = append(cores, channel.NewServer3Channel(conf))
		default:
			panic("invalid channel")
		}
	}
	return zap.New(zapcore.NewTee(cores...))
}
