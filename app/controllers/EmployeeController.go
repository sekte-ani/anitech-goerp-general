package controllers

import (
	"anierp/app/models"
	"anierp/database"
	"anierp/responses"

	"github.com/gin-gonic/gin"
)

func IndexEmployee(ctx *gin.Context) {
	divisionName := ctx.Query("division")

	var employees []models.Employee
	query := database.DB.Table("employees").Where("id != ?", 1)

	if divisionName != "" {
		var divisionID uint
		err := database.DB.Table("divisions").Select("id").Where("name = ?", divisionName).Scan(&divisionID).Error
		if err != nil || divisionID == 0 {
			ctx.JSON(400, gin.H{"message": "Division Not Found"})
			return
		}
		query = query.Where("division_id = ?", divisionID)
	}

	err := query.Find(&employees).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	var employeeResponses []responses.EmployeeResponse
	for _, emp := range employees {
		var roles []string

		err := database.DB.Table("role_employee").Select("roles.name").Joins("JOIN roles ON role_employee.role_id = roles.id").Where("role_employee.employee_id = ?", emp.ID).Pluck("roles.name", &roles).Error
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "Role Doesn't Exists",
			})
			return
		}

		employeeResponses = append(employeeResponses, responses.EmployeeResponse{
			ID:           emp.ID,
			Name:         emp.Name,
			RoleName:     roles,
			Images:       emp.Images,
		})
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    employeeResponses,
	})
}