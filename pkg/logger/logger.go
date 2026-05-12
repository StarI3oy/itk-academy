package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level zapcore.Level, debug bool) *zap.Logger {
	var logCfg zapcore.EncoderConfig

	if debug {
		logCfg = zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "lvl",
			TimeKey:        "ts",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			SkipLineEnding: false,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout(time.RFC3339),
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
	} else {
		logCfg = zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "lvl",
			TimeKey:        "ts",
			SkipLineEnding: false,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout(time.RFC3339),
		}
	}

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(logCfg), zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(level))

	return zap.New(core, zap.AddCaller())
}
