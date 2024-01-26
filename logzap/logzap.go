package logzap

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

func init() {

	// config := zap.NewProductionConfig()

	// config.EncoderConfig.TimeKey = "timestamp"
	// config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// config.EncoderConfig.StacktraceKey = "" // for close stacktrace

	// Create a file syncer
	lumberjackLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/" + getCurrentDate() + ".log", // Adjust the filename format as needed
		MaxSize:    10,                                  // Max megabytes before log is rolled
		MaxBackups: 5,                                   // Max number of old log files to retain
		MaxAge:     30,                                  // Max days to retain old log files
		Compress:   true,                                // Whether to compress rolled files
	})

	// var err error

	// Create a Zap encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(time.RFC3339Nano))
	}
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.StacktraceKey = ""
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	encoder := zapcore.NewJSONEncoder(encoderConfig)

	

	// Create a Zap core with the Lumberjack logger
	core := zapcore.NewCore(encoder, zapcore.AddSync(lumberjackLogger), zap.InfoLevel)

	// Create a Zap logger
	// log = zap.New(core,zap.AddCaller(), zap.AddCallerSkip(1))
	log = zap.New(core)
	// flushes buffer, if any
	defer log.Sync()
}

func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func Info(msg string, req interface{},res interface{}) {
	request := zap.Any("request", req)
	response := zap.Any("response", res)
	log.Info(msg, request,response)
}

func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}

func Error(msg interface{}, fields ...zapcore.Field) {

	switch t := msg.(type) {
	case error:
		log.Error(t.Error(), fields...)
	case string:
		log.Error(t, fields...)
	}
}
