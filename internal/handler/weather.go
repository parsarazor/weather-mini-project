	// what does handler?
	// registers a route bound to a specific HTTP method
package handler

import (
	"github.com/gofiber/fiber/v2"
	"practice_fiber/internal/service"
)

type WeatherResponse struct {
	City 	string `json:"city"`
	Temperature float64 `json:"tempreture"`
}

func GetWeatherHandler(c *fiber.Ctx) error {
	city := c.Query("city")
	if city == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "city name required",
		})
	}
	temp, err := service.GetWeather(city)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":  "weather provide error",
		})
	}
	return c.JSON(WeatherResponse{
		City:		city,
		Temperature: temp,
	})
	
}