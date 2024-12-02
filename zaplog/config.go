package zaplog

import (
	"github.com/shaco-go/go-kit/zaplog/core"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var DefaultConfig = &Config{
	Name:    "app",
	Debug:   false,
	Level:   "info",
	Channel: []string{"console", "file"},
	Console: Console{
		Level: "info",
	},
	File: File{
		Logger: &lumberjack.Logger{
			Filename:   "log/app.log",
			MaxSize:    100,
			MaxAge:     30,
			MaxBackups: 30,
		},
		Level: "info",
	},
}

type Config struct {
	Name    string
	Debug   bool
	Level   string
	Channel []string
	Console Console
	File    File
	Lark    Lark
	Server3 Server3
}

type Console struct {
	Level string
}

type File struct {
	*lumberjack.Logger
	Level string
}

type Lark struct {
	Level    string
	Webhook  string
	Detailed bool
}

type Server3 struct {
	Level    string
	SendKey  string
	Detailed bool
}

func toCoreConfig(config *Config) *core.Config {
	c := &core.Config{
		Name:  config.Name,
		Debug: config.Debug,
		// Channel: nil,
		ConsoleConf: core.ConsoleConfig{
			Debug: config.Debug,
		},
		FileConf: core.FileConfig{
			Logger: config.File.Logger,
			Debug:  config.Debug,
		},
		LarkConf: core.LarkConfig{
			Name:     config.Name,
			Webhook:  config.Lark.Webhook,
			Detailed: config.Lark.Detailed,
		},
		Server3Conf: core.Server3Config{
			Name:     config.Name,
			SendKey:  config.Server3.SendKey,
			Detailed: config.Server3.Detailed,
		},
	}
	c.Level, _ = zapcore.ParseLevel(config.Level)
	c.ConsoleConf.Level, _ = zapcore.ParseLevel(config.Console.Level)
	c.FileConf.Level, _ = zapcore.ParseLevel(config.File.Level)
	c.LarkConf.Level, _ = zapcore.ParseLevel(config.Lark.Level)
	c.Server3Conf.Level, _ = zapcore.ParseLevel(config.Server3.Level)
	if c.ConsoleConf.Level.Enabled(c.Level) {
		c.ConsoleConf.Level = c.Level
	}
	if c.FileConf.Level.Enabled(c.Level) {
		c.FileConf.Level = c.Level
	}
	if c.LarkConf.Level.Enabled(c.Level) {
		c.LarkConf.Level = c.Level
	}
	if c.Server3Conf.Level.Enabled(c.Level) {
		c.Server3Conf.Level = c.Level
	}
	for _, item := range config.Channel {
		if val, err := core.ParseChannel(item); err == nil {
			c.Channel = append(c.Channel, val)
		}
	}
	return c
}
