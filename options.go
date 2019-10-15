package xlog

import "go.uber.org/zap/zapcore"

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

type options struct {
	// logging priority. Higher levels are more important.
	level zapcore.Level

	// name of log file
	filename string

	// maxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	maxSize int

	// maxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	maxBackups int

	// maxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	maxAge int

	// localTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	localTime bool

	// compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	compress bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func Level(level zapcore.Level) Option {
	return optionFunc(func(o *options) {
		o.level = level
	})
}

func StringLevel(level string) Option {
	return optionFunc(func(o *options) {
		switch level {
		case "debug":
			o.level = DebugLevel
		case "info":
			o.level = InfoLevel
		case "warn":
			o.level = WarnLevel
		case "error":
			o.level = ErrorLevel
		case "panic":
			o.level = PanicLevel
		case "fatal":
			o.level = FatalLevel
		default:
			o.level = InfoLevel
		}
	})
}

func Filename(filename string) Option {
	return optionFunc(func(o *options) {
		o.filename = filename
	})
}

func MaxSize(maxSize int) Option {
	return optionFunc(func(o *options) {
		o.maxSize = maxSize
	})
}

func MaxBackups(maxBackups int) Option {
	return optionFunc(func(o *options) {
		o.maxBackups = maxBackups
	})
}

func MaxAge(maxAge int) Option {
	return optionFunc(func(o *options) {
		o.maxAge = maxAge
	})
}

func LocalTime(localTime bool) Option {
	return optionFunc(func(o *options) {
		o.localTime = localTime
	})
}

func Compress(compress bool) Option {
	return optionFunc(func(o *options) {
		o.compress = compress
	})
}
