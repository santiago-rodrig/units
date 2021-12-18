package main

const MetersFeet = 3.28084
const KilogramsPounds = 2.20462

const (
	length      UnitKind = "length"
	weight               = "weight"
	temperature          = "temperature"
)

const (
	meters            Unit = "meters"
	feet                   = "feet"
	degreesCelsius         = "celsius"
	degreesFahrenheit      = "fahrenheit"
	kilograms              = "kilograms"
	pounds                 = "pounds"
)

var UnitsMap = map[UnitKind][]Unit{
	length:      {meters, feet},
	weight:      {kilograms, pounds},
	temperature: {degreesCelsius, degreesFahrenheit},
}
