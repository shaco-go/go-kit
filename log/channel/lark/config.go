package lark

import "go.uber.org/zap/zapcore"

type Config struct {
	Name     string
	Level    zapcore.Level
	Webhook  string
	Detailed bool
}
