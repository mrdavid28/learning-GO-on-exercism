// Package weather is overall package for Weather Forecast.
package weather

// CurrentCondition describes current weather condition.
var CurrentCondition string

// CurrentLocation variable stores current Location.
var CurrentLocation string

// Forecast function shows the weather for the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
