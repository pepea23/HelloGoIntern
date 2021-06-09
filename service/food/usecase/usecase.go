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
	if err != nil {
		return nil, err
	}
	if len(foods) == 0 {
		return foods, err
	}

	for _, food := range foods {
		myFoods, err := f.psqlFoodRepo.FetchMyFoodFromFoodsId(food.Id)
		if err != nil {
			return nil, err
		}
		food.MyFoods = myFoods
	}
	return foods, err
}
