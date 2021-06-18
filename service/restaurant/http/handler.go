package http

import (
	"net/http"
	"github.com/HelloGoIntern/middleware"
	
	"github.com/HelloGoIntern/service/restaurant"
	"github.com/labstack/echo/v4"
)


type restaurantHandler struct {
	restaurantUs restaurant.RestaurantUseCaseInterface
}

func NewRestaurantHandler(e *echo.Echo, middL *middleware.GoMiddleware, us restaurant.RestaurantUseCaseInterface) {
	restaurantHandler := restaurantHandler{
		restaurantUs: us,
	}
//	e.POST("/restaurants", restaurantHandler, middL.InputForm)
	e.GET("/restaurants", restaurantHandler.FetchAll)
}

func (r *restaurantHandler) FetchAll(e echo.Context) error {

	restaurants, err := r.restaurantUs.FetchAllRestaurants()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "error")
	}

	var response = map[string]interface{}{
		"restaurants": restaurants,
	}

	return e.JSON(http.StatusOK, response)
}