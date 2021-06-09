package orm

import (
	"strings"

	"github.com/HelloGoIntern/models"
	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
)

func OrmMyFood(ptx *models.MyFood, rows *sqlx.Rows) (*models.MyFood, error) {
	tableName := GetTableName(ptx)
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	_, ptrColumnMap := GetStructFields(ptx)

	if err := rows.Err(); err != nil {
		return nil, err
	}

	values, err := rows.SliceScan()
	if err != nil {
		return nil, err
	}

	if len(values) > 0 {
		for index, col := range columns {
			orderCol := strings.ReplaceAll(col, tableName+".", "")
			if field, ok := ptrColumnMap.Load(orderCol); ok {
				if err := SetFieldFromType(field.(*structs.Field), values[index]); err != nil {
					return nil, err
				}
			}
		}
	}

	return ptx, nil
}
