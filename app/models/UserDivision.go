package models

import "time"

type UserDivision struct {
	ID 			uint 		`json:"id"`
	EmployeeID  uint 		`json:"employee_id"`
	DivisionID  uint 		`json:"division_id"`
	CreatedAt   time.Time 	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time 	`json:"updated_at" gorm:"autoUpdateTime"`
}

func (UserDivision) TableName() string {
	return "user_division"
}