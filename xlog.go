package xlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	DefaultLog           *Log
	DefaultSugaredLogger *zap.SugaredLogger
	_defaultOptions      *options
)

func init() {
	_defaultOptions = &options{
		level:      InfoLevel,
		filename:   "",
		maxSize:    100,
		maxBackups: 0,
		maxAge:     0,
		localTime:  true,
		compress:   true,
	}
	DefaultLog = &Log{New().WithOptions(zap.AddCallerSkip(1))}
	DefaultSugaredLogger = DefaultLog.Sugar()
}

func Init(opts ...Option) *Log {
	DefaultLog = &Log{New(opts...).WithOptions(zap.AddCallerSkip(1))}
	DefaultSugaredLogger = DefaultLog.Sugar()
	return DefaultLog
}

type Log struct {
	*zap.Logger
}

func New(opts ...Option) *Log {
	options := *_defaultOptions
	for _, o := range opts {
		o.apply(&options)
	}

	var w zapcore.WriteSyncer = os.Stdout
	if options.filename != "" {
		w = zapcore.AddSync(&lumberjack.Logger{
			Filename:   options.filename,
			MaxSize:    options.maxSize,
			MaxBackups: options.maxBackups,
			MaxAge:     options.maxAge,
			LocalTime:  options.localTime,
			Compress:   options.compress,
		})
	}
	e := zap.NewProductionEncoderConfig()
	e.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(e),
		w,
		options.level,
	)

	return &Log{zap.New(core, zap.AddCaller())}
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

func WithZapOptions(opts ...zap.Option) *Log {
	return &Log{DefaultLog.Logger.WithOptions(opts...)}
}
