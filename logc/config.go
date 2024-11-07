package logc

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Channel int

const (
	ConsoleChannel Channel = iota
	FileChannel
	LarkChannel     // 飞书
	DingTalkChannel // 钉钉
	WeComChannel    // 企业微信
)

type Env int

const (
	Dev Env = iota
	Prod
)

type Config struct {
	Env     Env           `json:"env" yaml:"env"`
	Level   zapcore.Level `json:"level" yaml:"level"`
	Channel []Channel     `json:"channel" yaml:"channel"`
	Console Console       `json:"console" yaml:"console"`
	Lark    Lark          `json:"lark" yaml:"lark"`
	File    File          `json:"file" yaml:"file"`
}

// Lark 飞书
type Lark struct {
	Webhook string        `json:"webhook"`
	Level   zapcore.Level `json:"level" yaml:"level"`
}

type Console struct {
	Level zapcore.Level `json:"level" yaml:"level"`
}

type File struct {
	// Filename 是要写入日志的文件。备份日志文件将保留在同一目录中。
	// 如果为空，它将使用 <processname>-lumberjack.log 在 os.TempDir() 中。
	Filename string `json:"filename" yaml:"filename"`

	// MaxSize 是日志文件在轮换之前的最大大小（以兆字节为单位）。默认为 100 兆字节。
	MaxSize int `json:"max_size" yaml:"max_size"`

	// MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
	// 注意，一天被定义为 24 小时，可能由于夏令时、闰秒等原因与日历天数不完全对应。
	// 默认情况下，不会根据年龄删除旧日志文件。
	MaxAge int `json:"max_age" yaml:"max_age"`

	// MaxBackups 是要保留的旧日志文件的最大数量。默认情况下保留所有旧日志文件
	// （尽管 MaxAge 仍可能导致它们被删除）。
	MaxBackups int `json:"max_backups" yaml:"max_backups"`

	// LocalTime 决定用于格式化备份文件中时间戳的时间是否为计算机的本地时间。
	// 默认情况下使用 UTC 时间。
	LocalTime bool `json:"local_time" yaml:"local_time"`

	// Compress 决定轮换的日志文件是否应使用 gzip 压缩。默认情况下不进行压缩。
	Compress bool `json:"compress" yaml:"compress"`

	Level zapcore.Level `json:"level" yaml:"level"`
}

// ConvChannel 转换
func ConvChannel(channel []string) []Channel {
	var c []Channel
	for _, item := range channel {
		switch strings.ToLower(item) {
		case "console":
			c = append(c, ConsoleChannel)
		case "file":
			c = append(c, FileChannel)
		case "lark":
			c = append(c, LarkChannel)
		case "dingtalk":
			c = append(c, DingTalkChannel)
		case "wecom":
			c = append(c, WeComChannel)
		default:
			panic("unknown channel")
		}
	}
	return c
}

// ConvEnv 转换
func ConvEnv(env string) Env {
	switch strings.ToLower(env) {
	case "dev":
		return Dev
	case "prod":
		return Prod
	default:
		return Prod
	}
}

// ConvLevel 转换
func ConvLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
