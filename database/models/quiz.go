package models

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model `json:"-"`
	Id         uint       `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name" gorm:"name"`
	UserId     uint       `json:"-"`
	User       *User      `json:"user,omitempty" gorm:"foreignKey:UserId"`
	Questions  []Question `json:"questions,omitempty" gorm:"foreignKey:QuizId"`
}
