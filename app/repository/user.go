package repository

import (
	"fmt"
	"github.com/bndrmrtn/goquiz_api/database"
	"github.com/bndrmrtn/goquiz_api/database/models"
)

type user struct{}

var User user

func (user) IsUsernameExists(username string) bool {
	var count int64
	err := database.Database.Model(models.User{}).
		Where("username = ?", username).
		Limit(1).
		Count(&count).
		Error

	if err != nil {
		fmt.Println("Query error:", err)
		return false
	}

	return count > 0
}
