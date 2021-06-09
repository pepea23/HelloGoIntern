package models

import (
	helperModel "git.innovasive.co.th/backend/models"
)

const FK_FIELD_MY_FOODS = "MyFoods"

type Food struct {
	TableName struct{}               `json:"-" db:"food"`
	Id        int64                  `json:"id" db:"id" type:"int64"`
	Name      string                 `json:"name" db:"name" type:"string"`
	CreatedAt *helperModel.Timestamp `json:"created_at" db:"created_at" type:"timestamp"`
	UpdatedAt *helperModel.Timestamp `json:"updated_at" db:"updated_at" type:"timestamp"`
	DeletedAt *helperModel.Timestamp `json:"deleted_at" db:"deleted_at" type:"timestamp"`

	MyFoods MyFoods `json:"my_foods" db:"-" fk:"relation:many,fk_field1:Id,fk_field2:food_id"`
}

func NewFoodWithParam(params map[string]interface{}) *Food {
	food := new(Food)

	if v, ok := params["name"]; ok {
		food.Name = v.(string)
	}

	return food
}
