package logger

import (
	"fmt"
	"io"
	"log"
)

type appLogger struct {
	loggerName string
	stdLogger  *log.Logger
	errLogger  *log.Logger
}

func NewLogger(_loggerName string, _stdWriter, _errWriter io.Writer) *appLogger {
	return &appLogger{
		loggerName: _loggerName,
		stdLogger:  log.New(_stdWriter, "", log.LstdFlags),
		errLogger:  log.New(_errWriter, "", log.LstdFlags),
	}
}

func (al *appLogger) ChangeWriter(_stdWriter, _errWriter io.Writer) {
	al.stdLogger.SetOutput(_stdWriter)
	al.errLogger.SetOutput(_errWriter)
}

func (al appLogger) compose(_flag string) []interface{} {
	v := []interface{}{
		fmt.Sprintf("[%s]", _flag),
	}
	if len(al.loggerName) > 0 {
		v = append(v, fmt.Sprintf("[%s]", al.loggerName))
	}
	return v
}

func (al appLogger) Error(v ...interface{}) {
	v = append(al.compose("ERROR"), v...)
	al.errLogger.Println(v...)
}

func (al appLogger) Fatal(v ...interface{}) {
	v = append(al.compose("FATAL"), v...)
	al.errLogger.Fatalln(v...)
}

func (al appLogger) Panic(v ...interface{}) {
	v = append(al.compose("PANIC"), v...)
	al.errLogger.Panicln(v...)
}

func (al appLogger) Warn(v ...interface{}) {
	v = append(al.compose(" WARN"), v...)
	al.stdLogger.Println(v...)
}

func (al appLogger) Info(v ...interface{}) {
	v = append(al.compose(" INFO"), v...)
	al.stdLogger.Println(v...)
}
