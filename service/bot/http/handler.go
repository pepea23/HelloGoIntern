package http

import (
	"log"
	"strings"
	"sync"

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

func (b botHandler) GetAllMenu(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	foodName, err := b.botUs.GetAllFood()
	
	var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error // something")
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, foodName)
	}

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

func (b botHandler) RandomMenu(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	foodName, err := b.botUs.RandomFood()
	var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "error // something")
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, foodName)
	}

	bot.Send(msg)
}

func (b botHandler) Filter(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := update.Message.Text
	log.Print(text)
	s := strings.Split(text, " ")
	args := new(sync.Map)
	if len(s) == 1 {
		args.Store("food_name", s[0])
	} 
	if len(s) == 2 {
		args.Store("food_type", s[1])
	}
	if len(s) == 3 {
		args.Store("food_price", s[2])
	}
	
	log.Print(s)
	foodName, err := b.botUs.FilterFoods(args)
	
	var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, foodName)
	}

	bot.Send(msg)
}

func (b botHandler) Test(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := update.Message.Text
	log.Print(text)
	s := strings.Split(text, " ")
	log.Print(s)
}

