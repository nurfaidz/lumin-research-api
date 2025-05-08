package main

import (
	"os"
	"project-research-gin/config"
	"project-research-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})

	})

	router.Run(":" + os.Getenv("APP_PORT"))
}
