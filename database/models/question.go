package models

import (
	"gorm.io/datatypes"
)

type Question struct {
	Base     `json:"-"`
	Id       uint           `json:"id" gorm:"primaryKey"`
	Question string         `json:"question" gorm:"question"`
	Image    *string        `json:"image" gorm:"image,default:null"`
	Answers  datatypes.JSON `json:"answers" gorm:"answers"`
	Answer   string         `json:"answer,omitempty" gorm:"answer"`
	QuizId   uint           `json:"-"`
	Quiz     *Quiz          `json:"quiz,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:QuizId"`
}
