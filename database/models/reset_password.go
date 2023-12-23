package models

import (
	"gorm.io/gorm"
	"time"
)

type ResetPassword struct {
	gorm.Model `json:"-"`
	Id         uint      `json:"id" gorm:"primaryKey"`
	UserId     uint      `json:"-"`
	User       User      `json:"user" gorm:"foreignKey:UserId"`
	Expiry     time.Time `json:"expiry" gorm:"expiry"`
}
