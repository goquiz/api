package models

import "time"

type EmailVerification struct {
	Base
	Id         uint      `json:"id" gorm:"primaryKey"`
	Expiration time.Time `json:"expiration" gorm:"expiration"`
	UserId     uint      `json:"-"`
	User       *User     `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserId"`
	Token      string    `json:"token" gorm:"token"`
}
