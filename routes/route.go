package routes

import (
	"project-research-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.Register)

	return router
}
