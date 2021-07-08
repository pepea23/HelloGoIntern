package restaurant

import (
	"github.com/HelloGoIntern/models"
	"github.com/gofrs/uuid"
)

type RestaurantUseCaseInterface interface {
	CreateRestaurant(restaurant *models.Restaurant) error
	FetchAllRestaurants() ([]*models.Restaurant, error)
	DeleteRestaurant(id uuid.UUID) error
}
