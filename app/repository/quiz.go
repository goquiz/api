package repository

import "github.com/bndrmrtn/goquiz_api/database/models"

type quiz struct{}

var Quiz quiz

func (quiz) All() []models.Quiz {
	return []models.Quiz{}
}
