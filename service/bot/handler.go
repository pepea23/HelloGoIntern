package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BOTHandlerInterface interface {
	GetAllMenu(bot *tgbotapi.BotAPI, msg tgbotapi.Update)
	ABC(bot *tgbotapi.BotAPI, msg tgbotapi.Update)
	RandomMenu(bot *tgbotapi.BotAPI, msg tgbotapi.Update)
	FilterFood(bot *tgbotapi.BotAPI, msg tgbotapi.Update)
	FilterRestaurant(bot *tgbotapi.BotAPI, msg tgbotapi.Update)
}
