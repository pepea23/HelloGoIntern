package usecase

import (
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/HelloGoIntern/service/bot"
	"github.com/HelloGoIntern/service/food"
	"github.com/HelloGoIntern/service/restaurant"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type botAPIUsercase struct {
	botapi *tgbotapi.BotAPI
	foodRepo food.FoodRepositorynterface
	restaurantRepo restaurant.RestaurantRepositorynterface
	
}

func NewBotAPIUsecase(botapi *tgbotapi.BotAPI, foodRepo food.FoodRepositorynterface, restaurantRepo restaurant.RestaurantRepositorynterface) bot.BOTUseCaseInterface {
	return &botAPIUsercase{
		botapi: botapi,
		foodRepo: foodRepo ,
		restaurantRepo: restaurantRepo ,
	}
}

func (b botAPIUsercase) GetSomeThing() (string, error) {
	food, err := b.foodRepo.FetchAllFoods()
	if err != nil {
		return "", err
	}
	return food[0].FoodName, nil
}

func (b botAPIUsercase) GetAllFood() (string, error) {

	food, err := b.foodRepo.FetchAllFoods()
	if err != nil {
		return "", err
	}
	var s []string
	s = append(s," \n —————————————————————")
		s = append(s,"|                               อาหารทั้งหมด   \n —————————————————————")
	for i := 0; i < len(food); i++ {
		
		s = append(s,"|  " +food[i].FoodName )
		log.Print(s)
	}
	stringArray := s
  	String := strings.Join(stringArray,"\n")

	return String , nil
}

func (b botAPIUsercase) RandomFood() (string, error) {
	food, err := b.foodRepo.FetchAllFoods()
	if err != nil {
		return "", err
	}
	
	rand.Seed(time.Now().UnixNano())
    min := 0
    max := len(food)-1
	ran := rand.Intn(max - min + 1) + min
	var s []string
	s = append(s," \n —————————————————————")
		s = append(s,"|                               อาหารที่สุ่มได้คือ!!   \n —————————————————————")
		s = append(s,"|  " +food[ran].FoodName  )
		stringArray := s
  		String := strings.Join(stringArray,"\n")
	return String , nil
}


func (b botAPIUsercase) FilterFoods(args *sync.Map) (string, error) { 
	foods, err := b.foodRepo.FetchFoodWithFilter(args)
	log.Print(foods)
	if err != nil {
		return "", err
	}
	if len(foods) == 0 {
		return "nothing" , nil
	}
	var ar []string
	ar = append(ar," \n —————————————————————")
	ar = append(ar,"|                               รายการอาหารที่ค้นหาเจอ   \n —————————————————————")
	for i := 0; i < len(foods); i++ {
	
		ar = append(ar,"|  " + foods[i].FoodName )
		log.Print(ar)
	}
	stringArray := ar
	  String := strings.Join(stringArray,"\n")

	return String, nil
}

func (b botAPIUsercase) GetAllFoodInRestaurant(ss string) (string, error) {
	restaurant, err := b.restaurantRepo.FetchIdRestaurantsFromName(ss)
	if err != nil {
		return "", err
	}
	if len(restaurant) != 0 {
		food, err := b.restaurantRepo.FetchFoodFromRestaurantsId(restaurant[0].Id)
	if err != nil {
		return "", err
	}
	var s []string
	s = append(s," \n —————————————————————")
		s = append(s,"|  ร้าน : "+ restaurant[0].RestaurantName + "   \n —————————————————————")
		s = append(s,"|  เวลาเปิดปิด : "+ restaurant[0].OpenCloseTime + "   \n  —————————————————————")
		s = append(s,"|  สถานที่ตั้ง : "+ restaurant[0].Address + "   \n  —————————————————————")
		s = append(s,"|  เบอร์โทรติดต่อ : "+ restaurant[0].PhoneNumber + "   \n  —————————————————————")
		s = append(s,"|  อาหารที่มีใน"+ restaurant[0].RestaurantName + "   \n ")
	for i := 0; i < len(food); i++ {
		
		s = append(s,"|  " +food[i].FoodName + " " + food[i].Price + " บาท" + "             ")
		log.Print(s)
	}
	stringArray := s
  	String := strings.Join(stringArray,"\n")

	return String , nil

	}
	return "ไม่มีร้านนี้!" , nil
}

func (b botAPIUsercase) GetAllRestaurant() (string, error) {

	restaurant, err := b.restaurantRepo.FetchAllRestaurants()
	if err != nil {
		return "", err
	}
	var s []string
		s = append(s," \n —————————————————————")
		s = append(s,"|                               รายชื่อร้านอาหารทั้งหมด \n —————————————————————")
	for i := 0; i < len(restaurant); i++ {
		
		s = append(s,"|  " + restaurant[i].RestaurantName   )
		log.Print(s)
	}
	stringArray := s
  	String := strings.Join(stringArray,"\n")

	return String , nil
}

func (b botAPIUsercase) FilterFoodsInRestaurants(args *sync.Map) (string, error) { 
	foods, err := b.restaurantRepo.FetchFoodInRestaurantWithFilter(args)
	log.Print(foods)
	if err != nil {
		return "", err
	}
	if len(foods) == 0 {
		return "nothing" , nil
	}
	var ar []string
	ar = append(ar," \n —————————————————————")
	ar = append(ar,"|                               รายการอาหารที่ค้นหาเจอ   \n —————————————————————")
	for i := 0; i < len(foods); i++ {
	
		ar = append(ar,"|  " + foods[i].FoodName )
		log.Print(ar)
	}
	stringArray := ar
	  String := strings.Join(stringArray,"\n")

	return String, nil
}