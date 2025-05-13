package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// BadRequest 返回错误请求响应
func BadRequest(c *gin.Context, message string, err ...interface{}) {
	res := Response{
		Code:    http.StatusBadRequest,
		Message: message,
	}
	if len(err) > 0 {
		res.Error = err[0].(string)
	}
	c.JSON(http.StatusBadRequest, res)
}

func InternalServerError(c *gin.Context, message string, err ...interface{}) {
	res := Response{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
	if len(err) > 0 {
		res.Error = err[0].(string)
	}
	c.JSON(http.StatusInternalServerError, res)
}
