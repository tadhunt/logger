package logger

type CompatLogWriter interface {
	SetLevel(level LogLevel)
	Syslog(bool) error
	Level() LogLevel
	SetPrefix(prefix string)
	Packetf(format string, args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(f string, args ...any)
	Fatalf(format string, args ...any)
	ErrFmt(format string, args ...any) error
	Prefix() string
	Write(data []byte) (int, error)
	Json(label string, data any, indent bool)
	Flush()
	SetId(id int)
	Id() int
}
