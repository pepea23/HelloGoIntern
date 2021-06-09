package repository

import (
	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/orm"
	"github.com/HelloGoIntern/service/food"
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

func (f foodRepository) CreateFood(food *models.Food) error {
	tx, err := f.db.Begin()
	if err != nil {
		return err
	}
	sql := `INSERT INTO food(name,quntity) VALUES($1::TEXT, $2::numeric)`

	stmt, err := tx.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		food.Name,
		food.Quntity,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (f foodRepository) FetchAllFoods() ([]*models.Food, error) {
	//add comment
	sql := `SELECT * FROM food`

	rows, err := f.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return f.orm(rows)
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
