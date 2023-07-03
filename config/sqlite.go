package config

import (
	"os"

	"github.com/LagLabs/backend-go.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqLIte() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"
	//Verify if DB exists

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database not found creating...")
		//Create the databa file and directory
		err = os.Mkdir("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		logger.Info("Database file not found creating...")
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}
	// Create and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQlite opening error")
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Oppening{})
	if err != nil {
		logger.Errorf("SQlite AutoMigrate error")
		return nil, err
	}

	return db, nil
}
