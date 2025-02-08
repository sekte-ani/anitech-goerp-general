package routes

import (
	"anierp/app/controllers"
	appconfig "anierp/config/app_config"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine){
	route := app
	route.Static(appconfig.STATIC_ROUTE, appconfig.STATIC_DIR)
	
	route.GET("/api/employee", controllers.IndexEmployee)
	route.GET("/api/structure", controllers.IndexStructure)
	route.GET("/api/product", controllers.IndexProduct)

	route.POST("/visitor", func (ctx *gin.Context) { 
		ctx.JSON(200, gin.H{
			"message": "Visitor tracked",
		})
	})
}