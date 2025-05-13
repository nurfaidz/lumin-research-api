package routes

import (
	"project-research-gin/controllers"
	middleware "project-research-gin/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/api/")

	apiRouter.POST("register", controllers.Register)
	apiRouter.POST("login", controllers.Login)

	apiRouter.GET("users", middleware.AuthMiddleware(), controllers.FindUsers)
	apiRouter.POST("users", middleware.AuthMiddleware(), controllers.CreateUser)

	return router
}
