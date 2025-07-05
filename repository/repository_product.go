package repository

import (
	"project-gorm/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


func AddProduct(name string, price uint) *models.Product{
	return &models.Product{
		ID: uuid.New().String(),
		Name: name,
		Price: price,
	}
}

func InsertProduct(db *gorm.DB, product *models.Product, nameTable string) error {
	// Inserindo o Produto Na tabela
	result := db.Table(nameTable).Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}