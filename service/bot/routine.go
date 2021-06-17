package bot

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func RoutineBot(bot *tgbotapi.BotAPI, handle BOTHandlerInterface) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	log.Print(updates, err)

	
	go func() {
		for update := range updates {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}

			text := update.Message.Text
			switch text {
			case "อาหารทั้งหมด":
				handle.GetAllMenu(bot, update)
				
			case "aaa":
				handle.ABC(bot, update)

			case "สุ่มอาหาร":
				handle.RandomMenu(bot, update)

			case "!ค้นหาอาหารด้วยชื่อ":
				handle.Test(bot, update)

			default:
				handle.Filter(bot, update)
			}


			time.Sleep(3000)
		}
	}()
}
