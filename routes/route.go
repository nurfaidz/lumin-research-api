package routes

import (
	"project-research-gin/controllers"
	middleware "project-research-gin/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	apiRouter := router.Group("/api/")

	apiRouter.POST("register", controllers.Register)
	apiRouter.POST("login", controllers.Login)

	apiRouter.GET("users", middleware.AuthMiddleware(), controllers.FindUsers)
	apiRouter.POST("users", middleware.AuthMiddleware(), controllers.CreateUser)
	apiRouter.GET("users/:id", middleware.AuthMiddleware(), controllers.FindUserById)
	apiRouter.PUT("users/:id", middleware.AuthMiddleware(), controllers.UpdateUser)
	apiRouter.DELETE("users/:id", middleware.AuthMiddleware(), controllers.DeleteUser)

	return router
}
