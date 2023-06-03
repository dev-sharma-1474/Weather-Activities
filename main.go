package main

import (
	"fmt"
	"sync"
)

func main() {

	var wait sync.WaitGroup
	var cities []string
	cityData := make(map[string]*Forecast)
	//Allows user to enter cities they want data for
	for {
		var city string
		fmt.Print("Enter a city name. If you are done adding cities, type 'Done': ")
		fmt.Scanln(&city)
		if city == "Done" {
			break
		}
		cities = append(cities, city) //each city entered is stored in a table
	}
	//goroutine to get the data from each city from cities
	for _, city := range cities {
		wait.Add(1)
		go func(city string) {
			defer wait.Done()
			forecast := GetWeatherData(city)

			cityData[city] = forecast //Stores data in a map for future access
			forecast.Display(city)
		}(city)
	}

	wait.Wait()

	activity := GetIdealLocation()
	CompareCities(cityData, activity)

}
