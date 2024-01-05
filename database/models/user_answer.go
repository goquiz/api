package models

import "gorm.io/gorm"

type UserAnswer struct {
	gorm.Model   `json:"-"`
	Id           uint        `json:"id" gorm:"primaryKey"`
	UserId       uint        `json:"-"`
	User         *User       `json:"user,omitempty" gorm:"foreignKey:UserId"`
	HostedQuizId uint        `json:"-"`
	HostedQuiz   *HostedQuiz `json:"hosted_quiz,omitempty" gorm:"foreignKey:HostedQuizId"`
	Answers      []*Answer   `json:"answers,omitempty" gorm:"foreignKey:UserAnswerId"`
}
