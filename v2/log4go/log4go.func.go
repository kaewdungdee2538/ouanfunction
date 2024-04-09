package log4go

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alecthomas/log4go"
	"github.com/kaewdungdee2538/ouanfunction/directoryfunc"
	"github.com/kaewdungdee2538/ouanfunction/formatString"
	"github.com/kaewdungdee2538/ouanfunction/json"
)

// ----------------- verson2
func (l Log4goLogger) WriteLog(model Log4goModel, logMode string) {
	switch logMode {
	case "warning":
		go l.writeNewLogWarning(model)
	case "error":
		go l.writeNewLogError(model)
	case "critical":
		go l.writeNewLogCritical(model)
	default:
		go l.writeNewLogInfo(model)
	}
}

// info log
func (l Log4goLogger) writeNewLogInfo(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.INFO)
	defer log.Close()
	// check directory
	currentDirectory := model.Driectory
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := l.generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.INFO, flw)
	logInfo := l.newCreateLogFormat(model)
	log.Info(logInfo)
}

// warning log
func (l Log4goLogger) writeNewLogWarning(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.WARNING)
	defer log.Close()
	// check directory
	currentDirectory := model.Driectory
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := l.generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.WARNING, flw)
	logInfo := l.newCreateLogFormat(model)
	log.Warn(logInfo)
}

// error log
func (l Log4goLogger) writeNewLogError(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.ERROR)
	defer log.Close()
	// check directory
	currentDirectory := model.Driectory
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := l.generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.ERROR, flw)
	logInfo := l.newCreateLogFormat(model)
	log.Error(logInfo)
}

// critical log
func (l Log4goLogger) writeNewLogCritical(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.CRITICAL)
	defer log.Close()
	// check directory
	currentDirectory := model.Driectory
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := l.generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.CRITICAL, flw)
	logInfo := l.newCreateLogFormat(model)
	log.Critical(logInfo)
}

// for
func (l Log4goLogger) newCreateLogFormat(model Log4goModel) string {
	dataJsonText, _ := json.PrettyStructJson(model.Data)

	logInfo := fmt.Sprintf(`
	APP_NAME=> %s
	FUNCTION_NAME => [%s] 
	REMOTE_ADDRESS => %s
	REMOTE_PORT => %s 
	ORIGINALURL => %s
	NODE_ID => %s
	MESSAGE => [%s]
	DATA => %v
	`,
		model.AppName, model.FunctionName, model.ClientIP, model.RequestID, model.OriginPath, model.NodeId, model.Msg, dataJsonText,
	)
	return logInfo
}

// get current date
func (l Log4goLogger) getCurrentDate(appName string) string {
	appname := formatString.StringToCamelCase(appName)
	now := time.Now()
	currentdate_str := fmt.Sprintf("%s.%s", appname, now.Format("2006-01-02"))
	return currentdate_str
}

// get older date
func (l Log4goLogger) getOlderDate(appName string) string {
	maxDays := 30
	if l.MaxDaysBackup > 0 {
		maxDays = l.MaxDaysBackup
	}
	appname := formatString.StringToCamelCase(appName)
	now := time.Now().Add(time.Duration(-maxDays) * (time.Hour * 24))
	currentdate_str := fmt.Sprintf("%s.%s", appname, now.Format("2006-01-02"))
	return currentdate_str
}

// generate format
func (l Log4goLogger) generateFormatLog(currentDirectory string, appName string) *log4go.FileLogWriter {
	// delete old log
	go l.deleteOldLogFile(currentDirectory, appName)
	logFileName := fmt.Sprintf(`%s/%s.log`, currentDirectory, l.getCurrentDate(appName))
	flw := log4go.NewFileLogWriter(logFileName, false)
	flw.SetFormat("[%D %T] [%L], %M")
	return flw
}

// delete older log file
func (l Log4goLogger) deleteOldLogFile(currentDirectory string, appName string) *log4go.FileLogWriter {
	// get old date
	l.getOlderDate(appName)
	logFileName := fmt.Sprintf(`%s/%s.log`, currentDirectory, l.getOlderDate(appName))
	if _, err := os.Stat(logFileName); err == nil {
		e := os.Remove(logFileName)
		if e != nil {
			log.Println(e)
		}
		log.Println("delete older log file : ", logFileName)
	}
	return nil
}
