package main

func MetersToFeet(meters float64) float64 {
	return meters * MetersFeet
}

func FeetToMeters(feet float64) float64 {
	return feet / MetersFeet
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9
}

func KilogramsToPounds(kilograms float64) float64 {
	return kilograms * KilogramsPounds
}

func PoundsToKilograms(pounds float64) float64 {
	return pounds / KilogramsPounds
}
