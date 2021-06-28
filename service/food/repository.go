package food

import (
	"database/sql"
	"sync"

	"github.com/HelloGoIntern/models"
	"github.com/gofrs/uuid"
)

type FoodRepositorynterface interface {
	CreateFood(food *models.Food, tx *sql.Tx) error
	FetchAllFoods() ([]*models.Food, error)
	FetchMyFoodFromFoodsId(id *uuid.UUID) (models.MyFoods, error)
	FetchFoodFromFoodsName(FoodName string) ([]*models.Food, error)
	FetchFoodFromTypeOfFood(TypeOfFood string) ([]*models.Food, error)
	FetchFoodFromPrice(Price string) ([]*models.Food, error)
	FetchFoodWithFilter(args *sync.Map) ([]*models.Food, error)
}
