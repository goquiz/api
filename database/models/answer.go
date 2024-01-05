package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model   `json:"-"`
	Id           uint        `json:"id" gorm:"primaryKey"`
	UserAnswerId uint        `json:"-"`
	UserAnswer   *UserAnswer `json:"user_answer,omitempty" gorm:"foreignKey:UserAnswerId"`
	Answer       string      `json:"answer" gorm:"answer"`
	QuestionId   uint        `json:"-"`
	Question     *Question   `json:"question,omitempty" gorm:"QuestionId"`
}
