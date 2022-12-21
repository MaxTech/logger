package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

type appLogger struct {
	level  LogLevel
	logger *log.Logger
}

// SetLogFlags
// default: log.LstdFlags
// example 1: log.LstdFlags | log.Lmicroseconds | log.Llongfile
// example 2: log.LstdFlags | log.Lmicroseconds | log.Lshortfile
// example 3: log.LstdFlags | log.Lmicrosecondsw
func (al *appLogger) SetLogFlags(_flagCode int) {
	al.logger.SetFlags(_flagCode)
}

func (al *appLogger) SetLogLevel(_level LogLevel) {
	if _level.LowerThan(LogLevelDebug) {
		_level = LogLevelDebug
	}
	if _level.HigherThan(LogLevelFatal) {
		_level = LogLevelFatal
	}
	al.level = _level
}

func (al *appLogger) SetWriter(_writer io.Writer) {
	al.logger.SetOutput(_writer)
}

func (al *appLogger) SetName(_name string) {
	if len(_name) > 0 {
		_name = fmt.Sprintf("[%s] ", _name)
	}
	al.logger.SetPrefix(_name)
}

func (al *appLogger) Error(v ...interface{}) {
	if al.level.HigherThan(LogLevelError) {
		return
	}
	headList := []interface{}{
		"[ERROR]",
	}
	if _, fileName, lineNum, ok := runtime.Caller(1); ok {
		headList = append(headList, fmt.Sprintf("%s:%d", fileName, lineNum))
	}
	v = append(headList, v...)
	err := al.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "appLogger Error Output error:", err)
	}
}

func (al *appLogger) Fatal(v ...interface{}) {
	if al.level.HigherThan(LogLevelFatal) {
		return
	}
	headList := []interface{}{
		"[FATAL]",
	}
	if _, fileName, lineNum, ok := runtime.Caller(1); ok {
		headList = append(headList, fmt.Sprintf("%s:%d", fileName, lineNum))
	}
	v = append(headList, v...)
	err := al.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "appLogger Fatal Output error:", err)
	}
	os.Exit(1)
}

func (al *appLogger) Warn(v ...interface{}) {
	if al.level.HigherThan(LogLevelWarn) {
		return
	}
	headList := []interface{}{
		"[ WARN]",
	}
	v = append(headList, v...)
	err := al.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "appLogger Warn Output error:", err)
	}
}

func (al *appLogger) Info(v ...interface{}) {
	if al.level.HigherThan(LogLevelInfo) {
		return
	}
	headList := []interface{}{
		"[ INFO]",
	}
	v = append(headList, v...)
	err := al.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "appLogger Info Output error:", err)
	}
}

func (al *appLogger) Debug(v ...interface{}) {
	if al.level.HigherThan(LogLevelDebug) {
		return
	}
	headList := []interface{}{
		"[DEBUG]",
	}
	if _, fileName, lineNum, ok := runtime.Caller(1); ok {
		headList = append(headList, fmt.Sprintf("%s:%d", fileName, lineNum))
	}
	v = append(headList, v...)
	err := al.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "appLogger Debug Output error:", err)
	}
}

func NewDefaultAppLogger() *appLogger {
	return &appLogger{logger: log.New(os.Stdout, "", log.LstdFlags), level: LogLevelInfo}
}

func NewAppLoggerWithName(_name string) *appLogger {
	if len(_name) > 0 {
		_name = fmt.Sprintf("[%s] ", _name)
	}
	return &appLogger{logger: log.New(os.Stdout, _name, log.LstdFlags), level: LogLevelInfo}
}
