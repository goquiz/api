package models

import (
	"time"
)

type ResetPassword struct {
	Base   `json:"-"`
	Id     uint      `json:"id" gorm:"primaryKey"`
	UserId uint      `json:"-"`
	User   User      `json:"user" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserId"`
	Expiry time.Time `json:"expiry" gorm:"expiry"`
}
