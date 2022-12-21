package logger

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
