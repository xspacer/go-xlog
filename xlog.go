package xlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Log struct {
	*zap.Logger
}

type Config struct {
	// Logging priority. Higher levels are more important.
	Level zapcore.Level

	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" yaml:"filename"`

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted.)
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	// Compress determines if the rotated log files should be compressed
	// using gzip. The default is not to perform compression.
	Compress bool `json:"compress" yaml:"compress"`
}

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

func Level(level string) zapcore.Level {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	default:
		return InfoLevel
	}
}

func New(c *Config) *Log {
	var w zapcore.WriteSyncer
	if c.Filename == "" {
		w = os.Stdout
	} else {
		w = zapcore.AddSync(&lumberjack.Logger{
			Filename:   c.Filename,
			MaxSize:    c.MaxSize,
			MaxAge:     c.MaxAge,
			MaxBackups: c.MaxBackups,
			LocalTime:  c.LocalTime,
			Compress:   c.Compress,
		})
	}
	e := zap.NewProductionEncoderConfig()
	e.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(e),
		w,
		c.Level,
	)

	return &Log{zap.New(core, zap.AddCaller())}
}

var DefaultLog *Log
var DefaultSugaredLogger *zap.SugaredLogger

func init() {
	DefaultLog = &Log{New(&Config{Level: InfoLevel}).WithOptions(zap.AddCallerSkip(1))}
	DefaultSugaredLogger = DefaultLog.Sugar()
}

func Init(c *Config) *Log {
	DefaultLog = &Log{New(c).WithOptions(zap.AddCallerSkip(1))}
	DefaultSugaredLogger = DefaultLog.Sugar()
	return DefaultLog
}

func Debug(args ...interface{}) {
	DefaultSugaredLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	DefaultSugaredLogger.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Debugw(msg, keysAndValues...)
}

func Debugz(msg string, fields ...zap.Field) {
	DefaultLog.Debug(msg, fields...)
}

func Info(args ...interface{}) {
	DefaultSugaredLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	DefaultSugaredLogger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Infow(msg, keysAndValues...)
}

func Infoz(msg string, fields ...zap.Field) {
	DefaultLog.Info(msg, fields...)
}

func Warn(args ...interface{}) {
	DefaultSugaredLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	DefaultSugaredLogger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Warnw(msg, keysAndValues...)
}

func Warnz(msg string, fields ...zap.Field) {
	DefaultLog.Warn(msg, fields...)
}

func Error(args ...interface{}) {
	DefaultSugaredLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	DefaultSugaredLogger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Errorw(msg, keysAndValues...)
}

func Errorz(msg string, fields ...zap.Field) {
	DefaultLog.Error(msg, fields...)
}

func Panic(args ...interface{}) {
	DefaultSugaredLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	DefaultSugaredLogger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Panicw(msg, keysAndValues...)
}

func Panicz(msg string, fields ...zap.Field) {
	DefaultLog.Panic(msg, fields...)
}

func Fatal(args ...interface{}) {
	DefaultSugaredLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	DefaultSugaredLogger.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	DefaultSugaredLogger.Fatalw(msg, keysAndValues...)
}

func Fatalz(msg string, fields ...zap.Field) {
	DefaultLog.Fatal(msg, fields...)
}

func WithOptions(opts ...zap.Option) *Log {
	return &Log{DefaultLog.Logger.WithOptions(opts...)}
}
