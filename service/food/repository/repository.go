package repository

import (
	"database/sql"
	"fmt"

	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/orm"
	"github.com/HelloGoIntern/service/food"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type foodRepository struct {
	db *sqlx.DB
}

func NewPsqlFoodRepository(dbcon *sqlx.DB) food.FoodRepositorynterface {
	return &foodRepository{
		db: dbcon,
	}
}

func (f foodRepository) CreateFood(food *models.Food, tx *sql.Tx) error {
	sql := `INSERT INTO food(id,name) VALUES($1::UUID,$2::TEXT)`

	stmt, err := tx.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		food.Id,
		food.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (f foodRepository) CreateMyFood(myFood *models.MyFood, tx *sql.Tx) error {
	sql := `INSERT INTO my_food(id,food_id,my) VALUES($1::UUID,$2::UUID,$3::TEXT)`

	stmt, err := tx.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		myFood.Id,
		myFood.FoodId,
		myFood.My,
	)

	if err != nil {
		return err
	}

	return nil
}

func (f foodRepository) FetchAllFoods() ([]*models.Food, error) {
	sql := `SELECT * FROM food`

	rows, err := f.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return f.orm(rows)
}

func (f foodRepository) FetchMyFoodFromFoodsId(id *uuid.UUID) (models.MyFoods, error) {
	sql := fmt.Sprintf(`SELECT * FROM my_food WHERE food_id='%s'`, id.String())

	rows, err := f.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return f.ormMyFood(rows)
}

func (f foodRepository) ormMyFood(rows *sqlx.Rows) (models.MyFoods, error) {
	var foods = make([]*models.MyFood, 0)

	for rows.Next() {
		var food = new(models.MyFood)
		food, err := orm.OrmMyFood(food, rows)
		if err != nil {
			return nil, err
		}
		if food != nil {
			foods = append(foods, food)
		}
	}

	return foods, nil
}

func (f foodRepository) orm(rows *sqlx.Rows) ([]*models.Food, error) {
	var foods = make([]*models.Food, 0)

	for rows.Next() {
		var food = new(models.Food)
		food, err := orm.OrmFood(food, rows)
		if err != nil {
			return nil, err
		}
		if food != nil {
			foods = append(foods, food)
		}
	}

	return foods, nil
}
