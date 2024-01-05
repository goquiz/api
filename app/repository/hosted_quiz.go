package repository

import (
	"fmt"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/helpers"
	"gorm.io/gorm"
)

type hostedQuiz struct{}

var HostedQuiz hostedQuiz

// NewUniqueCode generates a unique code for a quiz to host
func (hostedQuiz) NewUniqueCode() string {
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
func (hostedQuiz) FindForUser(id uint, userId uint) *models.HostedQuiz {
	var hosted models.HostedQuiz
	database.Database.Model(&hosted).
		Joins("Quiz").
		Where("hosted_quizzes.id = ? and Quiz.user_id = ?", id, userId).
		Find(&hosted)
	return &hosted
}

// AllForUser returns a list of models.HostedQuiz by a quiz and a user id
func (hostedQuiz) AllForUser(quizId uint, userId uint) []*models.HostedQuiz {
	var hosts []*models.HostedQuiz
	database.Database.Model(&models.HostedQuiz{}).
		Joins("Quiz").
		Order("updated_at DESC").
		Where("Quiz.id = ? and Quiz.user_id = ?", quizId, userId).
		Find(&hosts)
	return hosts
}

// FindByPublicKeyWithQuizUser returns a models.HostedQuiz with a preloaded quiz that contains a preloaded user
func (hostedQuiz) FindByPublicKeyWithQuizUser(publicKey string) *models.HostedQuiz {
	var hostedQuiz models.HostedQuiz
	database.Database.Model(&models.HostedQuiz{}).
		Preload("Quiz", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User")
		}).
		Where("public_key = ?", publicKey).
		Find(&hostedQuiz)
	return &hostedQuiz
}

func (hostedQuiz) FindByPublicKey(publicKey string) *models.HostedQuiz {
	var hostedQuiz models.HostedQuiz
	database.Database.Model(&models.HostedQuiz{}).
		Where("public_key = ?", publicKey).
		Find(&hostedQuiz)
	return &hostedQuiz
}

func (hostedQuiz) CountForQuizId(quizId uint) int64 {
	var count int64
	database.Database.Model(&models.HostedQuiz{}).
		Where("quiz_id = ?", quizId).
		Count(&count)
	return count
}

func (hostedQuiz) PaginateSubmissions(hostId uint, c int, p int) *models.HostedQuiz {
	offset := c * (p - 1) // the count of the quizzes multiplied by the page (-1 required cuz page 1 has no offset)

	var submission *models.HostedQuiz

	database.Database.Model(&models.HostedQuiz{}).
		Select("id, name").
		Preload("UserAnswers", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, username, profile_image_url")
			}).
				Preload("Answers", func(db *gorm.DB) *gorm.DB {
					return db.Preload("Question")
				}).Limit(c).Offset(offset)
		}).
		Order("updated_at desc").
		Where("id = ?", hostId).
		Find(&submission)

	return submission
}
