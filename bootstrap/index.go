package bootstrap

import (
	"anierp/app/middleware"
	"anierp/config"
	appconfig "anierp/config/app_config"
	corsconfig "anierp/config/cors_config"
	"anierp/database"
	"anierp/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp(){
	err := godotenv.Load()
	if err != nil {
		log.Println("Can't Connect to .env file")
	}

	config.InitConfig()
	database.ConnectDB()

	app := gin.Default()
	app.Use(corsconfig.CorsConfig())
	app.Use(middleware.VisitorMiddleware())

	routes.InitRoutes(app)
	app.Run(":" + appconfig.PORT)
}