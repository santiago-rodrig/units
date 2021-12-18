// units takes a number with its unit, and converts to a list
// of available destination units
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) != 3 {
		fmt.Println("Expected exactly 3 arguments")
		os.Exit(1)
	}
	n, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	srcUnit := Unit(os.Args[2])
	dstUnit := Unit(os.Args[3])
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(4)
		}
	}()
	srcUnitKind := SearchUnitKind(srcUnit)
	dstUnitKind := SearchUnitKind(dstUnit)
	if srcUnitKind != dstUnitKind {
		fmt.Printf("cannot convert %q into %q\n", srcUnitKind, dstUnitKind)
		os.Exit(3)
	}
	fmt.Printf("%f %s = %f %s\n", n, srcUnit, Convert(n, srcUnit, dstUnit), dstUnit)
}

func Convert(scalar float64, srcUnit, dstUnit Unit) float64 {
	switch srcUnit {
	case meters:
		switch dstUnit {
		case feet:
			return MetersToFeet(scalar)
		}
	case feet:
		switch dstUnit {
		case meters:
			return FeetToMeters(scalar)
		}
	case degreesCelsius:
		switch dstUnit {
		case degreesFahrenheit:
			return CelsiusToFahrenheit(scalar)
		}
	case degreesFahrenheit:
		switch dstUnit {
		case degreesCelsius:
			return FahrenheitToCelsius(scalar)
		}
	case kilograms:
		switch dstUnit {
		case pounds:
			return KilogramsToPounds(scalar)
		}
	case pounds:
		switch dstUnit {
		case kilograms:
			return PoundsToKilograms(scalar)
		}
	}
	panic("unhandled conversion")
}

func SearchUnitKind(unit Unit) UnitKind {
	for kind, units := range UnitsMap {
		for _, unitValue := range units {
			if unit == unitValue {
				return kind
			}
		}
	}
	panic("nonexistent kind")
}
