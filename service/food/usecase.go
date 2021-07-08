package food

import (
	"github.com/HelloGoIntern/models"
	"github.com/gofrs/uuid"
)

type FoodUseCaseInterface interface {
	CreateFood(food *models.Food) error
	FetchAllFoods() ([]*models.Food, error)
	FetchFoodFromFoodsName(FoodName string) ([]*models.Food, error)
	FetchFoodFromTypeOfFood(TypeOfFood string) ([]*models.Food, error)
	FetchFoodFromPrice(Price string) ([]*models.Food, error)
	Deletefood(id uuid.UUID) error
}
