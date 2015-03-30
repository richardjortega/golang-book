package main

import "fmt"

func main() {
	// pg 28
	// Playing with integers and floats and arguments
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.01-0.2)

	// pg 30
	// Messing with various print styles
	// Strings indexed at 0
	fmt.Println(len("Hello World")) // spaces are of course characters
	fmt.Println("Hello World"[1])   // each character represented by a byte (byte is an integer)
	fmt.Println("Hello " + "World")

	// pg 32
	// Booleans with logical operators
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false

	// pg 34
	// problem sets

	// 3 - multiplication
	fmt.Println(321325 * 424521)

	// 4 - find string's length
	fmt.Println(len("Hello World")) // func `len` used to establish length

	// 5 - evaluate expression for value [answer: true]
	// Note: Found out you can break apart expression
	//  into easier to read parts of expression without
	//  Go complaining! W00t!
	fmt.Println((true && false) ||
		(false && true) ||
		!(false && false))
}
