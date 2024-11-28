package log

import (
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Name  string
	Debug bool
	Level zapcore.Level
}
