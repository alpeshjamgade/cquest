package logger

import (
	"context"
	"cquest/config"
	"cquest/internal/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger
var loggerCount int32

func CreateLoggerWithCtx(ctx context.Context) *zap.SugaredLogger {
	if logger == nil {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg := zap.Config{
			Level:             GetLevel(),
			Development:       false,
			DisableCaller:     false,
			DisableStacktrace: false,
			Sampling:          nil,
			Encoding:          config.LogEncoding,
			EncoderConfig:     encoderCfg,
			OutputPaths:       []string{"stderr"},
			ErrorOutputPaths:  []string{"stderr"},
		}
		logger = zap.Must(cfg.Build()).Sugar()
		loggerCount++
		logger.Debugf("create new logger, count: %d", loggerCount)
	}

	if ctx != nil && ctx.Value(constants.TraceID) != nil {
		traceId := ctx.Value(constants.TraceID).(string)
		return logger.WithOptions(zap.Fields(zap.String(constants.TraceID, traceId), zap.String(constants.Service, constants.ServiceName)))
	}

	return logger
}

func GetLevel() zap.AtomicLevel {
	switch config.LogLevel {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case "dpanic":
		return zap.NewAtomicLevelAt(zap.DPanicLevel)
	}
	return zap.NewAtomicLevelAt(zap.InfoLevel)
}
