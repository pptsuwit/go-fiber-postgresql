package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func InitLogs() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	cfg.EncoderConfig.StacktraceKey = ""
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	log, _ = cfg.Build(zap.AddCallerSkip(1))
}

func Info(message string, field ...zap.Field) {
	log.Info(message, field...)
}

func Debug(message string, field ...zap.Field) {
	log.Debug(message, field...)
}
func Error(message interface{}, field ...zap.Field) {
	switch msg := message.(type) {
	case error:
		log.Error(msg.Error(), field...)
		break
	case string:
		log.Error(msg, field...)
		break
	}
}
