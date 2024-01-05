package models

type HostedQuiz struct {
	Base        `json:"-"`
	Id          uint         `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"name"`
	PublicKey   string       `json:"public_key" gorm:"public_key"`
	QuizId      uint         `json:"-"`
	Quiz        *Quiz        `json:"quiz,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:QuizId"`
	IsActive    bool         `json:"is_active" gorm:"is_active,default:true"`
	UserAnswers []UserAnswer `json:"user_answer,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:HostedQuizId"`
}
