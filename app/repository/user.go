package repository

import (
	"fmt"
	"github.com/bndrmrtn/goquiz_api/database"
	"github.com/bndrmrtn/goquiz_api/database/models"
)

type user struct{}

var User user

func (user) IsUsernameOrEmailExists(username string, email string) bool {
	var count int64
	err := database.Database.Model(models.User{}).
		Where("username = ? or email = ?", username, email).
		Limit(1).
		Count(&count).
		Error

	if err != nil {
		fmt.Println("Query error:", err)
		return false
	}

	return count > 0
}

// FindByUsername returns the user with the given username or an error
func (u user) FindByUsername(username string) (models.User, error) {
	return u.FindBy([]string{"username"}, username)
}

// FindById returns the user with the given id or an error
func (u user) FindById(id uint) (models.User, error) {
	return u.FindBy([]string{"id"}, id)
}

// FindBy returns the user with the given fields and values or an error
func (user) FindBy(fields []string, values ...interface{}) (models.User, error) {
	var user models.User
	var query string
	for _, f := range fields {
		query += fmt.Sprintf(" %v = ?", f)
	}
	err := database.Database.Model(models.User{}).
		Where(query, values...).
		Limit(1).
		Find(&user).
		Error

	if err != nil {
		return user, err
	}

	return user, nil
}
