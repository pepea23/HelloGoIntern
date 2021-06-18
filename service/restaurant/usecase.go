package restaurant

import "github.com/HelloGoIntern/models"

type RestaurantUseCaseInterface interface {
	FetchAllRestaurants() ([]*models.Restaurant, error)
}