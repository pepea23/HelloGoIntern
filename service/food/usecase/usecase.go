package usecase

import (
	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/service/food"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type foodUseCase struct {
	psqlFoodRepo food.FoodRepositorynterface
	dbcon        *sqlx.DB
}

func NewFoodUsecase(repo food.FoodRepositorynterface, dbcon *sqlx.DB) food.FoodUseCaseInterface {
	return &foodUseCase{
		psqlFoodRepo: repo,
		dbcon:        dbcon,
	}
}

func (f foodUseCase) CreateFood(food *models.Food) error {
	tx, err := f.dbcon.Begin()
	if err != nil {
		return err
	}

	if err = f.psqlFoodRepo.CreateFood(food, tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (f foodUseCase) FetchAllFoods() ([]*models.Food, error) {
	foods, err := f.psqlFoodRepo.FetchAllFoods()
	if err != nil {
		return nil, err
	}
	if len(foods) == 0 {
		return foods, err
	}

	return foods, err
}

func (f foodUseCase) FetchFoodFromFoodsName(FoodName string) ([]*models.Food, error) {
	foods, err := f.psqlFoodRepo.FetchFoodFromFoodsName(FoodName)
	if err != nil {
		return nil, err
	}
	if len(foods) == 0 {
		return foods, err
	}

	return foods, err
}

func (f foodUseCase) FetchFoodFromTypeOfFood(TypeOfFood string) ([]*models.Food, error) {
	foods, err := f.psqlFoodRepo.FetchFoodFromTypeOfFood(TypeOfFood)
	if err != nil {
		return nil, err
	}
	if len(foods) == 0 {
		return foods, err
	}

	return foods, err
}

func (f foodUseCase) FetchFoodFromPrice(Price string) ([]*models.Food, error) {
	foods, err := f.psqlFoodRepo.FetchFoodFromPrice(Price)
	if err != nil {
		return nil, err
	}
	if len(foods) == 0 {
		return foods, err
	}

	return foods, err
}
func (f foodUseCase) Deletefood(id uuid.UUID) error {
	return f.psqlFoodRepo.Deletefood(id)
}
