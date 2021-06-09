package http

import (
	"net/http"

	"github.com/HelloGoIntern/middleware"
	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/service/food"
	"github.com/labstack/echo/v4"
)

type foodHandler struct {
	foodUs food.FoodUseCaseInterface
}

func NewFoodHandler(e *echo.Echo, middL *middleware.GoMiddleware, us food.FoodUseCaseInterface) {
	foodHandler := foodHandler{
		foodUs: us,
	}
	e.POST("/foods", foodHandler.CreateFood, middL.InputForm)
	e.GET("/foods", foodHandler.FetchAll)
}

func (f *foodHandler) CreateFood(e echo.Context) error {
	var params = e.Get("params")
	var food = models.NewFoodWithParam(params.(map[string]interface{}))

	err := f.foodUs.CreateFood(food)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "error")
	}

	var response = map[string]interface{}{
		"message": "ok",
	}

	return e.JSON(http.StatusOK, response)
}

func (f *foodHandler) FetchAll(e echo.Context) error {

	foods, err := f.foodUs.FetchAllFoods()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "error")
	}

	var response = map[string]interface{}{
		"foods": foods,
	}

	return e.JSON(http.StatusOK, response)
}
