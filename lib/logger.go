package lib

import (
	"api.service.com/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func InitLogger() {
	conf := config.GetConfig()
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.GetString("log.file-name"),
		MaxSize:    conf.GetInt("log.max-size"), // megabytes
		MaxBackups: conf.GetInt("log.max-backups"),
		MaxAge:     conf.GetInt("log.max-age"), // days
		LocalTime:  true,
	})
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.TimeKey = "time"
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		w,
		zap.InfoLevel,
	)
	logger = zap.New(core)
}

func GetLogger() *zap.Logger {
	return logger
}
