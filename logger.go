package logger

import (
	"io"
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = 1
	LogLevelInfo  LogLevel = 2
	LogLevelWarn  LogLevel = 3
	LogLevelError LogLevel = 4
	LogLevelFatal LogLevel = 5
)

type appLogger struct {
	logger *log.Logger
}

var (
	stdLogger *appLogger
	errLogger *appLogger
	logLevel  int
)

func init() {
	stdLogger = &appLogger{logger: log.New(os.Stdout, "", log.LstdFlags)}
	errLogger = &appLogger{logger: log.New(os.Stderr, "", log.LstdFlags)}
	logLevel = 2
}

func SetLogLevel(_level LogLevel) {
	if int(_level) < int(LogLevelDebug) {
		_level = LogLevelDebug
	}
	if int(_level) > int(LogLevelFatal) {
		_level = LogLevelFatal
	}
	logLevel = int(_level)
}

func InitLoggerByWriter(_stdWriter, _errWriter io.Writer) {
	stdLogger = &appLogger{logger: log.New(_stdWriter, "", log.LstdFlags)}
	errLogger = &appLogger{logger: log.New(_errWriter, "", log.LstdFlags)}
}

func Error(v ...interface{}) {
	if logLevel > int(LogLevelError) {
		return
	}
	v = append([]interface{}{
		"[ERROR]",
	}, v...)
	errLogger.logger.Println(v...)
}

func Fatal(v ...interface{}) {
	if logLevel > int(LogLevelFatal) {
		return
	}
	Error(v...)
	os.Exit(2)
}

func Warn(v ...interface{}) {
	if logLevel > int(LogLevelWarn) {
		return
	}
	v = append([]interface{}{
		"[ WARN]",
	}, v...)
	errLogger.logger.Println(v...)
}

func Info(v ...interface{}) {
	if logLevel > int(LogLevelInfo) {
		return
	}
	v = append([]interface{}{
		"[ INFO]",
	}, v...)
	stdLogger.logger.Println(v...)
}

func Debug(v ...interface{}) {
	if logLevel > int(LogLevelDebug) {
		return
	}
	v = append([]interface{}{
		"[DEBUG]",
	}, v...)
	stdLogger.logger.Println(v...)
}
