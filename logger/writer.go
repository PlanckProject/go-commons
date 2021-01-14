package logger

import (
	"io"

	"github.com/natefinch/lumberjack"
)

func GetRotatedWriter(config *Config) io.Writer {
	return &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    int(config.MaxSize),
		MaxBackups: int(config.MaxBackups),
		MaxAge:     int(config.MaxAge),
		Compress:   config.Compress,
	}
}
