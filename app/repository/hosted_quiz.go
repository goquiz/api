package repository

import (
	"fmt"
	"github.com/goquiz/api/database"
	"github.com/goquiz/api/database/models"
	"github.com/goquiz/api/helpers"
)

type hosted_quiz struct{}

var HostedQuiz hosted_quiz

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

func (hosted_quiz) FindForUser(id uint, userId uint) *models.HostedQuiz {
	var hosted models.HostedQuiz
	database.Database.Joins("Quiz").Model(&hosted).
		Where("hosted_quizzes.id = ? and quiz.user_id = ?", id, userId).
		Find(&hosted)
	return &hosted
}

func (hosted_quiz) AllForUser(userId uint) []*models.HostedQuiz {
	var hosts []*models.HostedQuiz
	database.Database.Joins("Quiz").
		Where("quiz.user_id = ?", userId).
		Find(&hosts)
	return hosts
}
