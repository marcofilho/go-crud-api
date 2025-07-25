package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger

	LOG_LEVEL  = "LOG_LEVEL"
	LOG_OUTPUT = "LOG_OUTPUT"

	LEVEL_KEY   = "level"
	TIME_KEY    = "time"
	MESSAGE_KEY = "message"
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputLogger()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     LEVEL_KEY,
			TimeKey:      TIME_KEY,
			MessageKey:   MESSAGE_KEY,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	log, _ = logConfig.Build()
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))

	log.Error(message, tags...)
	log.Sync()
}

func getOutputLogger() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		output = "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	level := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL)))

	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn", "warning":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
