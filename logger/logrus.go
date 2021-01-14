package logger

import (
	"context"
	"fmt"
	"io"
	"runtime"

	"github.com/sirupsen/logrus"
)

var (
	logrusLevels = map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
		"trace": logrus.TraceLevel,
	}

	logrusFormatters = map[string]logrus.Formatter{
		"text": &logrus.TextFormatter{FullTimestamp: true,
			CallerPrettyfier: reportCallerFilenameWithLineNumber},
		"json": &logrus.JSONFormatter{CallerPrettyfier: reportCallerFilenameWithLineNumber},
	}
)

func newLogrusLogger() Logger {
	l := logrus.New()
	l.SetLevel(logrusLevels["debug"])
	l.SetFormatter(logrusFormatters["json"])
	return &logrusLogger{logger: l}
}

type logrusLogEntry struct {
	entry *logrus.Entry
}

func (e *logrusLogEntry) Debug(args ...interface{}) {
	e.entry.Debug(args...)
}

func (e *logrusLogEntry) Debugf(format string, args ...interface{}) {
	e.entry.Debugf(format, args...)
}

func (e *logrusLogEntry) Debugln(args ...interface{}) {
	e.entry.Debugln(args...)
}

func (e *logrusLogEntry) Error(args ...interface{}) {
	e.entry.Error(args...)
}

func (e *logrusLogEntry) Errorf(format string, args ...interface{}) {
	e.entry.Errorf(format, args...)
}

func (e *logrusLogEntry) Errorln(args ...interface{}) {
	e.entry.Errorln(args...)
}

func (e *logrusLogEntry) Fatal(args ...interface{}) {
	e.entry.Fatal(args...)
}

func (e *logrusLogEntry) Fatalf(format string, args ...interface{}) {
	e.entry.Fatalf(format, args...)
}

func (e *logrusLogEntry) Fatalln(args ...interface{}) {
	e.entry.Fatalln(args...)
}

func (e *logrusLogEntry) Info(args ...interface{}) {
	e.entry.Info(args...)
}

func (e *logrusLogEntry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

func (e *logrusLogEntry) Infoln(args ...interface{}) {
	e.entry.Infoln(args...)
}

func (e *logrusLogEntry) Trace(args ...interface{}) {
	e.entry.Trace(args...)
}

func (e *logrusLogEntry) Tracef(format string, args ...interface{}) {
	e.entry.Tracef(format, args...)
}

func (e *logrusLogEntry) Traceln(args ...interface{}) {
	e.entry.Traceln(args...)
}

func (e *logrusLogEntry) Panic(args ...interface{}) {
	e.entry.Panic(args...)
}

func (e *logrusLogEntry) Panicf(format string, args ...interface{}) {
	e.entry.Panicf(format, args...)
}

func (e *logrusLogEntry) Panicln(args ...interface{}) {
	e.entry.Panicln(args...)
}

func (e *logrusLogEntry) Print(args ...interface{}) {
	e.entry.Print(args...)
}

func (e *logrusLogEntry) Printf(format string, args ...interface{}) {
	e.entry.Printf(format, args...)
}

func (e *logrusLogEntry) Println(args ...interface{}) {
	e.entry.Println(args...)
}

func (e *logrusLogEntry) Warn(args ...interface{}) {
	e.entry.Warn(args...)
}

func (e *logrusLogEntry) Warnf(format string, args ...interface{}) {
	e.entry.Warnf(format, args...)
}

func (e *logrusLogEntry) Warnln(args ...interface{}) {
	e.entry.Warnln(args...)
}

func (e *logrusLogEntry) WithField(key string, value interface{}) LogEntry {
	return &logrusLogEntry{entry: e.entry.WithField(key, value)}
}

func (e *logrusLogEntry) WithFields(fields Fields) LogEntry {
	return &logrusLogEntry{entry: e.entry.WithFields(logrus.Fields(fields))}
}

func (e *logrusLogEntry) WithContext(ctx context.Context) LogEntry {
	return &logrusLogEntry{entry: e.entry.WithContext(ctx)}
}

type logrusLogger struct {
	logger *logrus.Logger
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

func (l *logrusLogger) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *logrusLogger) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

func (l *logrusLogger) Traceln(args ...interface{}) {
	l.logger.Traceln(args...)
}

func (l *logrusLogger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

func (l *logrusLogger) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}

func (l *logrusLogger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

func (l *logrusLogger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *logrusLogger) Println(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

func (l *logrusLogger) WithField(key string, value interface{}) LogEntry {
	return l.NewEntry().WithField(key, value)
}

func (l *logrusLogger) WithFields(fields Fields) LogEntry {
	return l.NewEntry().WithFields(fields)
}

func (l *logrusLogger) WithContext(ctx context.Context) LogEntry {
	return l.NewEntry().WithContext(ctx)
}

func (l *logrusLogger) AddHook(hook interface{}) error {
	if logrusHook, ok := hook.(logrus.Hook); ok {
		l.logger.AddHook(logrusHook)
		return nil
	}
	return fmt.Errorf("Unsupported hook type attached")
}

func (l *logrusLogger) NewEntry() LogEntry {
	return &logrusLogEntry{entry: logrus.NewEntry(l.logger)}
}

func (l *logrusLogger) SetWriter(writer io.Writer) {
	l.logger.SetOutput(writer)
}

func (l *logrusLogger) SetReportCaller(reportCaller bool) {
	l.logger.SetReportCaller(reportCaller)
}

func (l *logrusLogger) SetLevel(level string) {
	if logrusLevel, ok := logrusLevels[level]; ok {
		l.logger.SetLevel(logrusLevel)
	} else {
		panic(fmt.Sprintf("Unsupported level '%s'", level))
	}
}

func (l *logrusLogger) SetFormatter(formatter string) {
	if logrusFormatter, ok := logrusFormatters[formatter]; ok {
		l.logger.SetFormatter(logrusFormatter)
	} else {
		panic(fmt.Sprintf("Unsupported formatter '%s'", logrusFormatter))
	}
}

func reportCallerFilenameWithLineNumber(f *runtime.Frame) (string, string) {
	// Need to iterate through the runtime and report the actual caller
	// Leaving this with bug temporarily
	return "", fmt.Sprintf("%s:%d", f.File, f.Line)
}
