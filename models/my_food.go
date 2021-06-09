package models

import (
	helperModel "git.innovasive.co.th/backend/models"
)

type MyFood struct {
	TableName struct{}               `json:"-" db:"my_food"`
	Id        int64                  `json:"id" db:"id" type:"int64"`
	FoodId    int64                  `json:"food_id" db:"food_id" type:"int64"`
	CreatedAt *helperModel.Timestamp `json:"created_at" db:"created_at" type:"timestamp"`
	UpdatedAt *helperModel.Timestamp `json:"updated_at" db:"updated_at" type:"timestamp"`
	DeletedAt *helperModel.Timestamp `json:"deleted_at" db:"deleted_at" type:"timestamp"`
}

type MyFoods []*MyFood
