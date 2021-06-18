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
	var ar []string
	ar = append(ar," \n —————————————————————")
	ar = append(ar,"|                               คำสั่งทั้งหมด   \n —————————————————————")
	ar = append(ar,"| !คำสั่ง   \n ")
	ar = append(ar,"| !สุ่มอาหาร   \n ")
	ar = append(ar,"| !ค้นหาร้านอาหาร (สามารถ filter [ชื่อร้าน ชื่ออาหาร ประเภท ราคา] ได้)   \n")
	ar = append(ar,"| !ค้นหาอาหาร (สามารถ filter [ชื่อ ประเภท ราคา] ได้)  \n")
	stringArray := ar
	command := strings.Join(stringArray,"\n")
	
	var msg tgbotapi.MessageConfig
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, command)
	

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

func (b botHandler) FilterFood(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := update.Message.Text
	log.Print(text)
	s := strings.Split(text, " ")
	args := new(sync.Map)

	if len(s) == 1   {
		foodName, err := b.botUs.GetAllFood()
	
	var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, foodName)
	}

	bot.Send(msg)
	} 


	if len(s) == 2   {
		args.Store("get_one", s[1])
	} 
	if len(s) > 2  {
		args.Store("food_name", s[1])
	} 
	if len(s) == 3 || len(s) > 3 {
		args.Store("food_type", s[2])
	}
	if len(s) == 4 || len(s) > 4 {
		args.Store("food_price", s[3])
	}
	if len(s) > 1 {
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
	
}

func (b botHandler) FilterRestaurant(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := update.Message.Text
	log.Print(text)
	s := strings.Split(text, " ")
	args := new(sync.Map)
	if len(s) == 1   {
		restaurantName, err := b.botUs.GetAllRestaurant()
		var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, restaurantName)
	}

	bot.Send(msg)
	}
	if len(s) == 2   {
		foodName, err := b.botUs.GetAllFoodInRestaurant(s[1])
	
	var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, foodName)
	}

	bot.Send(msg)
	} 
		

	if len(s) == 3   {
		args.Store("get_one", s[2])
		args.Store("restaurant_name", s[1])
	} 
	if len(s) > 3  {
		args.Store("food_name", s[2])
		args.Store("restaurant_name", s[1])
	} 
	if len(s) == 4 || len(s) > 4 {
		args.Store("food_type", s[3])
		args.Store("restaurant_name", s[1])
	}
	if len(s) == 5 || len(s) > 5 {
		args.Store("food_price", s[4])
		args.Store("restaurant_name", s[1])
	}

	if len(s) > 2   {
		foodName, err := b.botUs.FilterFoodsInRestaurants(args)
	
	var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, foodName)
	}

	bot.Send(msg)
	} 
	
	
/* 	log.Print(s)
	foodName, err := b.botUs.FilterFoods(args)
	
	var msg tgbotapi.MessageConfig
	if err != nil {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, foodName)
	}

	bot.Send(msg) */
	
	
}



