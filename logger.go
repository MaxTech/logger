package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type appLogger struct {
	loggerName string
	logger     *log.Logger
}

func NewLogger(_loggerName string, _writer io.Writer) *appLogger {
	return &appLogger{loggerName: _loggerName, logger: log.New(_writer, "", log.LstdFlags)}
}

func (al *appLogger) ChangeWriter(_writer io.Writer) {
	al.logger.SetOutput(_writer)
}

func (al appLogger) Error(v ...interface{}) {
	v = append([]interface{}{
		"[ERROR]",
	}, v...)
	al.output(v...)
}

func (al appLogger) Fatal(v ...interface{}) {
	al.Error(v...)
	os.Exit(2)
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
	al.logger.Println(v...)
}
