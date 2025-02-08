package models

import "time"

type Employee struct {
	ID 			uint 		`json:"id"`
	Name     	string 		`json:"name"`
	Slug     	string 		`json:"slug"`
	DivisionID  uint 		`json:"division_id"`
	RoleID     	uint 		`json:"role_id"`
	Images     	string 		`json:"images"`
	Phone     	string 		`json:"phone"`
	Address     string 		`json:"address"`
	BankName    string 		`json:"bank_name"`
	BankAccount string 		`json:"bank_account"`
	CreatedAt   time.Time 	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time 	`json:"updated_at" gorm:"autoUpdateTime"`
}

func (Employee) TableName() string {
	return "employees"
}