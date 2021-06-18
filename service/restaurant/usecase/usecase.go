package usecase

import (
	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/service/restaurant"
	"github.com/jmoiron/sqlx"
)


type restaurantUseCase struct {
	psqlRestaurantRepo restaurant.RestaurantRepositorynterface
	dbcon              *sqlx.DB
}

func NewRestaurantUsecase(repo restaurant.RestaurantRepositorynterface, dbcon *sqlx.DB) restaurant.RestaurantUseCaseInterface {
	return &restaurantUseCase{
		psqlRestaurantRepo: repo,
		dbcon:        dbcon,
	}
}


func (r restaurantUseCase) FetchAllRestaurants() ([]*models.Restaurant, error) {
	restaurants, err := r.psqlRestaurantRepo.FetchAllRestaurants()
	if err != nil {
		return nil, err
	}
	if len(restaurants) == 0 {
		return restaurants, err
	}

	for _, restaurant := range restaurants {
		Foods, err := r.psqlRestaurantRepo.FetchFoodFromRestaurantsId(restaurant.Id)
		if err != nil {
			return nil, err
		}
		restaurant.Foods = Foods
	}
	return restaurants, err
}