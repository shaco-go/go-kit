package file

import (
	"github.com/shaco-go/go-kit/log/channel"
	"go.uber.org/zap/zapcore"
)

// New 注册文件通道
func New(conf Config) zapcore.Core {
	config := channel.NewEncoderConfig(false)
	encoder := zapcore.NewJSONEncoder(config)
	return zapcore.NewCore(encoder, zapcore.AddSync(conf.Logger), conf.Level)
}
