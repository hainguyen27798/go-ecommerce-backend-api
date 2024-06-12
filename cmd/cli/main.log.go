package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	encoder := getEncodeLog()
	sync := getWriteSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log")
	logger.Error("Error log")
}

func getEncodeLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// format timestamp
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> Time
	encoderConfig.TimeKey = "time"

	// info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller": "cli/main.log.go:3"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./logs/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	fileSync := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stdout)
	return zapcore.NewMultiWriteSyncer(fileSync, syncConsole)
}
