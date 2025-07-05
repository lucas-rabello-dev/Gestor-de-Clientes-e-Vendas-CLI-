package repository

import (
	"project-gorm/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Create Client
func CreateClient(name string,email string, phoneNumber int ) *models.Client {
	return &models.Client{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		PhoneNumber: phoneNumber,
	}
}

// func ValidEmail() error {

// }

// Adicionar Cliente
func AddClient(db *gorm.DB, client *models.Client, nameTable string) error {
	// Inserindo o Cliente Na tabela
	result := db.Table(nameTable).Create(client)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Ver Slientes 
// func ReadClients() {

// }

