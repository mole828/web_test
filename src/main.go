package main

import (
	"web_test/src/cache"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var logger logrus.Logger = *logrus.New()

func main() {
	app_id := uuid.New().String()
	logger.Println(logrus.Fields{"app_id": app_id})

	app := gin.New()
	app.Use(gin.Recovery())

	count := cache.Default()
	app.GET("/count", func(ctx *gin.Context) {
		mulit_count, err := count.Inc("multi_count")
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err,
			})
		}
		logger.WithField("count", mulit_count).Info("count add")
		ctx.JSON(200, gin.H{"count": mulit_count, "app_id": app_id})
	})

	logger.Info("server running")
	app.Run(":8080")
}
