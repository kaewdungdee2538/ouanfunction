package log4go

import (
	"github.com/gin-gonic/gin"
)

type Log4goModel struct {
	Context      *gin.Context           `json:"context" validate:"required"`
	AppName      string                 `json:"appName" validate:"required"`
	Driectory    string                 `json:"driectory" validate:"required"`
	FunctionName string                 `json:"functionName" validate:"required"`
	Msg          string                 `json:"msg" validate:"required"`
	NodeId       string                 `json:"nodeId" validate:"required"`
	Data         map[string]interface{} `json:"data"`
}
