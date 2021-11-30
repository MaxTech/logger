package logger

import (
	"fmt"
	"io"
	"log"
)

type appLogger struct {
	loggerName string
	stdLogger     *log.Logger
	errLogger     *log.Logger
}

func NewLogger(_loggerName string, _stdWriter, _errWriter io.Writer) *appLogger {
	return &appLogger{
		loggerName: _loggerName,
		stdLogger: log.New(_stdWriter, "", log.LstdFlags),
		errLogger: log.New(_errWriter, "", log.LstdFlags),
	}
}

func (al *appLogger) ChangeWriter(_stdWriter, _errWriter io.Writer) {
	al.stdLogger.SetOutput(_stdWriter)
	al.errLogger.SetOutput(_errWriter)
}

func (al appLogger) Error(v ...interface{}) {
	v = append([]interface{}{
		"[ERROR]",
	}, v...)
	al.error(v...)
}

func (al appLogger) Fatal(v ...interface{}) {
	v = append([]interface{}{
		"[FATAL]",
	}, v...)
	al.fatal(v...)
}

func (al appLogger) Panic(v ...interface{}) {
	v = append([]interface{}{
		"[PANIC]",
	}, v...)
	al.panic(v...)
}

func (al appLogger) Warn(v ...interface{}) {
	v = append([]interface{}{
		"[ WARN]",
	}, v...)
	al.output(v...)
}

func (al appLogger) Info(v ...interface{}) {
	v = append([]interface{}{
		"[ INFO]",
	}, v...)
	al.output(v...)
}

func (al appLogger) output(v ...interface{}) {
	if len(al.loggerName) > 0 {
		v = append([]interface{}{
			fmt.Sprintf("[%s]", al.loggerName),
		}, v...)
	}
	al.stdLogger.Println(v...)
}

func (al appLogger) error(v ...interface{}) {
	if len(al.loggerName) > 0 {
		v = append([]interface{}{
			fmt.Sprintf("[%s]", al.loggerName),
		}, v...)
	}
	al.errLogger.Println(v...)
}

func (al appLogger) fatal(v ...interface{}) {
	if len(al.loggerName) > 0 {
		v = append([]interface{}{
			fmt.Sprintf("[%s]", al.loggerName),
		}, v...)
	}
	al.errLogger.Fatalln(v...)
}

func (al appLogger) panic(v ...interface{}) {
	if len(al.loggerName) > 0 {
		v = append([]interface{}{
			fmt.Sprintf("[%s]", al.loggerName),
		}, v...)
	}
	al.errLogger.Panicln(v...)
}
