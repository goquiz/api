package repository

import (
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
)

type question struct{}

var Question question

func (question) ForQuiz(id uint, quizId uint) *models.Question {
	var question models.Question
	database.Database.Model(models.Question{}).
		Where("id = ? and quiz_id = ?", id, quizId).
		Find(&question)
	return &question
}
