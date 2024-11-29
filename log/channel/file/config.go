package file

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	*lumberjack.Logger
	Level zapcore.Level
}
