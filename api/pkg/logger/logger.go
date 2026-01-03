package logger

import (
	"os"

	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	var config zap.Config

	if os.Getenv("APP_ENV") == "dev" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	logFile := os.Getenv("APP_LOG_FILE")
	if logFile != "" {
		config.OutputPaths = append(config.OutputPaths, logFile)
	}

	l, err := config.Build()
	if err != nil {
		panic(err)
	}
	logger = l.Sugar()
}

func Sync() {
	_ = logger.Sync()
}

func Info(msg string, keysAndValues ...any) {
	logger.Infow(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...any) {
	logger.Debugw(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...any) {
	logger.Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...any) {
	logger.Errorw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...any) {
	logger.Fatalw(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...any) {
	logger.Panicw(msg, keysAndValues...)
}
