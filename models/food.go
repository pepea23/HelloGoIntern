package models

import (
	helperModel "git.innovasive.co.th/backend/models"
	"github.com/gofrs/uuid"
)

type Food struct {
	TableName  struct{}               `json:"-" db:"food"`
	Id         *uuid.UUID             `json:"id" db:"id" type:"uuid"`
	FoodName   string                 `json:"food_name" db:"food_name" type:"string"`
	TypeOfFood string                 `json:"type_of_food" db:"type_of_food" type:"string"`
	Price      string                 `json:"price" db:"price" type:"string"`
	CreatedAt  *helperModel.Timestamp `json:"created_at" db:"created_at" type:"timestamp"`
	UpdatedAt  *helperModel.Timestamp `json:"updated_at" db:"updated_at" type:"timestamp"`
	DeletedAt  *helperModel.Timestamp `json:"deleted_at" db:"deleted_at" type:"timestamp"`
}

type Foods []*Food

func (f *Food) GenarateUUID() {
	uuid, _ := uuid.NewGen().NewV4()
	f.Id = &uuid
}

func NewFoodWithParam(params map[string]interface{}) *Food {
	food := new(Food)

	if v, ok := params["food_name"]; ok {
		food.FoodName = v.(string)
	}
	if v, ok := params["type_of_food"]; ok {
		food.TypeOfFood = v.(string)
	}
	if v, ok := params["price"]; ok {
		food.Price = v.(string)
	}
	return food
}
