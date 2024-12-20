package channel

import (
	"github.com/shaco-go/go-kit/zaplog/core"
	"go.uber.org/zap/zapcore"
)

// NewFileChannel 注册文件通道
func NewFileChannel(conf *core.Config) zapcore.Core {
	config := core.NewEncoderConfig(false)
	encoder := zapcore.NewJSONEncoder(config)
	return zapcore.NewCore(encoder, zapcore.AddSync(conf.FileConf.Logger), conf.FileConf.Level)
}
