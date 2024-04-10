package logzap

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func getConfig() *zap.Logger {

	path := "/log_app/"
	// Create a file syncer
	lumberjackLogger := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", path, getCurrentDate()), // Adjust the filename format as needed
		MaxSize:    10,                                               // Max megabytes before log is rolled
		MaxBackups: 5,                                                // Max number of old log files to retain
		MaxAge:     30,                                               // Max days to retain old log files
		Compress:   true,                                             // Whether to compress rolled files
	})

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
	// log = zap.New(core)
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	// flushes buffer, if any
	defer log.Sync()
	return log
}

func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func Info(msg string, req interface{}, res interface{}) {
	log := getConfig()
	request := zap.Any("request", req)
	response := zap.Any("response", res)
	log.Info(msg, request, response)
}

func Debug(msg string, req interface{}, res interface{}) {
	log := getConfig()
	request := zap.Any("request", req)
	response := zap.Any("response", res)
	log.Debug(msg, request, response)
}

func Error(msg interface{}, req interface{}, res interface{}) {
	log := getConfig()
	request := zap.Any("request", req)
	response := zap.Any("response", res)

	switch t := msg.(type) {
	case error:
		log.Error(t.Error(), request, response)
	case string:
		log.Error(t, request, response)
	}
}
