package db

import (
	"littlegobook/shopping/models"
)

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
