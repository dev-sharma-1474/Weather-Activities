package main

import (
	"fmt"
)

type Forecast struct {
	Temp        float64
	Humidity    float64
	Wind        float64
	WeatherType string
}

// Display Prints data in a readable format for the user
func (f *Forecast) Display(city string) {
	fmt.Println("_________________________")
	fmt.Println("Forecast for " + city + ":")
	fmt.Println("Temperature:", f.Temp, "°F")
	fmt.Println("Humidity:", f.Humidity, "%")
	fmt.Println("Wind:", f.Wind, "mph")
	fmt.Println("Weather type:", f.WeatherType)

}

// GetWeatherData communicates with api.go to get relevant data from json output
func GetWeatherData(city string) *Forecast {

	weatherData, _ := fetchWeatherData(city)

	forecast := &Forecast{
		Temp:        weatherData.Main.Temperature,
		Humidity:    weatherData.Main.Humidity,
		Wind:        weatherData.Wind.Speed,
		WeatherType: weatherData.WeatherInfo[0].Description,
	}

	return forecast
}
