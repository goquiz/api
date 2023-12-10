package authorized

import (
	"github.com/goquiz/api/app/repository"
	"github.com/goquiz/api/database/models"
)

type _user struct {
	User *models.User
}

var Authorized _user

func (u *_user) AuthUser(id uint) error {
	user, err := repository.User.FindById(id)
	if err != nil {
		return err
	}
	u.User = &user
	return nil
}
