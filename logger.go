package logger

import (
	"fmt"
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

func (ll LogLevel) LowerThan(_level LogLevel) bool {
	return int(ll) < int(_level)
}

func (ll LogLevel) HigherThan(_level LogLevel) bool {
	return int(ll) > int(_level)
}

type appLogger struct {
	level  LogLevel
	logger *log.Logger
}

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
	v = append([]interface{}{
		"[ERROR]",
	}, v...)
	err := al.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "appLogger Error Output error:", err)
	}
}

func (al *appLogger) Fatal(v ...interface{}) {
	if al.level.HigherThan(LogLevelFatal) {
		return
	}
	v = append([]interface{}{
		"[FATAL]",
	}, v...)
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
	v = append([]interface{}{
		"[ WARN]",
	}, v...)
	err := al.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "appLogger Warn Output error:", err)
	}
}

func (al *appLogger) Info(v ...interface{}) {
	if al.level.HigherThan(LogLevelInfo) {
		return
	}
	v = append([]interface{}{
		"[ INFO]",
	}, v...)
	err := al.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "appLogger Info Output error:", err)
	}
}

func (al *appLogger) Debug(v ...interface{}) {
	if al.level.HigherThan(LogLevelDebug) {
		return
	}
	v = append([]interface{}{
		"[DEBUG]",
	}, v...)
	err := al.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "appLogger Debug Output error:", err)
	}
}

var (
	stdLogger *appLogger
	errLogger *appLogger
)

func init() {
	stdLogger = &appLogger{logger: log.New(os.Stdout, "", log.LstdFlags), level: LogLevelInfo}
	errLogger = &appLogger{logger: log.New(os.Stderr, "", log.LstdFlags), level: LogLevelInfo}
}

func NewDefaultLogger() *appLogger {
	return &appLogger{logger: log.New(os.Stdout, "", log.LstdFlags), level: LogLevelInfo}
}

func NewCustomLoggerWithName(_name string) *appLogger {
	if len(_name) > 0 {
		_name = fmt.Sprintf("[%s] ", _name)
	}
	return &appLogger{logger: log.New(os.Stdout, _name, log.LstdFlags), level: LogLevelInfo}
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
	v = append([]interface{}{
		"[ERROR]",
	}, v...)
	err := errLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Error Output error:", err)
	}
}

func Fatal(v ...interface{}) {
	if errLogger.level.HigherThan(LogLevelFatal) {
		return
	}
	v = append([]interface{}{
		"[FATAL]",
	}, v...)
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
	v = append([]interface{}{
		"[ WARN]",
	}, v...)
	err := stdLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Warn Output error:", err)
	}
}

func Info(v ...interface{}) {
	if stdLogger.level.HigherThan(LogLevelInfo) {
		return
	}
	v = append([]interface{}{
		"[ INFO]",
	}, v...)
	err := stdLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Info Output error:", err)
	}
}

func Debug(v ...interface{}) {
	if stdLogger.level.HigherThan(LogLevelDebug) {
		return
	}
	v = append([]interface{}{
		"[DEBUG]",
	}, v...)
	err := stdLogger.logger.Output(2, fmt.Sprintln(v...))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Logger Debug Output error:", err)
	}
}
