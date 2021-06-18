package models

import (
	helperModel "git.innovasive.co.th/backend/models"
	"github.com/gofrs/uuid"
)

const FK_FIELD_FOODS = "Foods"

type Restaurant struct {
	TableName  struct{}                    `json:"-" db:"restaurant"`
	Id              *uuid.UUID             `json:"id" db:"id" type:"uuid"`
	RestaurantName  string                 `json:"restaurant_name" db:"restaurant_name" type:"string"`
	Address         string                 `json:"address" db:"address" type:"string"`
	PhoneNumber     string                 `json:"phone_number" db:"phone_number" type:"string"`
	OpenCloseTime   string				   `json:"open_close_time" db:"open_close_time" type:"string"`
	CreatedAt  *helperModel.Timestamp `json:"created_at" db:"created_at" type:"timestamp"`
	UpdatedAt  *helperModel.Timestamp `json:"updated_at" db:"updated_at" type:"timestamp"`
	DeletedAt  *helperModel.Timestamp `json:"deleted_at" db:"deleted_at" type:"timestamp"`

	Foods Foods `json:"my_foods" db:"-" fk:"relation:many,fk_field1:Id,fk_field2:food_id"`
}

