package repository

import (
	"project-gorm/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func CreateSale(product *models.Product, price uint, client *models.Client, amount uint) *models.Sale {

	now := time.Now()
	date := now.Format("2006-01-02")

	return &models.Sale{
		ID: uuid.New().String(),
		Product: product,
		Client: client,
		Date: date,
		Amount: amount,

	}
}


func AddSale(db *gorm.DB, sale *models.Client, nameTable string) error {

	result := db.Table(nameTable).Create(sale)
	if result.Error != nil {
		return result.Error
	}
	return nil
}