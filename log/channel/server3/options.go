package server3

import (
	"go.uber.org/zap/zapcore"
)

type Options func(*Config)

func WithLarkLevel(level string) Options {
	return func(c *Config) {
		parseLevel, err := zapcore.ParseLevel(level)
		if err == nil && c.Level.Enabled(parseLevel) {
			c.Level = parseLevel
		}
	}
}

func WithLarkDetailed(d bool) Options {
	return func(c *Config) {
		c.Detailed = d
	}
}
