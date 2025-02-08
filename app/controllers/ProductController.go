package controllers

import (
	"anierp/app/models"
	"anierp/database"
	"anierp/responses"

	"github.com/gin-gonic/gin"
)

func IndexProduct(ctx *gin.Context) {
	products := new([]models.Product)

	err := database.DB.Table("products").Find(&products).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Product Not Found",
		})
	}

	var productResponses []responses.ProductResponse
	for _, product := range *products {
		productResponses = append(productResponses, responses.ProductResponse{
			ID: product.ID,
			Name: product.Name,
			Description: product.Description,
			Images: product.Images,
		})
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data": productResponses,
	})
}