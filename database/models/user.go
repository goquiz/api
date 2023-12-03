package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id           uint   `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" gorm:"username"`
	Email        string `json:"email" gorm:"email"`
	Password     string `json:"-" gorm:"password"`
	PasswordSalt string `json:"-" gorm:"password_salt"`
	Quizzes      []Quiz `gorm:"foreignKey:UserId"`
}
