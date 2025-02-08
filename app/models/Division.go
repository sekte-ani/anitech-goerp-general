package models

import "time"

type Division struct {
	ID 			uint 		`json:"id"`
	Name     	string 		`json:"name"`
	Slug     	string 		`json:"slug"`
	CreatedAt   time.Time 	`json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time 	`json:"updated_at" gorm:"autoUpdateTime"`
}

func (Division) TableName() string {
	return "divisions"
}