package main

import (
	"fmt"
)

func main() {
	// pg 46 - problem 6
	// converts feet to meters converter
	fmt.Print("Enter number of feet (converts to meters): ")
	var input float64

	// Fills input with the number entered by user
	fmt.Scanf("%f", &input)

	output := (input * 0.3048)

	fmt.Println(output)

}
