package restaurant

import (
	"sync"

	"github.com/HelloGoIntern/models"
	"github.com/gofrs/uuid"
	
)


type RestaurantRepositorynterface interface {
	FetchAllRestaurants() ([]*models.Restaurant, error)
	FetchFoodFromRestaurantsId(id *uuid.UUID) (models.Foods, error)
	FetchIdRestaurantsFromName(s string) ([]*models.Restaurant, error)
	FetchFoodInRestaurantWithFilter(args *sync.Map) ([]*models.Food, error)
}