package repository

import (
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
)

type quiz struct{}

var Quiz quiz

func (quiz) AllForUser(userId uint) []models.Quiz {
	var quizzes []models.Quiz
	database.Database.Model(models.Quiz{}).
		Where("user_id = ?", userId).
		Find(&quizzes)
	return quizzes
}

func (quiz) WithQuestions(quizId uint, userId uint) *models.Quiz {
	var quiz models.Quiz
	database.Database.Model(models.Quiz{}).
		Preload("Questions").
		Where("id = ? and user_id = ?", quizId, userId).
		Find(&quiz)
	return &quiz
}
