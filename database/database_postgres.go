package database

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitPostgres(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	CreateProductTable(db)
	CreateStockTable(db)
	CreateAddressTable(db)

	return db, nil
}

func CreateProductTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
}

func CreateStockTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Stock{})
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
}

func CreateAddressTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Address{})
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
}
