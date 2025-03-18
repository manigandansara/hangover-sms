package routes

import (
	"github.com/gin-gonic/gin"
	"hangover/controllers"
)

func AuthRoutes(router *gin.RouterGroup) {
	auth := new(controllers.AuthRepo)
	router.POST("/login", auth.Login)
}
