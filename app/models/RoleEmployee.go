package models

import "time"

type RoleEmployee struct {
	ID 			uint 		`json:"id"`
	EmployeeID  uint 		`json:"employee_id"`
	RoleID     	uint 		`json:"role_id"`
	CreatedAt   time.Time 	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time 	`json:"updated_at" gorm:"autoUpdateTime"`
}

func (RoleEmployee) TableName() string {
	return "role_employee"
}