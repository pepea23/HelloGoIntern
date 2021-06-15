package http

import (
	"github.com/HelloGoIntern/middleware"
	mybot "github.com/HelloGoIntern/service/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/labstack/echo/v4"
)

type botHandler struct {
	botUs mybot.BOTUseCaseInterface
}

func NewBOTHandler(e *echo.Echo, middL *middleware.GoMiddleware, us mybot.BOTUseCaseInterface) mybot.BOTHandlerInterface {
	return botHandler{
		botUs: us,
	}
}

func (b botHandler) TestCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	bot.Send(msg)
}

func (b botHandler) ABC(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	foodName, err := b.botUs.GetSomeThing()
	var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error // something")
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, foodName)
	}

	bot.Send(msg)
}
