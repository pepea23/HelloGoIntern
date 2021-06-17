package bot

import "sync"

type BOTUseCaseInterface interface {
	GetSomeThing() (string, error)
	GetAllFood() (string, error)
	RandomFood() (string, error)
	FilterFood(s string) (string, error)
	FilterFoods(args *sync.Map) (string, error)
}
