package log

import (
	"github.com/shaco-go/go-kit/log/channel/console"
	"github.com/shaco-go/go-kit/log/channel/file"
	"github.com/shaco-go/go-kit/log/channel/lark"
	"github.com/shaco-go/go-kit/log/channel/server3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Default() *Logger {
	logger := &Logger{}
	logger.Use(
		WithName("app"),
		WithLevel("debug"),
		WithDebug(false),
	)
	// 通道
	WithConsoleChannel(
		console.WithConsoleLevel("debug"),
	)(logger)
	WithFileChannel(
		file.WithFileLevel("info"),
		file.WithFileConfig(&lumberjack.Logger{
			Filename:   "log/app.log",
			MaxSize:    100,
			MaxAge:     30,
			MaxBackups: 30,
		}),
	)(logger)
	return logger
}

func New(conf *Config) *Logger {
	logger := &Logger{}

	logger.Use(
		WithName(conf.Name),
		WithLevel(conf.Level),
		WithDebug(conf.Debug),
	)

	// 通道
	for _, item := range conf.Channel {
		c, err := ParseChannel(item)
		if err == nil {
			switch c {
			case ConsoleChannel:
				WithConsoleChannel(
					console.WithConsoleLevel(conf.Console.Level),
				)(logger)
			case FileChannel:
				WithFileChannel(
					file.WithFileLevel(conf.File.Level),
					file.WithFileConfig(conf.File.Logger),
				)(logger)
			case LarkChannel:
				WithLarkChannel(
					conf.Lark.Webhook,
					lark.WithLarkLevel(conf.Lark.Level),
					lark.WithLarkDetailed(conf.Lark.Detailed),
				)(logger)
			case Server3Channel:
				WithServer3Channel(
					conf.Server3.SendKey,
					server3.WithLarkLevel(conf.Server3.Level),
					server3.WithLarkDetailed(conf.Server3.Detailed),
				)(logger)
			default:
				continue
			}
		}
	}
	return logger
}

type Logger struct {
	Conf        config
	Channel     []Channel
	ConsoleConf console.Config
	FileConf    file.Config
	LarkConf    lark.Config
	Server3Conf server3.Config
}

func (l *Logger) Use(arg ...Option) *Logger {
	for _, op := range arg {
		op(l)
	}
	return l
}

// Zap 生成zap实例
func (l *Logger) Zap() *zap.Logger {
	var core []zapcore.Core
	for _, item := range l.Channel {
		switch item {
		case ConsoleChannel:
			if l.ConsoleConf.Level.Enabled(l.Conf.Level) {
				l.ConsoleConf.Level = l.Conf.Level
			}
			l.ConsoleConf.Debug = l.Conf.Debug
			core = append(core, console.New(l.ConsoleConf))
		case FileChannel:
			if l.FileConf.Level.Enabled(l.Conf.Level) {
				l.FileConf.Level = l.Conf.Level
			}
			core = append(core, file.New(l.FileConf))
		case LarkChannel:
			if l.LarkConf.Level.Enabled(l.Conf.Level) {
				l.LarkConf.Level = l.Conf.Level
			}
			core = append(core, lark.New(l.LarkConf))
		case Server3Channel:
			if l.Server3Conf.Level.Enabled(l.Conf.Level) {
				l.Server3Conf.Level = l.Conf.Level
			}
			core = append(core, server3.New(l.Server3Conf))
		default:
			panic("unknown channel")
		}
	}
	return zap.New(zapcore.NewTee(core...))
}
