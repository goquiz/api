package models

import (
	"time"
)

type ResetPassword struct {
	Base       `json:"-"`
	Id         uint      `json:"id" gorm:"primaryKey"`
	Token      string    `json:"token" gorm:"token"`
	UserId     uint      `json:"-"`
	User       User      `json:"user" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserId"`
	Expiration time.Time `json:"expiry" gorm:"expiration"`
}
