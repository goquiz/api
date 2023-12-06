package models

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"name"`
	IsActivated bool   `json:"is_activated" gorm:"is_activated,default:false"`
	UserId      uint
	User        User `gorm:"foreignKey:UserId"`
}
