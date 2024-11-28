package server3

import "go.uber.org/zap/zapcore"

type Config struct {
	Name     string
	Level    zapcore.Level
	SendKey  string
	Detailed bool
}
