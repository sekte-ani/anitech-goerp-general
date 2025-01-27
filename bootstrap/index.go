package bootstrap

import (
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

	routes.InitRoutes(app)
	app.Run(":" + appconfig.PORT)
}