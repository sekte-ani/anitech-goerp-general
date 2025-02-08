package models

import "time"

type Visitor struct {
	ID       		uint 		`json:"id"`
	IPAddress     	string 		`json:"ip_address" gorm:"size:45;not null"`
	VisitTime    	time.Time 	`json:"visit_time" gorm:"autoCreateTime"`
}