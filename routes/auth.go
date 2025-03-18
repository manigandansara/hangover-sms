package routes

import (
	"github.com/gin-gonic/gin"
	"hangover/controllers"
)

func AuthRoutes(router *gin.RouterGroup) {
	health := new(controllers.HealthRepo)
	router.GET("/health", health.Status)
}
