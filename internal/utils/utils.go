package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitializeLogger: init ZAP logger
func InitializeLogger(isDebug bool) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	consoleWriter := zapcore.AddSync(os.Stdout)

	/* add file logging
	fileEncoder := zapcore.NewJSONEncoder(config)
	logFile, _ := os.OpenFile(time.Now().Format("2006-01-02)+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fileWriter := zapcore.AddSync(logFile)
	*/

	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		//zapcore.NewCore(fileEncoder, fileWriter, defaultLogLevel), // add file logging
		zapcore.NewCore(consoleEncoder, consoleWriter, defaultLogLevel),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	if isDebug{
		core := zapcore.NewTee(
			//zapcore.NewCore(fileEncoder, fileWriter, defaultLogLevel), // add file logging
			zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel),
		)
		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.DebugLevel), zap.Development())
		//Logger, _ = zap.NewDevelopment(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	} else {
		core := zapcore.NewTee(
			//zapcore.NewCore(fileEncoder, fileWriter, defaultLogLevel), // add file logging
			zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel),
		)
		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		//Logger, _ = zap.NewProduction(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}
}