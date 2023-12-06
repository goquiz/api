package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Id         uint `json:"id" gorm:"primaryKey"`
	UserId     uint
	User       User   `json:"user" gorm:"foreignKey:UserId"`
	Answer     string `json:"answer" gorm:"answer"`
	QuestionId uint
	Question   Question `json:"question" gorm:"QuestionId"`
}
