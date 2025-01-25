package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"os"
)

type OpenWeatherResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`
		Humidity  int     `json:"humidity"`
		FeelsLike float64 `json:"feels_like"`
		Pressure  int     `json:"pressure"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
		Main        string `json:"main"`
	} `json:"weather"`
}

func GetWeatherByCoordinates(c *fiber.Ctx) error {
	lat := c.Query("lat")
	lon := c.Query("lon")
	if lat == "" || lon == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "lat and lon parameters are required",
		})
	}
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	baseURL := os.Getenv("OPENWEATHER_BASE_URL")

	url := fmt.Sprintf("%s/weather?lat=%s&lon=%s&appid=%s&units=metric", baseURL, lat, lon, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error making API request: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch weather data",
		})
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Error response from API: %s", string(body))
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": "Failed to fetch weather data from OpenWeather API",
		})
	}

	var weatherData OpenWeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		log.Printf("Error decoding API response: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode weather data",
		})
	}

	filteredResponse := fiber.Map{
		"main":    weatherData.Main,
		"weather": weatherData.Weather,
	}
	return c.Status(fiber.StatusOK).JSON(filteredResponse)
}
