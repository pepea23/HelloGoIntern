package models

import (
	helperModel "git.innovasive.co.th/backend/models"
	"github.com/gofrs/uuid"
)

type MyFood struct {
	TableName struct{}               `json:"-" db:"my_food"`
	Id        *uuid.UUID             `json:"id" db:"id" type:"uuid"`
	FoodId    *uuid.UUID             `json:"food_id" db:"food_id" type:"uuid"`
	My        string                 `json:"my" db:"my" type:"string"`
	CreatedAt *helperModel.Timestamp `json:"created_at" db:"created_at" type:"timestamp"`
	UpdatedAt *helperModel.Timestamp `json:"updated_at" db:"updated_at" type:"timestamp"`
	DeletedAt *helperModel.Timestamp `json:"deleted_at" db:"deleted_at" type:"timestamp"`
}

type MyFoods []*MyFood

func (f *MyFood) GenarateUUID() {
	uuid, _ := uuid.NewGen().NewV4()
	f.Id = &uuid
}

func NewMyFoodWithParam(params map[string]interface{}) *MyFood {
	myFood := new(MyFood)

	if v, ok := params["my"]; ok {
		myFood.My = v.(string)
	}

	return myFood
}
