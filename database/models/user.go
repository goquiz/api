package models

import "gorm.io/gorm"

type User struct {
	gorm.Model   `json:"-"`
	Id           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" gorm:"username"`
	Email        string `json:"email" gorm:"email"`
	Password     string `json:"-" gorm:"password"`
	PasswordSalt string `json:"-" gorm:"password_salt"`
	Quizzes      []Quiz `json:"quizzes,omitempty" gorm:"foreignKey:UserId"`
}
