package http

import (
	"net/http"

	"github.com/HelloGoIntern/middleware"
	"github.com/HelloGoIntern/models"
	"github.com/gofrs/uuid"

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
	e.POST("/restaurants", restaurantHandler.CreateRestaurant, middL.InputForm)
	e.GET("/restaurants", restaurantHandler.FetchAll)
	e.DELETE("/restaurants/:id", restaurantHandler.DeleteRestaurant)
}

func (r *restaurantHandler) CreateRestaurant(e echo.Context) error {
	var params = e.Get("params")
	var restaurant = models.NewRestaurantWithParam(params.(map[string]interface{}))
	restaurant.GenarateUUID()

	err := r.restaurantUs.CreateRestaurant(restaurant)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "error")
	}

	var response = map[string]interface{}{
		"message": "ok",
	}

	return e.JSON(http.StatusOK, response)
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

func (r *restaurantHandler) DeleteRestaurant(e echo.Context) error {
	id := e.Param("id")
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil
	}
	r.restaurantUs.DeleteRestaurant(uid)
	var response = map[string]interface{}{
		"message": "deleted",
	}
	return e.JSON(http.StatusOK, response)
}
