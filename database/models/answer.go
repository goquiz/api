package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Id       uint     `json:"id" gorm:"primaryKey"`
	User     User     `json:"user" gorm:"user"`
	Question Question `json:"question" gorm:"question"`
	Answer   string   `json:"answer" gorm:"answer"`
}
