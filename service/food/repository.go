package food

import (
	"database/sql"

	"github.com/HelloGoIntern/models"
	"github.com/gofrs/uuid"
)

type FoodRepositorynterface interface {
	CreateFood(food *models.Food, tx *sql.Tx) error
	CreateMyFood(myFood *models.MyFood, tx *sql.Tx) error
	FetchAllFoods() ([]*models.Food, error)
	FetchMyFoodFromFoodsId(id *uuid.UUID) (models.MyFoods, error)
}
