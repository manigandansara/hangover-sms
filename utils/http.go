package utils

import (
	"hangover/structs"

	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, message string, code int, data interface{}) {
	c.JSON(code, structs.Response{Code: code, Msg: message, Data: data})
}
