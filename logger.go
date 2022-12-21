package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

type logger struct {
	stdLogger *appLogger
	errLogger *appLogger
}

func NewLoggerWithName(_name string) *logger {
	if len(_name) > 0 {
		_name = fmt.Sprintf("[%s] ", _name)
	}
	return &logger{
		stdLogger: &appLogger{logger: log.New(os.Stdout, _name, log.LstdFlags), level: LogLevelInfo},
		errLogger: &appLogger{logger: log.New(os.Stderr, _name, log.LstdFlags), level: LogLevelInfo},
	}
}

func (l *logger) GetStdLogger() *appLogger {
	return l.stdLogger
}

func (l *logger) GetErrLogger() *appLogger {
	return l.errLogger
}

// SetStdFlags
// default: log.LstdFlags
// example 1: log.LstdFlags | log.Lmicroseconds | log.Llongfile
// example 2: log.LstdFlags | log.Lmicroseconds | log.Lshortfile
// example 3: log.LstdFlags | log.Lmicrosecondsw
func (l *logger) SetStdFlags(_flagCode int) {
	l.stdLogger.SetLogFlags(_flagCode)
}

// SetErrFlags
// default: log.LstdFlags
// example 1: log.LstdFlags | log.Lmicroseconds | log.Llongfile
// example 2: log.LstdFlags | log.Lmicroseconds | log.Lshortfile
// example 3: log.LstdFlags | log.Lmicrosecondsw
func (l *logger) SetErrFlags(_flagCode int) {
	l.errLogger.SetLogFlags(_flagCode)
}

func (l *logger) SetStdLevel(_level LogLevel) {
	l.stdLogger.SetLogLevel(_level)
	if l.stdLogger.level == LogLevelDebug {
		l.stdLogger.logger.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
	}
}

func (l *logger) SetErrLevel(_level LogLevel) {
	l.errLogger.SetLogLevel(_level)
	if l.errLogger.level == LogLevelDebug {
		l.errLogger.logger.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
	}
}

func (l *logger) SetStdWriter(_writer io.Writer) {
	l.stdLogger.SetWriter(_writer)
}

func (l *logger) SetErrWriter(_writer io.Writer) {
	l.errLogger.SetWriter(_writer)
}

func (l *logger) Error(v ...interface{}) {
	if l.errLogger.level.HigherThan(LogLevelError) {
		return
	}
	headList := []interface{}{
		"[ERROR]",
	}
	if _, fileName, lineNum, ok := runtime.Caller(1); ok {
		headList = append(headList, fmt.Sprintf("%s:%d", fileName, lineNum))
	}
	v = append(headList, v...)
	err := l.errLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Error Output error:", err)
	}
}

func (l *logger) Fatal(v ...interface{}) {
	if l.errLogger.level.HigherThan(LogLevelFatal) {
		return
	}
	headList := []interface{}{
		"[FATAL]",
	}
	if _, fileName, lineNum, ok := runtime.Caller(1); ok {
		headList = append(headList, fmt.Sprintf("%s:%d", fileName, lineNum))
	}
	v = append(headList, v...)
	err := l.errLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Fatal Output error:", err)
	}
	os.Exit(1)
}

func (l *logger) Warn(v ...interface{}) {
	if l.stdLogger.level.HigherThan(LogLevelWarn) {
		return
	}
	headList := []interface{}{
		"[ WARN]",
	}
	v = append(headList, v...)
	err := l.stdLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Warn Output error:", err)
	}
}

func (l *logger) Info(v ...interface{}) {
	if l.stdLogger.level.HigherThan(LogLevelInfo) {
		return
	}
	headList := []interface{}{
		"[ INFO]",
	}
	v = append(headList, v...)
	err := l.stdLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Info Output error:", err)
	}
}

func (l *logger) Debug(v ...interface{}) {
	if l.stdLogger.level.HigherThan(LogLevelDebug) {
		return
	}
	headList := []interface{}{
		"[DEBUG]",
	}
	if _, fileName, lineNum, ok := runtime.Caller(1); ok {
		headList = append(headList, fmt.Sprintf("%s:%d", fileName, lineNum))
	}
	v = append(headList, v...)
	err := l.stdLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Debug Output error:", err)
	}
}
