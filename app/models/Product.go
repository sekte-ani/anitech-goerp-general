package models

import "time"

type Product struct {
	ID 				uint 		`json:"id"`
	Name     		string 		`json:"name"`
	Slug     		string 		`json:"slug"`
	Description     string 		`json:"description"`
	Images     		string 		`json:"images"`
	CreatedAt   	time.Time 	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   	time.Time 	`json:"updated_at" gorm:"autoUpdateTime"`
}

func (Product) TableName() string {
	return "products"
}