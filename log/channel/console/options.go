package console

import "go.uber.org/zap/zapcore"

type Options func(*Config)

func WithConsoleLevel(level string) Options {
	return func(c *Config) {
		parseLevel, err := zapcore.ParseLevel(level)
		if err == nil {
			c.Level = parseLevel
		}
	}
}
