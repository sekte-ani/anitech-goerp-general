package controllers

import (
	"anierp/app/models"
	"anierp/database"
	"anierp/responses"

	"github.com/gin-gonic/gin"
)

func IndexEmployee(ctx *gin.Context) {
	var employee []models.Employee
	err := database.DB.Table("employees").Find(&employee).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	var employeeResponses []responses.EmployeeResponse
	for _, emp := range employee {
		var roleName, divisionName string
		database.DB.Table("roles").Where("id = ?", emp.RoleID).Pluck("name", &roleName).Scan(&roleName)
		database.DB.Table("divisions").Where("id = ?", emp.DivisionID).Pluck("name", &divisionName).Scan(&divisionName)

		employeeResponses = append(employeeResponses, responses.EmployeeResponse{
			ID:        emp.ID,
			Name:      emp.Name,
			DivisionName: divisionName,
			RoleName:  roleName,
			Images:    emp.Images,
		})
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data": employeeResponses,
	})
}

// func ShowEmployee(ctx *gin.Context){
// 	id := ctx.Param("id")
// 	employee := new(responses.EmployeeResponse)

// 	err := database.DB.Table("users").Where("id = ?", id).Find(&employee).Error
// 	if err != nil {
// 		ctx.JSON(500, gin.H{
// 			"message": "Internal Server Error",
// 		})
// 		return
// 	}
	
// 	if employee.ID == nil {
// 		ctx.JSON(404, gin.H{
// 			"message": "User Not Found",
// 		})
// 		return
// 	}

// 	var roleName string
// 	err = database.DB.Table("roles").Where("id = ?", employee.RoleID).Pluck("name", &roleName).Error
// 	if err != nil {
// 		ctx.JSON(500, gin.H{
// 			"message": "Failed to Get Role Name",
// 		})
// 	}

// 	employee.RoleName = roleName

// 	ctx.JSON(200, gin.H{
// 		"message": "Success",
// 		"data": employee,
// 	})
// }

// func StoreUser(ctx *gin.Context){
// 	userReq := new(request.UserRequest)

// 	if err := ctx.ShouldBind(&userReq); err != nil {
// 		ctx.JSON(400, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	userEmailExist := new(models.User)
// 	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist)

// 	if userEmailExist.Email != "" {
// 		ctx.JSON(400, gin.H{
// 			"message": "Email Already Exist",
// 		})
// 		return
// 	}

// 	user := new(models.User)
// 	user.Name = userReq.Name
// 	user.Email = userReq.Email
// 	user.Password, _ = utilities.HashPassword(userReq.Password)
// 	user.RoleID = userReq.RoleID

// 	err := database.DB.Table("users").Create(&user).Error
// 	if err != nil {
// 		ctx.JSON(500, gin.H{
// 			"message": "Create Data Failed",
// 		})
// 		return
// 	}

// 	ctx.JSON(200, gin.H{
// 		"message": "success",
// 		"data":    user,
// 	})
// }