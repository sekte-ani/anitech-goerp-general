package routes

import (
	"anierp/app/controllers"
	appconfig "anierp/config/app_config"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine){
	route := app
	route.Static(appconfig.STATIC_ROUTE, appconfig.STATIC_DIR)

	// AUTH ROUTES
	// authRoute := route.Group("/api/auth")
	// authRoute.POST("/login", controllers.Login)
	
	// USER ROUTES
	userRoute := route.Group("/api/employee")
	userRoute.GET("/", controllers.IndexEmployee)
	// userRoute.GET("/:id", controllers.ShowEmployee)
}