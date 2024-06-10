package database

import (
	"awesomeProject/config"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	SQL *gorm.DB
}

func InitDatabase(config *config.Config) (*Database, error) {
	var db Database
	var err error

	switch config.DBDriver {
	case "postgres":
		db.SQL, err = InitPostgres(config)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
		}
	default:
		log.Fatal("Unsupported database driver")
	}

	return &db, nil
}
