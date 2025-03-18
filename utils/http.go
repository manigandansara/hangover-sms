package utils

import (
	"github.com/gin-gonic/gin"
	"hangover/structs"
)

func SuccessResponse(c *gin.Context, message string, code int, data interface{}) {
	c.JSON(200, structs.Response{code, message, data})
}

func ErrorResponse(c *gin.Context, message string, code int, data interface{}) {
	c.JSON(200, structs.Response{code, message, data})
}
