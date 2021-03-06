package usecase

import (
	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/service/user"
)

type userUsecase struct {
	psqlUserRepo user.PsqlUserRepositoryInf
}

func NewUserUsecase(uRepo user.PsqlUserRepositoryInf) user.UserUsecaseInf {
	return &userUsecase{
		psqlUserRepo: uRepo,
	}
}

func (u userUsecase) FetchAll() ([]*models.User, error) {
	return u.psqlUserRepo.FetchAll()
}

func (u userUsecase) FetchOneById(id int64) (*models.User, error) {
	return u.psqlUserRepo.FetchOneById(id)
}

func (u userUsecase) Create(user *models.User) error {
	return u.psqlUserRepo.Create(user)
}
