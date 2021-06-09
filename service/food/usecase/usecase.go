package usecase

import (
	"database/sql"

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

	if err = f.createMyFoodsWithFoodID(food.Id, food.MyFoods, tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (f foodUseCase) createMyFoodsWithFoodID(foodId *uuid.UUID, myFoods models.MyFoods, tx *sql.Tx) error {
	for _, myFood := range myFoods {
		myFood.FoodId = foodId
		myFood.GenarateUUID()
		if err := f.psqlFoodRepo.CreateMyFood(myFood, tx); err != nil {
			return err
		}
	}
	return nil
}

func (f foodUseCase) FetchAllFoods() ([]*models.Food, error) {
	foods, err := f.psqlFoodRepo.FetchAllFoods()
	if err != nil {
		return nil, err
	}
	if len(foods) == 0 {
		return foods, err
	}

	for _, food := range foods {
		myFoods, err := f.psqlFoodRepo.FetchMyFoodFromFoodsId(food.Id)
		if err != nil {
			return nil, err
		}
		food.MyFoods = myFoods
	}
	return foods, err
}
