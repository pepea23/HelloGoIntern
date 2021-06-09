package food

import "github.com/HelloGoIntern/models"

type FoodRepositorynterface interface {
	CreateFood(food *models.Food) error
	FetchAllFoods() ([]*models.Food, error)
	FetchMyFoodFromFoodsId(id int64) (models.MyFoods, error)
}
