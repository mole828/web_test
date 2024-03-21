package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger logrus.Logger = *logrus.New()

func main() {

	app := gin.New()

	app.Use(gin.Recovery())

	count := 0
	app.GET("/count", func(ctx *gin.Context) {
		count += 1
		logger.WithField("count", count).Info("count add")
		ctx.JSON(200, gin.H{"count": count})
	})

	logger.Info("server running")
	app.Run(":8080")
}
