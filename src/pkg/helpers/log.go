package helper

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


var Logger *zap.Logger

func InitLogger() {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncode := zapcore.NewJSONEncoder(config)
	consoleEncode := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	logLevel := zapcore.InfoLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncode, writer, logLevel),
		zapcore.NewCore(consoleEncode, zapcore.AddSync(os.Stdout), logLevel),
	)

	Logger= zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}