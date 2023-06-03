package main

import (
	"fmt"
	"math"
)

type Activity struct {
	ActivityName string
	Temp         float64
}

// GetIdealLocation allows user to input data for activity & ideal forecast
func GetIdealLocation() *Activity {
	var act string
	var temperature float64
	fmt.Println("_________________________")
	fmt.Print("Name your activity: ")
	fmt.Scanln(&act)
	fmt.Print("Enter ideal temperature in Fahrenheit: ")
	fmt.Scanln(&temperature)
	activity := &Activity{
		ActivityName: act,
		Temp:         temperature,
	}
	return activity
}

// CompareCities compares the temperature difference between cities and activities, showing the best matches
func CompareCities(cityData map[string]*Forecast, activity *Activity) {
	var lowerTempDiff float64 = -9999
	var higherTempDiff float64 = 9999
	var bestHotCity string
	var bestColdCity string
	for city, forecast := range cityData {
		diff := forecast.Temp - activity.Temp
		if diff > 0 && diff < higherTempDiff {
			bestHotCity = city
			higherTempDiff = diff
		} else if diff <= 0 && diff > lowerTempDiff {
			bestColdCity = city
			lowerTempDiff = diff
		}
	}
	//truncates lowerTempDiff and higherTempDiff to 2 decimal places
	lowerTempDiff = math.Trunc(lowerTempDiff*10) / 10
	higherTempDiff = math.Trunc(higherTempDiff*10) / 10
	if higherTempDiff != 9999 {
		fmt.Print("With the given data, " + bestHotCity + " is ")
		fmt.Print(higherTempDiff)
		fmt.Print(" degrees hotter than the desired temperature. ")

	}
	if lowerTempDiff != -9999 {
		fmt.Print(bestColdCity + " is ")
		fmt.Print(math.Abs(lowerTempDiff))
		fmt.Print(" degrees colder than the desired temperature")
	}

}
