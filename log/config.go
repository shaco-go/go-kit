package log

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type config struct {
	Name  string
	Debug bool
	Level zapcore.Level
}

type Config struct {
	Name    string
	Debug   bool
	Level   string
	Channel []string
	Console Console
	File    File
	Lark    Lark
	Server3 Server3
}

type Console struct {
	Level string
}

type File struct {
	*lumberjack.Logger
	Level string
}

type Lark struct {
	Level    string
	Webhook  string
	Detailed bool
}

type Server3 struct {
	Level    string
	SendKey  string
	Detailed bool
}
