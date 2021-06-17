package main

import (
	"log"
	"net/http"

	_conf "github.com/HelloGoIntern/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	myMiddL "github.com/HelloGoIntern/middleware"

	_my_bot_ "github.com/HelloGoIntern/service/bot"
	_bot_handler "github.com/HelloGoIntern/service/bot/http"
	_bot_usecase "github.com/HelloGoIntern/service/bot/usecase"

	_food_handler "github.com/HelloGoIntern/service/food/http"
	_food_repository "github.com/HelloGoIntern/service/food/repository"
	_food_usecase "github.com/HelloGoIntern/service/food/usecase"

	_user_handler "github.com/HelloGoIntern/service/user/http"
	_user_repository "github.com/HelloGoIntern/service/user/repository"
	_user_usecase "github.com/HelloGoIntern/service/user/usecase"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
)

func sqlDB() *sqlx.DB {
	var connstr = _conf.GetEnv("PSQL_DATABASE_URL", "postgres://postgres:postgres@psql_db:5432/app_example?sslmode=disable")
	db, err := _conf.NewPsqlConnection(connstr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func connectBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(_conf.GetEnv("TOKEN_BOT", ""))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func main() {
	//set up bot and routine
	bot := connectBot()

	psqlDB := sqlDB()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	middL := myMiddL.InitMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	/* Inject Repository */

	userRepo := _user_repository.NewPsqlUserRepository(psqlDB)
	foodRepo := _food_repository.NewPsqlFoodRepository(psqlDB)
	/* Inject Usecase */

	foodUs := _food_usecase.NewFoodUsecase(foodRepo, psqlDB)
	botUs := _bot_usecase.NewBotAPIUsecase(bot, foodRepo)
	userUs := _user_usecase.NewUserUsecase(userRepo)
	/* Inject Handler */

	handle := _bot_handler.NewBOTHandler(e, middL, botUs)
	_user_handler.NewUserHandler(e, middL, userUs)
	_food_handler.NewFoodHandler(e, middL, foodUs)

	_my_bot_.RoutineBot(bot, handle)
	port := ":" + _conf.GetEnv("PORT", "3000")
	e.Logger.Fatal(e.Start(port))
}
