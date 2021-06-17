package usecase

import (
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/HelloGoIntern/models"
	"github.com/HelloGoIntern/service/bot"
	"github.com/HelloGoIntern/service/food"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type botAPIUsercase struct {
	botapi *tgbotapi.BotAPI
	foodRepo food.FoodRepositorynterface
	
}

func NewBotAPIUsecase(botapi *tgbotapi.BotAPI, foodRepo food.FoodRepositorynterface) bot.BOTUseCaseInterface {
	return &botAPIUsercase{
		botapi: botapi,
		foodRepo: foodRepo ,
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
		s = append(s,"|  อาหารทั้งหมด  | \n ———————")
	for i := 0; i < len(food); i++ {
		
		s = append(s,"|  " +food[i].FoodName + "             |")
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
		s = append(s,"|  อาหารที่สุ่มได้คือ!!  | \n —————————")
		s = append(s,"|  " +food[ran].FoodName + "                    |")
		stringArray := s
  		String := strings.Join(stringArray,"\n")
	return String , nil
}

func (b botAPIUsercase) FilterFood(s string) (string, error) {
//	foodP, errP := b.foodUs.FetchFoodFromPrice(s)
//	foodN, errN := b.foodUs.FetchFoodFromFoodsName(s) 
	foodT, errT := b.foodRepo.FetchFoodFromTypeOfFood(s)
	log.Print(foodT)
	if errT != nil {
		return "", errT
	}

 /* 	if s == foodN[0].FoodName{
		if errN != nil {
			return "", errN
		}
		return b.FilterFoodName(foodN) 
	
	} else */ if len(foodT) == 0 {
		return "nothing" , nil
		

	/* } else  if s == foodP[0].Price {
		if errP != nil {
			return "", errP
		}
		return b.FilterFoodPrice(foodP) */
	}  
		return b.FilterFoodType(foodT)
		
	
		
	
	
	
	
	  /* if s == foodN[0].FoodName {
		if errN != nil {
			return "", errN
		}
		if len(foodN) == 0 {
			return "nothing", nil
		}
		var ar []string
		for i := 0; i < len(foodN); i++ {
		
			ar = append(ar, foodN[i].FoodName)
			log.Print(ar)
		}
		stringArray := ar
		  String := strings.Join(stringArray,"\n")
		return String , nil


	} else  if s == foodT[0].TypeOfFood {
		if errT != nil {
			return "", errT
		}
		if len(foodT) == 0 {
			return "nothing", nil
		}
		var ar2 []string
		ar2 = append(ar2,"|  อาหารประเภท " + s +"   | \n —————————")
		for i := 0; i < len(foodT); i++ {
		
			ar2 = append(ar2,"|  " + foodT[i].FoodName + "                    |")
			log.Print(ar2)
		}
		stringArray2 := ar2
		  String2 := strings.Join(stringArray2,"\n")
		return String2 , nil

	}  else  if s == foodP[0].Price {
		if errP != nil {
			return "", errP
		}
		if len(foodP) == 0 {
			return "nothing", nil
		}
		var ar3 []string
		ar3 = append(ar3,"|  อาหารในราคา " + s + " บาท" +"   | \n ———————————")
		for i := 0; i < len(foodP); i++ {
		
			ar3 = append(ar3,"|  " + foodP[i].FoodName + "                    |")
			log.Print(ar3)
		}
		stringArray3 := ar3
		  String3 := strings.Join(stringArray3,"\n")
		return String3 , nil

	}  
		return "nothing", nil */
	
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

	return "", nil
}

func (b botAPIUsercase) FilterFoodName(f []*models.Food) (string, error) {
	if len(f) == 0 {
		return "nothing", nil
	}
	var ar []string
	for i := 0; i < len(f); i++ {
	
		ar = append(ar, f[i].FoodName)
		log.Print(ar)
	}
	stringArray := ar
	  String := strings.Join(stringArray,"\n")
	return String , nil

}

func (b botAPIUsercase) FilterFoodType(f []*models.Food) (string, error) {
	
	if len(f) == 0 {
		return "nothing", nil
	}
	var ar2 []string
	ar2 = append(ar2,"|  อาหารประเภท " + f[0].TypeOfFood +"   | \n —————————")
	for i := 0; i < len(f); i++ {
	
		ar2 = append(ar2,"|  " + f[i].FoodName + "                    |")
		log.Print(ar2)
	}
	stringArray2 := ar2
	  String2 := strings.Join(stringArray2,"\n")
	return String2 , nil

}

func (b botAPIUsercase) FilterFoodPrice(f []*models.Food) (string, error) {
	
	if len(f) == 0 {
		return "nothing", nil
	}
	var ar3 []string
	ar3 = append(ar3,"|  อาหารในราคา " + f[0].Price + " บาท" +"   | \n ———————————")
	for i := 0; i < len(f); i++ {
	
		ar3 = append(ar3,"|  " + f[i].FoodName + "                    |")
		log.Print(ar3)
	}
	stringArray3 := ar3
	  String3 := strings.Join(stringArray3,"\n")
	return String3 , nil


}