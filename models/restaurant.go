package models

import (
	helperModel "git.innovasive.co.th/backend/models"
	"github.com/gofrs/uuid"
)

const FK_FIELD_FOODS = "Foods"

type Restaurant struct {
	TableName      struct{}               `json:"-" db:"restaurant"`
	Id             *uuid.UUID             `json:"id" db:"id" type:"uuid"`
	RestaurantName string                 `json:"restaurant_name" db:"restaurant_name" type:"string"`
	Address        string                 `json:"address" db:"address" type:"string"`
	PhoneNumber    string                 `json:"phone_number" db:"phone_number" type:"string"`
	OpenCloseTime  string                 `json:"open_close_time" db:"open_close_time" type:"string"`
	CreatedAt      *helperModel.Timestamp `json:"created_at" db:"created_at" type:"timestamp"`
	UpdatedAt      *helperModel.Timestamp `json:"updated_at" db:"updated_at" type:"timestamp"`
	DeletedAt      *helperModel.Timestamp `json:"deleted_at" db:"deleted_at" type:"timestamp"`

	Foods Foods `json:"my_food" db:"-" fk:"relation:many,fk_field1:Id,fk_field2:food_id"`
}

func (r *Restaurant) GenarateUUID() {
	uuid, _ := uuid.NewGen().NewV4()
	r.Id = &uuid
}

func NewRestaurantWithParam(params map[string]interface{}) *Restaurant {
	restaurant := new(Restaurant)

	if v, ok := params["restaurant_name"]; ok {
		restaurant.RestaurantName = v.(string)
	}
	if v, ok := params["address"]; ok {
		restaurant.Address = v.(string)
	}
	if v, ok := params["phone_number"]; ok {
		restaurant.PhoneNumber = v.(string)
	}
	if v, ok := params["open_close_time"]; ok {
		restaurant.OpenCloseTime = v.(string)
	}
	return restaurant
}
