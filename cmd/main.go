package main

import (
	"github.com/gofiber/fiber/v2"
	"practice_fiber/internal/handler"
	"log"
)

func main() {

	app := fiber.New()
	
	app.Get("/weather", handler.GetWeatherHandler)
	
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}