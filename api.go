package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherAPIResponse struct {
	Main        WeatherAPIWeather `json:"main"`
	Wind        WeatherAPIWind    `json:"wind"`
	WeatherInfo []WeatherAPIInfo  `json:"weather"`
}

// WeatherAPIWeather gets temperature and humidity from json output
type WeatherAPIWeather struct {
	Temperature float64 `json:"temp"`
	Humidity    float64 `json:"humidity"`
}

// WeatherAPIWind gets wind speed from json output
type WeatherAPIWind struct {
	Speed float64 `json:"speed"`
}

// WeatherAPIInfo gets weather description(sunny, cloudy etc.) from json output
type WeatherAPIInfo struct {
	Description string `json:"description"`
}

// fetchWeatherData communicates with api to get the weather data
func fetchWeatherData(cityName string) (*WeatherAPIResponse, error) {

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=" + cityName + "&units=imperial&appid=6541b7d812295e29233904dbd1c13a7a")

	// Make the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// Reads the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Parses the JSON response
	var weatherResponse WeatherAPIResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return nil, err
	}

	return &weatherResponse, nil
}
