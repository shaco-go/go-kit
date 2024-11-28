package file

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Options func(*Config)

func WithFileLevel(level string) Options {
	return func(c *Config) {
		parseLevel, err := zapcore.ParseLevel(level)
		if err == nil && c.Level.Enabled(parseLevel) {
			c.Level = parseLevel
		}
	}
}

func WithFileConfig(conf *lumberjack.Logger) Options {
	return func(c *Config) {
		c.Logger = conf
	}
}
