package models

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	Id        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"name"`
	Questions []Question `json:"questions" gorm:"foreignKey:QuizId"`
	UserId    uint
	User      User `gorm:"foreignKey:UserId"`
}
