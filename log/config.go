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
	Level   string   `json:"level" yaml:"level"`
	Channel []string `json:"channel" yaml:"channel"`
	Console Console  `json:"console" yaml:"console"`
	File    File     `json:"file" yaml:"file"`
	Lark    Lark     `json:"lark" yaml:"lark"`
	Server3 Server3  `json:"server3" yaml:"server3"`
}

type Console struct {
	Level string `json:"level" yaml:"level"`
}

type File struct {
	*lumberjack.Logger
	Level string `json:"level" yaml:"level"`
}

type Lark struct {
	Level    string `json:"level" yaml:"level"`
	Webhook  string `json:"webhook" yaml:"webhook"`
	Detailed bool   `json:"detailed" yaml:"detailed"`
}

type Server3 struct {
	Level    string `json:"level" yaml:"level"`
	SendKey  string `json:"sendkey" yaml:"sendkey"`
	Detailed bool   `json:"detailed" yaml:"detailed"`
}
