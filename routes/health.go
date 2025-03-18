package routes

import (
	"github.com/gin-gonic/gin"
	"hangover/controllers"
)

func HealthRoutes(router *gin.Engine) {
	health := new(controllers.HealthRepo)
	router.GET("/health", health.Status)
}
