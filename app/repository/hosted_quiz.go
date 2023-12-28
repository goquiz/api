package repository

import (
	"fmt"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/helpers"
	"gorm.io/gorm"
)

type hosted_quiz struct{}

var HostedQuiz hosted_quiz

// NewUniqueCode generates a unique code for a quiz to host
func (hosted_quiz) NewUniqueCode() string {
	var publicKey int
	for {
		var count int64
		random := helpers.NewRandom()
		publicKey = random.Number(100000, 999999)
		database.Database.Model(&models.HostedQuiz{}).
			Where("public_key = ?", fmt.Sprintf("%v", publicKey)).
			Limit(1).
			Count(&count)
		if count == 0 {
			break
		}
	}
	return fmt.Sprintf("%v", publicKey)
}

// FindForUser returns a models.HostedQuiz by an id for a specific user
func (hosted_quiz) FindForUser(id uint, userId uint) *models.HostedQuiz {
	var hosted models.HostedQuiz
	database.Database.Joins("Quiz").Model(&hosted).
		Where("hosted_quizzes.id = ? and quiz.user_id = ?", id, userId).
		Find(&hosted)
	return &hosted
}

// AllForUser returns a list of models.HostedQuiz by a quiz and a user id
func (hosted_quiz) AllForUser(quizId uint, userId uint) []*models.HostedQuiz {
	var hosts []*models.HostedQuiz
	database.Database.Model(&models.HostedQuiz{}).
		Joins("Quiz").
		Order("updated_at DESC").
		Where("quiz.id = ? and quiz.user_id = ?", quizId, userId).
		Find(&hosts)
	return hosts
}

// FindByPublicKeyWithQuizUser returns a models.HostedQuiz with a preloaded quiz that contains a preloaded user
func (hosted_quiz) FindByPublicKeyWithQuizUser(publicKey string) *models.HostedQuiz {
	var hostedQuiz models.HostedQuiz
	database.Database.Model(&models.HostedQuiz{}).
		Preload("Quiz", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User")
		}).
		Where("public_key = ?", publicKey).
		Find(&hostedQuiz)
	return &hostedQuiz
}

func (hosted_quiz) FindByPublicKey(publicKey string) *models.HostedQuiz {
	var hostedQuiz models.HostedQuiz
	database.Database.Model(&models.HostedQuiz{}).
		Where("public_key = ?", publicKey).
		Find(&hostedQuiz)
	return &hostedQuiz
}
