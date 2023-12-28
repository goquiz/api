package repository

import (
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
)

type question struct{}

var Question question

// ForQuiz returns a question by a question and a quiz id
func (question) ForQuiz(id uint, quizId uint) *models.Question {
	var question models.Question
	database.Database.Model(models.Question{}).
		Where("id = ? and quiz_id = ?", id, quizId).
		Find(&question)
	return &question
}

func (question) ForPlayers(quizId uint) []*models.Question {
	var questions []*models.Question
	database.Database.Model(models.Question{}).
		Select("id, question, image, answers").
		Where("quiz_id = ?", quizId).
		Find(&questions)
	return questions
}
