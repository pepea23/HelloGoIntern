package http

import (
	"log"

	"github.com/HelloGoIntern/middleware"
	"github.com/HelloGoIntern/service/bot"
	"github.com/labstack/echo/v4"
)

type botHandler struct {
	botUs bot.BOTUseCaseInterface
}

func NewBOTHandler(e *echo.Echo, middL *middleware.GoMiddleware, us bot.BOTUseCaseInterface) {
	botHandler := botHandler{
		botUs: us,
	}
	log.Print(botHandler)
}
