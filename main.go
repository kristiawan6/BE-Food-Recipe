package main

import (
	"be_food_recipe/src/config"
	"be_food_recipe/src/helper"
	"be_food_recipe/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	
	app := fiber.New()
	
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	config.InitDB()
	helper.Migration()
	routes.Router(app)

	app.Listen(":8080")
}
