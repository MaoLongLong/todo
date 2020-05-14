package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func OK(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusOK, data, msg)
}

func BadRequest(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusBadRequest, data, msg)
}

func InternalServerError(c *gin.Context, data interface{}, msg string) {
	Response(c, http.StatusInternalServerError, data, msg)
}
