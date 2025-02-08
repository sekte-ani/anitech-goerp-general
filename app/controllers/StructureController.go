package controllers

import (
	"anierp/app/models"
	"anierp/database"
	"anierp/responses"

	"github.com/gin-gonic/gin"
)

func IndexStructure(ctx *gin.Context) {
	var employee []models.Employee
	var divisionID uint

	err := database.DB.Table("divisions").Where("name = ?", "Manajemen Eksekutif").Pluck("id", &divisionID).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Division Not Found",
		})
	}

	err = database.DB.Table("employees").Where("id != ?", 1).Where("division_id = ?", divisionID).Find(&employee).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	var structureResponses []responses.StructureResponse
	for _, emp := range employee {
		var roles []string

		err := database.DB.Table("role_employee").Select("roles.name").Joins("JOIN roles ON role_employee.role_id = roles.id").Where("role_employee.employee_id = ?", emp.ID).Pluck("roles.name", &roles).Error
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Role Doesn't Exists",
			})
			return
		}

		structureResponses = append(structureResponses, responses.StructureResponse{
			ID:        emp.ID,
			Name:      emp.Name,
			RoleName:  roles,
			Images:    emp.Images,
		})
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data": structureResponses,
	})
}