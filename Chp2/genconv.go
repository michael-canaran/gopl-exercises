package main

import "fmt"

type Pounds float64
type Feet float64
type Kilograms float64
type Centimeters float64

// Pounds to kilos
func PoundsToKilo(pounds Pounds) Kilograms {
	return Kilograms(pounds / 2.205)
}

//Kilos to pounds
func KilosToPounds(kilos Kilograms) Pounds {
	return Pounds(kilos * 2.205)
}

func FeetToCent(feet Feet) Centimeters {
	return Centimeters(feet * 30.48)
}

func CentToFeet(cent Centimeters) Feet {
	return Feet(cent / 30.48)
}

func MassConv(v float64) string {
	return fmt.Sprintf("Value: %g\nIn pounds: %g\nIn feet: %g\nIn kilos: %g\nIn centimeters: %g\n", v, KilosToPounds(Kilograms(v)), CentToFeet(Centimeters(v)), PoundsToKilo(Pounds(v)), FeetToCent(Feet(v)))
}

func (p Pounds) String() string {
	return fmt.Sprintf("%g lb", p)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g ft", f)
}

func (k Kilograms) String() string {
	return fmt.Sprintf("%g kg", k)
}

func (c Centimeters) String() string {
	return fmt.Sprintf("%g cm", c)
}
