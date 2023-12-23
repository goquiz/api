package repository

import (
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
)

type quiz struct{}

var Quiz quiz

// AllForUser returns all the quizzes for a user
func (quiz) AllForUser(userId uint) []models.Quiz {
	var quizzes []models.Quiz
	database.Database.Model(models.Quiz{}).
		Order("updated_at DESC").
		Where("user_id = ?", userId).
		Find(&quizzes)
	return quizzes
}

// WithQuestions returns a specific quiz for a user with all the questions
func (quiz) WithQuestions(quizId uint, userId uint) *models.Quiz {
	var quiz models.Quiz
	database.Database.Model(models.Quiz{}).
		Preload("Questions").
		Where("id = ? and user_id = ?", quizId, userId).
		Find(&quiz)
	return &quiz
}

// IsBelongsToUser checks if the quiz belongs to a user or not
func (quiz) IsBelongsToUser(quizId uint, userId uint) bool {
	var count int64
	database.Database.Model(models.Quiz{}).
		Where("id = ? and user_id = ?", quizId, userId).
		Count(&count)
	return count > 0
}
