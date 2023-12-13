package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model `json:"-"`
	Id         uint      `json:"id" gorm:"primaryKey"`
	UserId     uint      `json:"-"`
	User       User      `json:"user" gorm:"foreignKey:UserId"`
	Answer     string    `json:"answer" gorm:"answer"`
	QuestionId uint      `json:"-"`
	Question   *Question `json:"question" gorm:"QuestionId"`
}
