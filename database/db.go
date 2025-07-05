package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"project-gorm/models"
)

func CreateDataBaseAndTables() error {
	database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	database.AutoMigrate(&models.Product{}, &models.Client{}, &models.Sale{})

	return nil
}
