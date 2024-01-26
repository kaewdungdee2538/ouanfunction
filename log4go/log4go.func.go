package log4go

import (
	"fmt"
	"net"
	"time"

	"github.com/alecthomas/log4go"
	"github.com/kaewdungdee2538/ouanfunction/directoryfunc"
	"github.com/kaewdungdee2538/ouanfunction/formatString"
	"github.com/kaewdungdee2538/ouanfunction/json"
)

// ---------------------------------------------------------------------------------------------------------------//
func WriteLogDebugAll(model Log4goModel, logMode string) {

	log := log4go.NewDefaultLogger(log4go.DEBUG)

	defer log.Close()
	// check directory
	currentDirectory := fmt.Sprintf(`%s/debug`, model.Driectory)
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.DEBUG, flw)
	logInfo := createLogFormat(model)

	switch logMode {
	case "warning":
		log.Warn(logInfo)
		go WriteLogWarning(model)
	case "error":
		log.Error(logInfo)
		go WriteLogError(model)
	case "critical":
		log.Critical(logInfo)
		go WriteLogCritical(model)
	default:
		log.Info(logInfo)
		go WriteLogInfo(model)
	}
}

// info log
// ---------------------------------------------------------------------------------------------------------------//
func WriteLogInfo(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.INFO)
	defer log.Close()
	// check directory
	currentDirectory := fmt.Sprintf(`%s/info`, model.Driectory)
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.INFO, flw)
	logInfo := createLogFormat(model)
	log.Info(logInfo)
}

// warning log
func WriteLogWarning(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.WARNING)
	defer log.Close()
	// check directory
	currentDirectory := fmt.Sprintf(`%s/warning`, model.Driectory)
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.WARNING, flw)
	logInfo := createLogFormat(model)
	log.Warn(logInfo)
}

// error log
func WriteLogError(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.ERROR)
	defer log.Close()
	// check directory
	currentDirectory := fmt.Sprintf(`%s/error`, model.Driectory)
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.ERROR, flw)
	logInfo := createLogFormat(model)
	log.Error(logInfo)
}

// critical log
func WriteLogCritical(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.CRITICAL)
	defer log.Close()
	// check directory
	currentDirectory := fmt.Sprintf(`%s/critical`, model.Driectory)
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.CRITICAL, flw)
	logInfo := createLogFormat(model)
	log.Critical(logInfo)
}

// for *gin.Context
func createLogFormat(model Log4goModel) string {
	dataJsonText, _ := json.PrettyStructJson(model.Data)

	var clientAddress = ""
	var remotePort = ""
	var fullPath = ""
	if model.Context != nil {
		if ip, port, err := net.SplitHostPort(model.Context.Request.RemoteAddr); err == nil {
			clientAddress = ip
			remotePort = port
			fullPath = model.Context.FullPath()
		} else {
			clientAddress = model.Context.ClientIP()
			fullPath = model.Context.FullPath()
		}
	}

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
		model.AppName, model.FunctionName, clientAddress, remotePort, fullPath, model.NodeId, model.Msg, dataJsonText,
	)
	return logInfo
}


//----------------- new verson

// info log
func WriteNewLogInfo(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.INFO)
	defer log.Close()
	// check directory
	currentDirectory := fmt.Sprintf(`%s/info`, model.Driectory)
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.INFO, flw)
	logInfo := newCreateLogFormat(model)
	log.Info(logInfo)
}

// warning log
func WriteNewLogWarning(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.WARNING)
	defer log.Close()
	// check directory
	currentDirectory := fmt.Sprintf(`%s/warning`, model.Driectory)
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.WARNING, flw)
	logInfo := newCreateLogFormat(model)
	log.Warn(logInfo)
}

// error log
func WriteNewLogError(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.ERROR)
	defer log.Close()
	// check directory
	currentDirectory := fmt.Sprintf(`%s/error`, model.Driectory)
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.ERROR, flw)
	logInfo := newCreateLogFormat(model)
	log.Error(logInfo)
}

// critical log
func WriteNewLogCritical(model Log4goModel) {
	log := log4go.NewDefaultLogger(log4go.CRITICAL)
	defer log.Close()
	// check directory
	currentDirectory := fmt.Sprintf(`%s/critical`, model.Driectory)
	directoryfunc.CheckDirectory(currentDirectory)

	// generate file log writer formater
	flw := generateFormatLog(currentDirectory, model.AppName)

	log.AddFilter("log", log4go.CRITICAL, flw)
	logInfo := newCreateLogFormat(model)
	log.Critical(logInfo)
}
// for 
func newCreateLogFormat(model Log4goModel) string {
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

// ---------------------------------------------------------------------------------------------------------------//
func getCurrentDate(appName string) string {
	appname := formatString.StringToCamelCase(appName)
	now := time.Now()
	currentdate_str := fmt.Sprintf("%s.%s", appname, now.Format("2006-01-02"))
	return currentdate_str
}

// ---------------------------------------------------------------------------------------------------------------//
func generateFormatLog(currentDirectory string, appName string) *log4go.FileLogWriter {
	logFileName := fmt.Sprintf(`%s/%s.log`, currentDirectory, getCurrentDate(appName))
	flw := log4go.NewFileLogWriter(logFileName, false)
	flw.SetFormat("[%D %T] [%L], %M")
	return flw
}
