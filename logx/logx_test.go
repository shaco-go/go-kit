package logx

import (
	"context"
	"github.com/shaco-go/go-kit/logc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestWithContent(t *testing.T) {
	logger := logc.New(&logc.Config{
		Env:     logc.Dev,
		Level:   zapcore.DebugLevel,
		Channel: []logc.Channel{logc.ConsoleChannel},
	})
	logx := New(logger, func(ctx context.Context, logger *zap.Logger) *zap.Logger {
		rid := ctx.Value("rid").(string)
		return logger.With(zap.String("rid", rid))
	})
	ctx := context.WithValue(context.Background(), "rid", "26KRrzS938FH")
	logx.WithContext(ctx).Info("rid")
}
