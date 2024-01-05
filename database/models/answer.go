package models

type Answer struct {
	Base         `json:"-"`
	Id           uint        `json:"id" gorm:"primaryKey"`
	UserAnswerId uint        `json:"-"`
	UserAnswer   *UserAnswer `json:"user_answer,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserAnswerId"`
	Answer       string      `json:"answer" gorm:"answer"`
	QuestionId   uint        `json:"-"`
	Question     *Question   `json:"question,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:QuestionId"`
}
