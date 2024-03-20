package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	
)

var DB *gorm.DB

func InitDB() {
	url := "postgres://jglxdxwd:Yg-Dj1U_4lwhdRwsCR81eK8nnDWX9l22@floppy.db.elephantsql.com/jglxdxwd"
	// url := os.Getenv("URL") 
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

}
