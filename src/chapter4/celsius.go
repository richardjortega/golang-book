package main

import (
	"fmt"
)

func main() {
	// pg 46 - problem 5
	// Fahrenheit to Celsius converter
	fmt.Print("Enter temperature in Fahrenheit (converts to Celsius): ")
	var input float64

	// Fills input with the number entered by user
	fmt.Scanf("%f", &input)

	output := ((input - 32) * 5 / 9)

	fmt.Println(output)

}
