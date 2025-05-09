package routes

import (
	"project-research-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/api/")

	apiRouter.POST("register", controllers.Register)
	apiRouter.POST("login", controllers.Login)

	return router
}
