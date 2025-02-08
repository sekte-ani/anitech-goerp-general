package middleware

import (
	"anierp/app/models"
	"anierp/database"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func VisitorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIP := ctx.ClientIP()
		now := time.Now()

		var lastVisit models.Visitor
		err := database.DB.Where("ip_address = ?", userIP).Order("visit_time DESC").First(&lastVisit).Error

		if err == nil && now.Sub(lastVisit.VisitTime) < 24*time.Hour {
			ctx.Next()
			return
		}

		visitor := models.Visitor{
			IPAddress: userIP,
			VisitTime: now,
		}
		if err := database.DB.Create(&visitor).Error; err != nil {
			fmt.Println("Database error:", err)
		}

		ctx.Next()
	}
}