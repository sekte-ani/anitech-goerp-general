package routes

import (
	"anierp/app/controllers"
	appconfig "anierp/config/app_config"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine){
	route := app
	route.Static(appconfig.STATIC_ROUTE, appconfig.STATIC_DIR)
	
	// USER ROUTES
	userRoute := route.Group("/api/employee")
	userRoute.GET("/", controllers.IndexEmployee)
}