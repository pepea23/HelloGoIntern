package user

import (
	"github.com/HelloGoIntern/models"
)

type UserUsecaseInf interface {
	FetchAll() ([]*models.User, error)
	FetchOneById(id int64) (*models.User, error)
	Create(user *models.User) error
}
