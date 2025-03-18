package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/robertantonyjaikumar/hangover-common/database"
	"github.com/robertantonyjaikumar/hangover-common/logger"
	"go.uber.org/zap"
	"net/http"
)

type HealthRepo struct{}

func (h HealthRepo) Status(c *gin.Context) {
	var str string
	database.Db.Raw("select login from res_users order by id desc limit 1").Scan(&str)

	logger.Info("str", zap.String("str", str))
	c.String(http.StatusOK, "Gin-starter is working")
}
