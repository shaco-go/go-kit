package log

import (
	"fmt"
	"github.com/shaco-go/go-kit/log/channel/console"
	"github.com/shaco-go/go-kit/log/channel/file"
	"github.com/shaco-go/go-kit/log/channel/lark"
	"github.com/shaco-go/go-kit/log/channel/server3"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Option func(logger *Logger)

func WithLevel(level string) Option {
	return func(l *Logger) {
		parseLevel, err := zapcore.ParseLevel(level)
		if err == nil {
			l.Conf.Level = parseLevel
		}
	}
}

func WithName(name string) Option {
	return func(l *Logger) {
		l.Conf.Name = name
	}
}

func WithDebug(debug bool) Option {
	return func(l *Logger) {
		l.Conf.Debug = debug
	}
}

// WithConsoleChannel 添加控制台通道
func WithConsoleChannel(args ...console.Options) Option {
	return func(l *Logger) {
		l.Channel = append(l.Channel, ConsoleChannel)
		l.ConsoleConf = console.Config{
			Level: l.Conf.Level,
			Debug: l.Conf.Debug,
		}
		for _, op := range args {
			op(&l.ConsoleConf)
		}
	}
}

func WithFileChannel(args ...file.Options) Option {
	return func(l *Logger) {
		l.Channel = append(l.Channel, FileChannel)
		l.FileConf = file.Config{
			Logger: &lumberjack.Logger{
				Filename:   fmt.Sprintf("log/%s.log", l.Conf.Name),
				MaxBackups: 50,
				MaxAge:     30, // days
			},
			Level: l.Conf.Level,
		}
		for _, op := range args {
			op(&l.FileConf)
		}
	}
}

func WithLarkChannel(webhook string, args ...lark.Options) Option {
	return func(l *Logger) {
		l.Channel = append(l.Channel, LarkChannel)
		l.LarkConf = lark.Config{
			Name:    l.Conf.Name,
			Level:   l.Conf.Level,
			Webhook: webhook,
		}
		for _, op := range args {
			op(&l.LarkConf)
		}
	}
}

func WithServer3Channel(sendKey string, args ...server3.Options) Option {
	return func(l *Logger) {
		l.Channel = append(l.Channel, Server3Channel)
		l.Server3Conf = server3.Config{
			Name:    l.Conf.Name,
			Level:   l.Conf.Level,
			SendKey: sendKey,
		}
		for _, op := range args {
			op(&l.Server3Conf)
		}
	}
}
