package repository

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/orm"
	"github.com/HelloGoIntern/service/restaurant"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)


type restaurantRepository struct {
	db *sqlx.DB
}

func NewPsqlRestaurantRepository(dbcon *sqlx.DB) restaurant.RestaurantRepositorynterface {
	return &restaurantRepository{
		db: dbcon,
	}
}

func (r restaurantRepository) FetchAllRestaurants() ([]*models.Restaurant, error) {
	sql := `SELECT * FROM restaurant`

	rows, err := r.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.orm(rows)
}

func (r restaurantRepository) orm(rows *sqlx.Rows) ([]*models.Restaurant, error) {
	var restaurants = make([]*models.Restaurant, 0)

	for rows.Next() {
		var restaurant = new(models.Restaurant)
		restaurant, err := orm.OrmRestaurant(restaurant, rows)
		if err != nil {
			return nil, err
		}
		if restaurant != nil {
			restaurants = append(restaurants, restaurant)
		}
	}

	return restaurants, nil
}

func (r restaurantRepository) ormFood(rows *sqlx.Rows) (models.Foods, error) {
	var restaurants = make([]*models.Food, 0)

	for rows.Next() {
		var restaurant = new(models.Food)
		restaurant, err := orm.OrmFood(restaurant, rows)
		if err != nil {
			return nil, err
		}
		if restaurant != nil {
			restaurants = append(restaurants, restaurant)
		}
	}

	return restaurants, nil
}

func (r restaurantRepository) FetchFoodFromRestaurantsId(id *uuid.UUID) (models.Foods, error) {
	sql := fmt.Sprintf(`SELECT * FROM food WHERE restaurant_id='%s'`, id.String())

	rows, err := r.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.ormFood(rows)
}

func (r restaurantRepository) FetchIdRestaurantsFromName(s string) ([]*models.Restaurant, error) {
	sql := fmt.Sprintf(`SELECT * FROM restaurant WHERE restaurant_name='%s'`, s)

	rows, err := r.db.Queryx(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return r.orm(rows)
}

 func (r restaurantRepository) FetchFoodInRestaurantWithFilter(args *sync.Map) ([]*models.Food, error) {
	var wheresomthing []string 


	if restaurantName, ok := args.Load("restaurant_name"); ok {
		idRestaurant, err := r.FetchIdRestaurantsFromName(fmt.Sprintf(`%s`, restaurantName))
		if err != nil {
			return nil, err
		}
		foods, err := r.FetchFoodFromRestaurantsId(idRestaurant[0].Id)
		if getOne, ok := args.Load("get_one"); ok { 
			for i := 0; i < len(foods); i++ {
				if foods[i].FoodName == getOne {
					wheresomthing = append(wheresomthing, fmt.Sprintf(`food_name='%s'`, foods[i].FoodName))
					log.Print(wheresomthing)
				}	
				if foods[i].TypeOfFood == getOne {
					wheresomthing = append(wheresomthing, fmt.Sprintf(`type_of_food='%s'`, foods[i].TypeOfFood))
					log.Print(wheresomthing)
				}	
				if foods[i].Price == getOne {
					wheresomthing = append(wheresomthing, fmt.Sprintf(`price='%s'`, foods[i].Price))
					log.Print(wheresomthing)
				} 
				
			} 
			
		
		}
		for i := 0; i < len(foods); i++ {
		if foodName, ok := args.Load("food_name"); ok { 
			if foods[i].FoodName == foodName {
				wheresomthing = append(wheresomthing, fmt.Sprintf(`food_name='%s'`, foodName))
			log.Print(wheresomthing)
			}	
		}
		if foodType, ok := args.Load("food_type"); ok { 
			wheresomthing = append(wheresomthing, fmt.Sprintf(`type_of_food='%s'`, foodType))
			log.Print(wheresomthing)
		}
		if foodPrice, ok := args.Load("food_price"); ok { 
			wheresomthing = append(wheresomthing, fmt.Sprintf(`price='%s'`, foodPrice))
			log.Print(wheresomthing)
		}
	}
		var where string
		log.Print(wheresomthing)
		if len(wheresomthing) != 0 {
			where = "WHERE " + strings.Join(wheresomthing," AND ")
			
		}

		if where == "" {
			return nil, nil
		}
		sql := fmt.Sprintf(`SELECT * FROM food %s`, where)
		log.Print(sql)
	
		rows, err := r.db.Queryx(sql)
		if err != nil {
			return nil, err
		}
		
		defer rows.Close()
		return r.ormFood(rows)
		
	}

	return nil, nil
	
} 





