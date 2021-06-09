package main

import (
	"log"
	"net/http"

	_conf "github.com/HelloGoIntern/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	myMiddL "github.com/HelloGoIntern/middleware"

	_food_handler "github.com/HelloGoIntern/service/food/http"
	_food_repository "github.com/HelloGoIntern/service/food/repository"
	_food_usecase "github.com/HelloGoIntern/service/food/usecase"

	_user_handler "github.com/HelloGoIntern/service/user/http"
	_user_repository "github.com/HelloGoIntern/service/user/repository"
	_user_usecase "github.com/HelloGoIntern/service/user/usecase"
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

func main() {
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

	userUs := _user_usecase.NewUserUsecase(userRepo)
	foodUs := _food_usecase.NewFoodUsecase(foodRepo)
	/* Inject Handler */

	_user_handler.NewUserHandler(e, middL, userUs)
	_food_handler.NewFoodHandler(e, middL, foodUs)

	port := ":" + _conf.GetEnv("PORT", "3000")
	e.Logger.Fatal(e.Start(port))
}
