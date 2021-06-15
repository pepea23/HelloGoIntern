package usecase

import (
	"github.com/HelloGoIntern/service/bot"
	"github.com/HelloGoIntern/service/food"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type botAPIUsercase struct {
	botapi *tgbotapi.BotAPI
	foodUs food.FoodUseCaseInterface
}

func NewBotAPIUsecase(botapi *tgbotapi.BotAPI, foodUs food.FoodUseCaseInterface) bot.BOTUseCaseInterface {
	return &botAPIUsercase{
		botapi: botapi,
		foodUs: foodUs,
	}
}

func (b botAPIUsercase) GetSomeThing() (string, error) {
	food, err := b.foodUs.FetchAllFoods()
	if err != nil {
		return "", err
	}
	return food[0].Name, nil
}
