package user

import (
	"github.com/HelloGoIntern/models"
)

type PsqlUserRepositoryInf interface {
	FetchAll() ([]*models.User, error)
	FetchOneById(id int64) (*models.User, error)
	Create(user *models.User) error
}
