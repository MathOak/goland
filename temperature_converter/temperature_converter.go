package main

import "fmt"

const boilingKelvin = 373.0

func main() {
	tempK := boilingKelvin
	tempC := tempK - 273.0

	fmt.Printf("Temperature in Kelvin: %g, Temperature in Celsius: %g\n", tempK, tempC)
}
