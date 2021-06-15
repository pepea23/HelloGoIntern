package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type BOTHandlerInterface interface {
	TestCallback(bot *tgbotapi.BotAPI, msg tgbotapi.Update)
	ABC(bot *tgbotapi.BotAPI, msg tgbotapi.Update)
}
