package core

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Name        string
	Debug       bool
	Level       zapcore.Level
	Channel     []Channel
	ConsoleConf ConsoleConfig
	FileConf    FileConfig
	LarkConf    LarkConfig
	Server3Conf Server3Config
}

type ConsoleConfig struct {
	Level zapcore.Level
	Debug bool
}

type FileConfig struct {
	*lumberjack.Logger
	Level zapcore.Level
	Debug bool
}

type LarkConfig struct {
	Name     string
	Level    zapcore.Level
	Webhook  string
	Detailed bool
}

type Server3Config struct {
	Name     string
	Level    zapcore.Level
	SendKey  string
	Detailed bool
}
