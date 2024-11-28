package console

import "go.uber.org/zap/zapcore"

type Config struct {
	Level zapcore.Level `json:"level" yaml:"level"`
	Debug bool
}
