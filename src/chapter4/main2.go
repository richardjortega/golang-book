package main

import (
	"fmt"
)

func main() {
	// pg 45
	fmt.Print("Enter a number (we will multiply by 2): ")
	var input float64

	// Fills input with the number entered by user
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
