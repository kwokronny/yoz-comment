package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseStatus int

// ResponseStatus 常用返回Code枚举值
const (
	// Success is 正常返回
	ResponseSuccess ResponseStatus = 200
	// Authorized is 鉴权失败
	ResponseAuthorized ResponseStatus = 401
	// ParamError is 入参错误
	ResponseParamError ResponseStatus = 406
	// ServerError is 服务错误
	ResponseServerError ResponseStatus = 50
)

// Response is default return struct of API Request
type Response struct {
	Code    ResponseStatus `json:"code"`
	Message string         `json:"msg"`
	Data    interface{}    `json:"data"`
}

// Success is output success message and data struct
func (resp Response) Success(c *gin.Context, data interface{}) {
	var result Response
	result.Code = ResponseSuccess
	result.Message = "操作成功"
	result.Data = data
	c.JSON(http.StatusOK, result)
}

// Error is output error message struct
func (resp Response) Error(c *gin.Context, code ResponseStatus, message string) {
	var result Response
	result.Code = code
	result.Message = message
	c.JSON(http.StatusOK, result)
}
