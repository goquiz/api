package models

type UserAnswer struct {
	Base         `json:"-"`
	Id           uint        `json:"id" gorm:"primaryKey"`
	UserId       uint        `json:"-"`
	User         *User       `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserId"`
	HostedQuizId uint        `json:"-"`
	HostedQuiz   *HostedQuiz `json:"hosted_quiz,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:HostedQuizId"`
	Answers      []*Answer   `json:"answers,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserAnswerId"`
}
