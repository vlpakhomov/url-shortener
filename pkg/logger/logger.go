package logger

import (
	"context"

	"github.com/VlPakhomov/url_shortener/internal/config"
	"go.uber.org/zap"
)

type ctxKey string

const keyLogger ctxKey = "user_id"

func init() {
	var logger *zap.Logger
	if config.Get(config.LogLevel) == "debug" {
		logger = zap.Must(zap.NewDevelopment())
	}
	if config.Get(config.LogLevel) == "info" {
		logger = zap.Must(zap.NewProduction())
	}
	zap.ReplaceGlobals(logger)
}

func Ctx(ctx context.Context) context.Context {
	return context.WithValue(ctx, keyLogger, zap.S())
}

func Info(ctx context.Context, args ...interface{}) {
	if sl, ok := ctx.Value(keyLogger).(*zap.SugaredLogger); ok {
		sl.Info(args...)
	} else {
		zap.S().Info(args...)
	}
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	if sl, ok := ctx.Value(keyLogger).(*zap.SugaredLogger); ok {
		sl.Infof(template, args...)
	} else {
		zap.S().Infof(template, args...)
	}
}

func Debug(ctx context.Context, args ...interface{}) {
	if sl, ok := ctx.Value(keyLogger).(*zap.SugaredLogger); ok {
		sl.Debug(args...)
	} else {
		zap.S().Debug(args...)
	}
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	if sl, ok := ctx.Value(keyLogger).(*zap.SugaredLogger); ok {
		sl.Debugf(template, args...)
	} else {
		zap.S().Debugf(template, args...)
	}
}

func Error(ctx context.Context, args ...interface{}) {
	if sl, ok := ctx.Value(keyLogger).(*zap.SugaredLogger); ok {
		sl.Error(args...)
	} else {
		zap.S().Error(args...)
	}
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	if sl, ok := ctx.Value(keyLogger).(*zap.SugaredLogger); ok {
		sl.Errorf(template, args...)
	} else {
		zap.S().Errorf(template, args...)
	}
}

func Fatal(ctx context.Context, args ...interface{}) {
	if sl, ok := ctx.Value(keyLogger).(*zap.SugaredLogger); ok {
		sl.Fatal(args...)
	} else {
		zap.S().Fatal(args...)
	}
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	if sl, ok := ctx.Value(keyLogger).(*zap.SugaredLogger); ok {
		sl.Fatalf(template, args...)
	} else {
		zap.S().Fatalf(template, args...)
	}
}
