package food

import "github.com/HelloGoIntern/models"

type FoodUseCaseInterface interface {
	CreateFood(food *models.Food) error
	FetchAllFoods() ([]*models.Food, error)
}
