package bot

import "sync"

type BOTUseCaseInterface interface {
	//	CreateFood() (string,error)
	GetSomeThing() (string, error)
	GetAllFood() (string, error)
	RandomFood() (string, error)
	GetAllFoodInRestaurant(s string) (string, error)
	GetAllRestaurant() (string, error)
	FilterFoods(args *sync.Map) (string, error)
	FilterFoodsInRestaurants(args *sync.Map) (string, error)
}
