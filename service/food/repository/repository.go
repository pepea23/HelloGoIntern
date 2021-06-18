package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

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
		food.FoodName,
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

func (f foodRepository) FetchFoodFromFoodsName(FoodName string) ([]*models.Food, error) {
	sql := fmt.Sprintf(`SELECT * FROM food WHERE food_name='%s'`, FoodName)

	rows, err := f.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	return f.orm(rows)
}

func (f foodRepository) FetchFoodFromTypeOfFood(TypeOfFood string) ([]*models.Food, error) {
	sql := fmt.Sprintf(`SELECT * FROM food WHERE type_of_food='%s'`, TypeOfFood)

	rows, err := f.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	return f.orm(rows)
}

func (f foodRepository) FetchFoodFromPrice(Price string) ([]*models.Food, error) {
	sql := fmt.Sprintf(`SELECT * FROM food WHERE price='%s'`, Price)

	rows, err := f.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	return f.orm(rows)
}


func (f foodRepository) FetchFoodWithFilter(args *sync.Map) ([]*models.Food, error) {
	var wheresomthing []string 

	if getOne, ok := args.Load("get_one"); ok { 
		
		foodP, errP := f.FetchFoodFromPrice(fmt.Sprintf(`%s`, getOne))
		foodN, errN := f.FetchFoodFromFoodsName(fmt.Sprintf(`%s`, getOne)) 
		foodT, errT := f.FetchFoodFromTypeOfFood(fmt.Sprintf(`%s`, getOne))
		if errT != nil {
			return nil, errT
		}
		if errP != nil {
			return nil, errP
		}
		if errN != nil {
			return nil, errN
		}
		if len(foodP) | len(foodN) | len(foodT) == 0 {
			return nil, nil
		}
		if len(foodP) != 0 {
			
			wheresomthing = append(wheresomthing, fmt.Sprintf(`price='%s'`, foodP[0].Price))
		log.Print(wheresomthing)
		}
		if len(foodN) != 0 {
			wheresomthing = append(wheresomthing, fmt.Sprintf(`food_name='%s'`, foodN[0].FoodName))
		log.Print(wheresomthing)
		}
		if len(foodT) != 0 {
			wheresomthing = append(wheresomthing, fmt.Sprintf(`type_of_food='%s'`, foodT[0].TypeOfFood))
		log.Print(wheresomthing)
		} 
		
		
		
	}
	if foodName, ok := args.Load("food_name"); ok { 
		wheresomthing = append(wheresomthing, fmt.Sprintf(`food_name='%s'`, foodName))
		log.Print(wheresomthing)
	}
	if foodType, ok := args.Load("food_type"); ok { 
		wheresomthing = append(wheresomthing, fmt.Sprintf(`type_of_food='%s'`, foodType))
		log.Print(wheresomthing)
	}
	if foodPrice, ok := args.Load("food_price"); ok { 
		wheresomthing = append(wheresomthing, fmt.Sprintf(`price='%s'`, foodPrice))
		log.Print(wheresomthing)
	}
	var where string
	log.Print(wheresomthing)
	if len(wheresomthing) != 0 {
		where = "WHERE " + strings.Join(wheresomthing," AND ")
	}
 	
	sql := fmt.Sprintf(`SELECT * FROM food %s`, where)
	log.Print(sql)

	rows, err := f.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	return f.orm(rows)
	
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
