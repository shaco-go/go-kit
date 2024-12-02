package channel

import (
	"github.com/shaco-go/go-kit/zaplog/core"
	"go.uber.org/zap/zapcore"
	"os"
)

// NewConsoleChannel 控制台
func NewConsoleChannel(conf *core.Config) zapcore.Core {
	ec := core.NewEncoderConfig(conf.Debug)
	// 控制台修改颜色,开发模式下用console模式格式化
	var encoder zapcore.Encoder
	if conf.Debug {
		// 开发模式带有颜色
		ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(ec)
	} else {
		encoder = zapcore.NewJSONEncoder(ec)
	}
	level := conf.Level
	if conf.Level.Enabled(conf.ConsoleConf.Level) {
		level = conf.Server3Conf.Level
	}
	return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level)
}
