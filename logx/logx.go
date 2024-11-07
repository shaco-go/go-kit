package logx

import (
	"context"
	"go.uber.org/zap"
)

func New(logger *zap.Logger, handle ...HandleContext) *Logx {
	logx := Logx{
		logger:        logger,
		handleContext: DefaultHandleContext,
	}
	if len(handle) > 0 {
		logx.handleContext = handle[0]
	}
	return &logx
}

type HandleContext func(ctx context.Context, logger *zap.Logger) *zap.Logger

var DefaultHandleContext HandleContext = func(ctx context.Context, logger *zap.Logger) *zap.Logger {
	return logger
}

type Logx struct {
	logger        *zap.Logger
	handleContext HandleContext
}

func (l *Logx) WithContext(ctx context.Context) *zap.Logger {
	return l.handleContext(ctx, l.logger)
}

func (l *Logx) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.WithContext(ctx).Debug(msg, fields...)
}
func (l *Logx) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.WithContext(ctx).Info(msg, fields...)
}
func (l *Logx) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.WithContext(ctx).Warn(msg, fields...)
}
func (l *Logx) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.WithContext(ctx).Error(msg, fields...)
}
func (l *Logx) DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	l.WithContext(ctx).DPanic(msg, fields...)
}
func (l *Logx) Panic(ctx context.Context, msg string, fields ...zap.Field) {
	l.WithContext(ctx).Panic(msg, fields...)
}
func (l *Logx) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	l.WithContext(ctx).Fatal(msg, fields...)
}

func (l *Logx) Debugf(ctx context.Context, err error) {
	l.WithContext(ctx).Sugar().Debugf("%+v", err)
}
func (l *Logx) Infof(ctx context.Context, err error) {
	l.WithContext(ctx).Sugar().Infof("%+v", err)
}
func (l *Logx) Warnf(ctx context.Context, err error) {
	l.WithContext(ctx).Sugar().Warnf("%+v", err)
}
func (l *Logx) Errorf(ctx context.Context, err error) {
	l.WithContext(ctx).Sugar().Errorf("%+v", err)
}
func (l *Logx) DPanicf(ctx context.Context, err error) {
	l.WithContext(ctx).Sugar().DPanicf("%+v", err)
}
func (l *Logx) Panicf(ctx context.Context, err error) {
	l.WithContext(ctx).Sugar().Panicf("%+v", err)
}
func (l *Logx) Fatalf(ctx context.Context, err error) {
	l.WithContext(ctx).Sugar().Fatalf("%+v", err)
}
