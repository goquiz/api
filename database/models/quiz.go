package models

type Quiz struct {
	Base      `json:"-"`
	Id        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"name"`
	UserId    uint       `json:"-"`
	User      *User      `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserId"`
	Questions []Question `json:"questions,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:QuizId"`
}
