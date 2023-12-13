package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model `json:"-"`
	Id         uint           `json:"id" gorm:"primaryKey"`
	Question   string         `json:"question" gorm:"question"`
	Image      *string        `json:"image" gorm:"image,default:null"`
	Answers    datatypes.JSON `json:"answers" gorm:"answers"`
	Answer     string         `json:"answer" gorm:"answer"`
	QuizId     uint           `json:"-"`
	Quiz       *Quiz          `json:"quiz,omitempty" gorm:"foreignKey:QuizId"`
}
