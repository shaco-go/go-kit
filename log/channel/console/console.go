package console

import (
	"fmt"
	"github.com/shaco-go/go-kit/log/channel"
	"go.uber.org/zap/zapcore"
	"os"
)

// New 控制台
func New(conf Config) zapcore.Core {
	ec := channel.NewEncoderConfig(conf.Debug)
	// 控制台修改颜色,开发模式下用console模式格式化
	var encoder zapcore.Encoder
	if conf.Debug {
		// 开发模式带有颜色
		ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(ec)
	} else {
		encoder = zapcore.NewJSONEncoder(ec)
	}
	return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), conf.Level)
}

type ConsoleEncoder struct {
}

func (c ConsoleEncoder) Encode(i interface{}) error {
	fmt.Println(i)
	return nil
	panic("implement me")
}
