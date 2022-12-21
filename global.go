package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

var (
	stdLogger *appLogger
	errLogger *appLogger
)

func init() {
	stdLogger = &appLogger{logger: log.New(os.Stdout, "", log.LstdFlags), level: LogLevelInfo}
	errLogger = &appLogger{logger: log.New(os.Stderr, "", log.LstdFlags), level: LogLevelInfo}
}

func ResetName(_name string) {
	stdLogger = &appLogger{logger: log.New(os.Stdout, _name, log.LstdFlags), level: LogLevelInfo}
	errLogger = &appLogger{logger: log.New(os.Stderr, _name, log.LstdFlags), level: LogLevelInfo}
}

// SetLogFlags
// default: log.LstdFlags
// example 1: log.LstdFlags | log.Lmicroseconds | log.Llongfile
// example 2: log.LstdFlags | log.Lmicroseconds | log.Lshortfile
// example 3: log.LstdFlags | log.Lmicrosecondsw
func SetLogFlags(_flagCode int) {
	stdLogger.SetLogFlags(_flagCode)
	errLogger.SetLogFlags(_flagCode)
}

func SetLogLevel(_level LogLevel) {
	stdLogger.SetLogLevel(_level)
	if stdLogger.level == LogLevelDebug {
		stdLogger.logger.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
	}
	errLogger.SetLogLevel(_level)
	if errLogger.level == LogLevelDebug {
		errLogger.logger.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
	}
}

func SetLogWriter(_stdWriter, _errWriter io.Writer) {
	stdLogger.SetWriter(_stdWriter)
	errLogger.SetWriter(_errWriter)
}

func Error(v ...interface{}) {
	if errLogger.level.HigherThan(LogLevelError) {
		return
	}
	headList := []interface{}{
		"[ERROR]",
	}
	if _, fileName, lineNum, ok := runtime.Caller(1); ok {
		headList = append(headList, fmt.Sprintf("%s:%d", fileName, lineNum))
	}
	v = append(headList, v...)
	err := errLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Error Output error:", err)
	}
}

func Fatal(v ...interface{}) {
	if errLogger.level.HigherThan(LogLevelFatal) {
		return
	}
	headList := []interface{}{
		"[FATAL]",
	}
	if _, fileName, lineNum, ok := runtime.Caller(1); ok {
		headList = append(headList, fmt.Sprintf("%s:%d", fileName, lineNum))
	}
	v = append(headList, v...)
	err := errLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Fatal Output error:", err)
	}
	os.Exit(1)
}

func Warn(v ...interface{}) {
	if stdLogger.level.HigherThan(LogLevelWarn) {
		return
	}
	headList := []interface{}{
		"[ WARN]",
	}
	v = append(headList, v...)
	err := stdLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Warn Output error:", err)
	}
}

func Info(v ...interface{}) {
	if stdLogger.level.HigherThan(LogLevelInfo) {
		return
	}
	headList := []interface{}{
		"[ INFO]",
	}
	v = append(headList, v...)
	err := stdLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Info Output error:", err)
	}
}

func Debug(v ...interface{}) {
	if stdLogger.level.HigherThan(LogLevelDebug) {
		return
	}
	headList := []interface{}{
		"[DEBUG]",
	}
	if _, fileName, lineNum, ok := runtime.Caller(1); ok {
		headList = append(headList, fmt.Sprintf("%s:%d", fileName, lineNum))
	}
	v = append(headList, v...)
	err := stdLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Debug Output error:", err)
	}
}
