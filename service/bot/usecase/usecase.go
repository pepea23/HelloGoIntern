package usecase

import (
	"github.com/HelloGoIntern/service/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type botAPIUsercase struct {
	botapi *tgbotapi.BotAPI
}

func NewBotAPIUsecase(botapi *tgbotapi.BotAPI) bot.BOTUseCaseInterface {
	return &botAPIUsercase{
		botapi: botapi,
	}
}
