package main

import (
	"os"
	"project-research-gin/config"
	"project-research-gin/database"
	"project-research-gin/routes"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	r := routes.SetupRoutes()

	r.Run(":" + os.Getenv("APP_PORT"))
}
