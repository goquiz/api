package repository

import (
	"errors"
	"fmt"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/helpers"
	"time"
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
	var u models.User
	var query string
	for _, f := range fields {
		query += fmt.Sprintf(" %v = ?", f)
	}
	err := database.Database.Model(models.User{}).
		Where(query, values...).
		Limit(1).
		Find(&u).
		Error

	if err != nil {
		return u, err
	}

	return u, nil
}

func (user) EmailVerification(t string) (*models.EmailVerification, error) {
	var ev *models.EmailVerification
	database.Database.Model(&models.EmailVerification{}).
		Preload("User").
		Where("token = ? and expiration > ?", t, time.Now()).
		Find(&ev)
	if ev.Id == 0 {
		return nil, errors.New("invalid or expired verification token")
	}

	return ev, nil
}

func (user) HasRequestedNewPassword(uId uint) bool {
	var count int64
	database.Database.Model(&models.EmailVerification{}).
		Joins("User").
		Where("email_verifications.expire > ? and User.id = ?", time.Now(), uId).
		Limit(1).
		Count(&count)
	return count > 0
}

func (user) ResetPassword(t string) (*models.ResetPassword, error) {
	var rp *models.ResetPassword
	database.Database.Model(&models.ResetPassword{}).
		Preload("User").
		Where("token = ? and expiration > ?", t, time.Now()).
		Find(&rp)
	if rp.Id == 0 {
		return nil, errors.New("invalid or expired reset password token")
	}

	return rp, nil
}

func (user) NewTokenFor(model interface{}) string {
	var token string
	for {
		var count int64
		random := helpers.NewRandom()
		token = random.String(random.Number(15, 50))
		database.Database.Model(&model).
			Where("token = ?", token).
			Limit(1).
			Count(&count)
		if count == 0 {
			break
		}
	}
	return token
}
