package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitializeLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	consoleWriter := zapcore.AddSync(os.Stdout)

	/* add file logging
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fileWriter := zapcore.AddSync(logFile)
	*/

	defaultLogLevel := zapcore.ErrorLevel
	core := zapcore.NewTee(
		//zapcore.NewCore(fileEncoder, fileWriter, defaultLogLevel), // add file logging
		zapcore.NewCore(consoleEncoder, consoleWriter, defaultLogLevel),
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}