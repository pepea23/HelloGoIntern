package models

type Food struct {
	TableName struct{} `json:"-" db:"food"`
	Name      string   `json:"name" db:"name" type:"string"`
	Quntity   int64    `json:"quntity" db:"quntity" type:"int64"`
}

func NewFoodWithParam(params map[string]interface{}) *Food {
	food := new(Food)

	if v, ok := params["name"]; ok {
		food.Name = v.(string)
	}

	if v, ok := params["quntity"]; ok {
		food.Quntity = int64(v.(float64))
	}

	return food
}
