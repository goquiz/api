package models

import (
	"time"
)

type User struct {
	Base            `json:"-"`
	Id              uint       `json:"id" gorm:"primaryKey"`
	Username        string     `json:"username" gorm:"username"`
	Email           string     `json:"email" gorm:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty" gorm:"email_verified_at,default:null"`
	ProfileImageURL *string    `json:"profile_image_url,omitempty" gorm:"profile_image_url,default:null"`
	Password        string     `json:"-" gorm:"password"`
	PasswordSalt    string     `json:"-" gorm:"password_salt"`
	Quizzes         []Quiz     `json:"quizzes,omitempty" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserId"`
}
