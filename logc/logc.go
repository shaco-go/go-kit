package logc

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func New(conf *Config) *zap.Logger {
	logc := &Logc{Conf: conf}
	return logc.NewZap()
}

type Logc struct {
	Conf *Config
}

// EncoderConfig zap EncoderConfig
func (l *Logc) EncoderConfig() zapcore.EncoderConfig {
	var encoder zapcore.EncoderConfig
	if l.Conf.Env == Dev {
		encoder = zap.NewDevelopmentEncoderConfig()
	} else {
		encoder = zap.NewProductionEncoderConfig()
		encoder.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	return encoder
}

// RegisterConsole 注册控制台
func (l *Logc) RegisterConsole() zapcore.Core {
	level := l.Conf.Level
	if l.Conf.Console.Level > level {
		level = l.Conf.Console.Level
	}
	encoderConfig := l.EncoderConfig()
	// 控制台修改颜色,开发模式下用console模式格式化
	var encoder zapcore.Encoder
	if l.Conf.Env == Dev {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
}

// RegisterFile 注册文件通道
func (l *Logc) RegisterFile() zapcore.Core {
	level := l.Conf.Level
	if l.Conf.Console.Level > level {
		level = l.Conf.Console.Level
	}
	encoderConfig := l.EncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	// 如果文件名称为空，使用默认配置
	l.FileConfDefault()
	ws := &lumberjack.Logger{
		Filename:   l.Conf.File.Filename,
		MaxSize:    l.Conf.File.MaxSize,
		MaxAge:     l.Conf.File.MaxAge,
		MaxBackups: l.Conf.File.MaxBackups,
		LocalTime:  l.Conf.File.LocalTime,
		Compress:   l.Conf.File.Compress,
	}
	return zapcore.NewCore(encoder, zapcore.AddSync(ws), level)
}

// RegisterLark 注册飞书通道
func (l *Logc) RegisterLark() zapcore.Core {
	level := l.Conf.Level
	if l.Conf.Console.Level > level {
		level = l.Conf.Console.Level
	}
	encoderConfig := l.EncoderConfig()
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewCore(encoder, zapcore.AddSync(NewLarkWriter(l.Conf.Lark)), level)
}

// FileConfDefault 设置默认配置
func (l *Logc) FileConfDefault() {
	if l.Conf.File.Filename == "" {
		l.Conf.File.Filename = "./log/app.log"
		l.Conf.File.MaxSize = 10
		l.Conf.File.MaxAge = 30
		l.Conf.File.MaxBackups = 100
	}
}

// NewZap 生成zap实例
func (l *Logc) NewZap() *zap.Logger {
	var core []zapcore.Core
	for _, item := range l.Conf.Channel {
		switch item {
		case ConsoleChannel:
			core = append(core, l.RegisterConsole())
		case FileChannel:
			core = append(core, l.RegisterFile())
		case LarkChannel:
			core = append(core, l.RegisterLark())
		case DingTalkChannel:
			panic("unrealized channel")
		case WeComChannel:
			panic("unrealized channel")
		default:
			panic("unknown channel")
		}
	}
	return zap.New(zapcore.NewTee(core...))
}
