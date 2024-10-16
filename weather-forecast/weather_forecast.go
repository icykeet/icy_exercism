// Package weather provides tools for forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition present current state of the Weather.
var CurrentCondition string

// CurrentLocation present current location of the Weather.
var CurrentLocation string

// Forecast returns a readable output with weather information for a given city and condition, given the city and condition values as input strings.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
