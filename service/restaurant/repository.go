package restaurant

import (
	"database/sql"
	"sync"

	"github.com/HelloGoIntern/models"
	"github.com/gofrs/uuid"
)

type RestaurantRepositorynterface interface {
	CreateRestaurant(restaurant *models.Restaurant, tx *sql.Tx) error
	FetchAllRestaurants() ([]*models.Restaurant, error)
	FetchFoodFromRestaurantsId(id *uuid.UUID) (models.Foods, error)
	FetchIdRestaurantsFromName(s string) ([]*models.Restaurant, error)
	FetchFoodInRestaurantWithFilter(args *sync.Map) ([]*models.Food, error)
	DeleteRestaurant(id uuid.UUID) error
}
