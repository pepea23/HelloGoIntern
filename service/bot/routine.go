package bot

import (
	"log"
	"strings"
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
			case "!รายการอาหาร":
				handle.GetAllMenu(bot, update)

			case "!คำสั่ง":
				handle.ABC(bot, update)

			case "!สุ่มอาหาร":
				handle.RandomMenu(bot, update)

			default:

			}

			if strings.Contains(text, "!ค้นหาร้านอาหาร") {
				handle.FilterRestaurant(bot, update)
			}

			if strings.Contains(text, "!ค้นหาอาหาร") {
				handle.FilterFood(bot, update)
			}

			time.Sleep(3000)
		}
	}()
}
