package models

import (
	"gorm.io/gorm"
)

type HostedQuiz struct {
	gorm.Model  `json:"-"`
	Id          uint         `json:"id" gorm:"primary_key"`
	Name        string       `json:"name" gorm:"name"`
	PublicKey   string       `json:"public_key" gorm:"public_key"`
	QuizId      uint         `json:"-"`
	Quiz        *Quiz        `json:"quiz,omitempty" gorm:"foreignKey:QuizId"`
	IsActive    bool         `json:"is_active" gorm:"is_active,default:true"`
	UserAnswers []UserAnswer `json:"user_answer,omitempty" gorm:"foreignKey:HostedQuizId"`
}
