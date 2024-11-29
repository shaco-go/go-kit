package console

import "go.uber.org/zap/zapcore"

type Config struct {
	Level zapcore.Level
	Debug bool
}
