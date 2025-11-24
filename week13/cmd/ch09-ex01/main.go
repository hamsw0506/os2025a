package main

import "fmt"

type Meters float64
type Kilomiters float64
type Miles float64

func (k Kilomiters) ToMiles() Miles {
	return Miles(k * 0.621371)
}
func (m Meters) ToMiles() Miles {
	return Miles(m * 0.000621371)
}

func main() {
	kmph := Kilomiters(150)
	fmt.Printf("%0.2f kilometer per hour equals %0.2f mile per hour\n", kmph, kmph.ToMiles())
	meter := Meters(500)
	fmt.Printf("%0.2f meter equals %0.2f mile\n", meter, meter.ToMiles())
}
