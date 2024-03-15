package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	
)

var DB *gorm.DB

func InitDB() {
	url := "postgres://tjryujji:5vllCYGk1i33xU7e3ZdF_ecGCnOGYm-N@topsy.db.elephantsql.com:5432/tjryujji"
	// url := os.Getenv("URL") 
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

}
