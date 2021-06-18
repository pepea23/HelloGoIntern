package models

import (
	helperModel "git.innovasive.co.th/backend/models"
	"github.com/gofrs/uuid"
)

const FK_FIELD_MY_FOODS = "MyFoods"

type Food struct {
	TableName struct{}               `json:"-" db:"food"`
	Id        *uuid.UUID             `json:"id" db:"id" type:"uuid"`
	FoodName   string                `json:"food_name" db:"food_name" type:"string"`
	TypeOfFood string				 `json:"type_of_food" db:"type_of_food" type:"string"`
	Price      string 				 `json:"price" db:"price" type:"string"`
	CreatedAt *helperModel.Timestamp `json:"created_at" db:"created_at" type:"timestamp"`
	UpdatedAt *helperModel.Timestamp `json:"updated_at" db:"updated_at" type:"timestamp"`
	DeletedAt *helperModel.Timestamp `json:"deleted_at" db:"deleted_at" type:"timestamp"`

	MyFoods MyFoods `json:"my_foods" db:"-" fk:"relation:many,fk_field1:Id,fk_field2:food_id"`
}

type Foods []*Food

func (f *Food) GenarateUUID() {
	uuid, _ := uuid.NewGen().NewV4()
	f.Id = &uuid
}

func NewFoodWithParam(params map[string]interface{}) *Food {
	food := new(Food)

	if v, ok := params["name"]; ok {
		food.FoodName = v.(string)
	}

	if v, ok := params["my_foods"]; ok {
		for _, item := range v.([]interface{}) {
			myFood := NewMyFoodWithParam(item.(map[string]interface{}))
			food.MyFoods = append(food.MyFoods, myFood)
		}
	}

	return food
}
