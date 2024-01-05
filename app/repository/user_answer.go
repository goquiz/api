package repository

import (
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"gorm.io/gorm"
)

type userAnswer struct{}

var UserAnswer userAnswer

func (userAnswer) IsUserAlreadyPlayed(quizId uint, userId uint) bool {
	var count int64
	database.Database.Model(&models.UserAnswer{}).
		Where("quiz_id = ? and user_id = ?", quizId, userId).
		Limit(1).
		Count(&count)

	return count > 0
}

func (userAnswer) Paginate(userId uint, c int, p int, specQuizId uint) []*models.UserAnswer {
	offset := c * (p - 1) // the count of the quizzes multiplied by the page (-1 required cuz page 1 has no offset)

	var userAnswers []*models.UserAnswer
	database.Database.Model(&models.UserAnswer{}).
		Preload("Answers").
		Preload("HostedQuiz", func(db *gorm.DB) *gorm.DB {
			if specQuizId != 0 {
				db.Where("quiz_id = ?", specQuizId)
			}
			return db.Preload("Quiz", func(db *gorm.DB) *gorm.DB {
				return db.Preload("User").Preload("Questions")
			})
		}).
		Order("updated_at desc").
		Where("user_id = ?", userId).
		Limit(c).
		Offset(offset).
		Find(&userAnswers)

	return userAnswers
}
