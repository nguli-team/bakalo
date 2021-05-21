package logger

var Log Logger

// Logger represent common interface for logging function
type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
	Printf(format string, args ...interface{})
	Print(args ...interface{})
}

func SetLogger(logger Logger) {
	Log = logger
}
