package models

import "gorm.io/gorm"

type HostedQuiz struct {
	gorm.Model `json:"-"`
	Id         uint   `json:"id" gorm:"primary_key"`
	PublicKey  string `json:"public_key" gorm:"public_key"`
	QuizId     uint
	Quiz       Quiz       `json:"quiz" gorm:"foreignKey:QuizId"`
	Questions  []Question `json:"questions" gorm:"foreignKey:HostedQuizId"`
}
