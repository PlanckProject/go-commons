package logger

import (
	"context"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

var instance Logger

type (
	LogEntry interface {
		Debug(...interface{})
		Debugf(string, ...interface{})
		Debugln(...interface{})
		Error(...interface{})
		Errorf(string, ...interface{})
		Errorln(...interface{})
		Fatal(...interface{})
		Fatalf(string, ...interface{})
		Fatalln(...interface{})
		Info(...interface{})
		Infof(string, ...interface{})
		Infoln(...interface{})
		Trace(...interface{})
		Tracef(string, ...interface{})
		Traceln(...interface{})
		Panic(...interface{})
		Panicf(string, ...interface{})
		Panicln(...interface{})
		Print(...interface{})
		Printf(string, ...interface{})
		Println(...interface{})
		Warn(...interface{})
		Warnf(string, ...interface{})
		Warnln(...interface{})

		WithField(string, interface{}) LogEntry
		WithFields(Fields) LogEntry
		WithContext(context.Context) LogEntry
	}

	Logger interface {
		LogEntry

		AddHook(interface{}) error
		NewEntry() LogEntry
		SetWriter(io.Writer)
		SetReportCaller(bool)
		SetLevel(string)
		SetFormatter(string)
	}

	// Config represents logger configuration
	Config struct {
		Base         string `mapstructure:"base"`
		Level        string `mapstructure:"level"`
		Format       string `mapstructure:"format"`
		Enabled      bool   `mapstructure:"enabled"`
		MaxAge       uint   `mapstructure:"max_age"` // days
		MaxBackups   uint   `mapstructure:"max_backups"`
		MaxSize      uint   `mapstructure:"max_size"` // MBs
		Compress     bool   `mapstructure:"compress"`
		ReportCaller bool   `mapstructure:"report_caller"`
		Filename     string `mapstructure:"filename"`
	}

	Fields map[string]interface{}
)

func init() {
	instance = newLogrusLogger()
}

func Configure(config *Config, writer io.Writer) {
	switch config.Base {
	case "logrus":
		instance = newLogrusLogger()
	}

	if config.Level == "" {
		config.Level = "debug"
	}
	instance.SetLevel(config.Level)

	instance.SetReportCaller(config.ReportCaller)

	if config.Format == "" {
		config.Format = "json"
	}
	instance.SetFormatter(config.Format)

	if config.Enabled {
		instance.SetWriter(writer)
	}
}

func Debug(args ...interface{}) {
	instance.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	instance.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	instance.Debugln(args...)
}

func Error(args ...interface{}) {
	instance.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	instance.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	instance.Errorln(args...)
}

func Fatal(args ...interface{}) {
	instance.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	instance.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	instance.Fatalln(args...)
}

func Info(args ...interface{}) {
	instance.Info(args...)
}

func Infof(format string, args ...interface{}) {
	instance.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	instance.Infoln(args...)
}

func Trace(args ...interface{}) {
	instance.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	instance.Tracef(format, args...)
}

func Traceln(args ...interface{}) {
	instance.Traceln(args...)
}

func Panic(args ...interface{}) {
	instance.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	instance.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	instance.Panicln(args...)
}

func Print(args ...interface{}) {
	instance.Print(args...)
}

func Printf(format string, args ...interface{}) {
	instance.Printf(format, args...)
}

func Println(args ...interface{}) {
	instance.Println(args...)
}

func Warn(args ...interface{}) {
	instance.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	instance.Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	instance.Warnln(args...)
}

func WithField(key string, value interface{}) LogEntry {
	return instance.WithField(key, value)
}

func WithFields(fields Fields) LogEntry {
	return instance.WithFields(fields)
}

func WithContext(ctx context.Context) LogEntry {
	return instance.WithContext(ctx)
}

func AddHook(hook interface{}) error {
	if logrusHook, ok := hook.(logrus.Hook); ok {
		instance.AddHook(logrusHook)
		return nil
	}
	return fmt.Errorf("Unsupported hook type attached")
}

func NewEntry() LogEntry {
	return instance.NewEntry()
}

func SetWriter(writer io.Writer) {
	instance.SetWriter(writer)
}

func SetReportCaller(reportCaller bool) {
	instance.SetReportCaller(reportCaller)
}

func SetLevel(level string) {
	instance.SetLevel(level)
}

func SetFormatter(formatter string) {
	instance.SetFormatter(formatter)
}
