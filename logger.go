package logger

import (
	"io"
	"log"
	"os"
)

type appLogger struct {
	logger *log.Logger
}

var (
	stdLogger *appLogger
	errLogger *appLogger
)

func init() {
	stdLogger = &appLogger{logger: log.New(os.Stdout, "", log.LstdFlags)}
	errLogger = &appLogger{logger: log.New(os.Stderr, "", log.LstdFlags)}
}

func InitLoggerByWriter(_stdWriter, _errWriter io.Writer)  {
	stdLogger = &appLogger{logger: log.New(_stdWriter, "", log.LstdFlags)}
	errLogger = &appLogger{logger: log.New(_errWriter, "", log.LstdFlags)}
}

func Error(v ...interface{}) {
	v = append([]interface{}{
		"[ERROR]",
	}, v...)
	errLogger.logger.Println(v...)
}

func Warn(v ...interface{}) {
	v = append([]interface{}{
		"[ WARN]",
	}, v...)
	errLogger.logger.Println(v...)
}

func Info(v ...interface{}) {
	v = append([]interface{}{
		"[ INFO]",
	}, v...)
	stdLogger.logger.Println(v...)
}

func Fatal(v ...interface{}) {
	Error(v...)
	os.Exit(1)
}
