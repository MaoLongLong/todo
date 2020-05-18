package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func NewResponse(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(code, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(c *gin.Context, data interface{}, msg string) {
	NewResponse(c, http.StatusOK, data, msg)
}

func BadRequest(c *gin.Context, data interface{}, msg string) {
	NewResponse(c, http.StatusBadRequest, data, msg)
}

func InternalServerError(c *gin.Context, data interface{}, msg string) {
	NewResponse(c, http.StatusInternalServerError, data, msg)
}
