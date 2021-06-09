package usecase

import (
	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/service/food"
)

type foodUseCase struct {
	psqlFoodRepo food.FoodRepositorynterface
}

func NewFoodUsecase(repo food.FoodRepositorynterface) food.FoodUseCaseInterface {
	return &foodUseCase{
		psqlFoodRepo: repo,
	}
}

func (f foodUseCase) CreateFood(food *models.Food) error {
	err := f.psqlFoodRepo.CreateFood(food)
	return err
}

func (f foodUseCase) FetchAllFoods() ([]*models.Food, error) {
	foods, err := f.psqlFoodRepo.FetchAllFoods()
	return foods, err
}
