package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewEncoderConfig zap EncoderConfig
func NewEncoderConfig(debug bool) zapcore.EncoderConfig {
	var encoder zapcore.EncoderConfig
	if debug {
		encoder = zap.NewDevelopmentEncoderConfig()
	} else {
		encoder = zap.NewProductionEncoderConfig()
		encoder.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	return encoder
}
