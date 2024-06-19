package database

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func InitPostgres(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	CreateProductTable(db)
	CreateTransactionTable(db)
	CreateStockTable(db)
	CreateOrderTable(db)
	CreateItemTable(db)

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

func CreateOrderTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
}

func CreateTransactionTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Transaction{})
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
}

func CreateItemTable(db *gorm.DB) {
	err := db.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
}
