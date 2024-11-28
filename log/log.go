package log

import (
	"github.com/shaco-go/go-kit/log/channel/console"
	"github.com/shaco-go/go-kit/log/channel/file"
	"github.com/shaco-go/go-kit/log/channel/lark"
	"github.com/shaco-go/go-kit/log/channel/server3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(arg ...Option) *zap.Logger {
	logger := &Logger{Conf: &Config{
		Name:  "app",
		Debug: false,
		Level: zap.DebugLevel,
	}}
	for _, op := range arg {
		op(logger)
	}
	if len(logger.Channel) == 0 {
		WithConsoleChannel()(logger)
		WithFileChannel()(logger)
	}
	return logger.newZap()
}

type Logger struct {
	Conf        *Config
	Channel     []Channel
	ConsoleConf console.Config
	FileConf    file.Config
	LarkConf    lark.Config
	Server3Conf server3.Config
}

// ParseLevel 解析level,如果级别比默认高,优先使用
func (l *Logger) ParseLevel(arg ...string) zapcore.Level {
	level := l.Conf.Level
	if len(arg) > 0 {
		if val, err := zapcore.ParseLevel(arg[0]); err == nil {
			if level.Enabled(val) {
				// 指定level比默认级别高,使用指定level
				level = val
			}
		}
	}
	return level
}

// NewZap 生成zap实例
func (l *Logger) newZap() *zap.Logger {
	var core []zapcore.Core
	for _, item := range l.Channel {
		switch item {
		case ConsoleChannel:
			core = append(core, console.New(l.ConsoleConf))
		case FileChannel:
			core = append(core, file.New(l.FileConf))
		case LarkChannel:
			core = append(core, lark.New(l.LarkConf))
		case Server3Channel:
			core = append(core, server3.New(l.Server3Conf))
		default:
			panic("unknown channel")
		}
	}
	return zap.New(zapcore.NewTee(core...))
}
