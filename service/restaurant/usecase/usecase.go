package usecase

import (
	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/service/restaurant"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type restaurantUseCase struct {
	psqlRestaurantRepo restaurant.RestaurantRepositorynterface
	dbcon              *sqlx.DB
}

func NewRestaurantUsecase(repo restaurant.RestaurantRepositorynterface, dbcon *sqlx.DB) restaurant.RestaurantUseCaseInterface {
	return &restaurantUseCase{
		psqlRestaurantRepo: repo,
		dbcon:              dbcon,
	}
}

func (r restaurantUseCase) CreateRestaurant(restaurant *models.Restaurant) error {
	tx, err := r.dbcon.Begin()
	if err != nil {
		return err
	}

	if err = r.psqlRestaurantRepo.CreateRestaurant(restaurant, tx); err != nil {
		return err
	}

	return tx.Commit()
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

func (r restaurantUseCase) DeleteRestaurant(id uuid.UUID) error {
	return r.psqlRestaurantRepo.DeleteRestaurant(id)
}
