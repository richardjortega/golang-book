// Takeaway Notes
// [1] Use shortest form possible in code when using vars
// [2] Go is lexically scoped using blocks ( e.g. { } )

package main

import (
	"fmt"
)

// This style of shorthand outer scope doesnt work
// var global_var := "Hello World [global scope]"
var global_var string = "Hello World [global scope]"

func main() {
	/// Example of optional/alt way of variable
	// var x string
	// x = "Hello World"
	// fmt.Println(x)

	/// Variable declaration with type and value on single line
	/// Note: Cannot redeclare a variable, but you can reassign it's value
	// var x string = "Hello World"
	// fmt.Println(x)

	// pg 37
	// Reassigning variables
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	// x = x + "third"
	x += "third" // special assignment operator
	fmt.Println(x)

	// pg 38
	// change example var names avoid redeclaration of var
	// Note: If you declare a variable and it's not used it will not compile
	var y string = "hello"
	var z string = "world"
	fmt.Println(y == z) // false

	z = "hello"
	fmt.Println(y == z) // true after reassignment

	// pg 39
	// var declaration and assignment (shorthand method)
	// type infered by Go compiler
	// Note: Cannot use shorthand for reassignment only as it's also doing a declartion
	a := "Hello World"
	fmt.Println(a)

	// type inference can also happen on var declaration as well
	var b = "Hello World"
	fmt.Println(b)

	// type inference works across various inputs
	c := 5
	fmt.Println(c)

	// pg 40
	// naming vars
	d := "Max" // terrible var name
	fmt.Println("My dog's name is", d)

	name := "Max"
	fmt.Println("My dog's name is", name)

	dogsName := "Max" // camelcase var declaration is common
	fmt.Println("My dog's name is", dogsName)

	// pg 41
	// scoping vars (using outside global var)
	fmt.Println(global_var)

	// pg 43
	// constants
	const e string = "I'm a constant string"
	fmt.Println(e)
	// Cannot reassign values to const
	// e = "Trying to reassign a const value"

	// pg 44
	// defining multiple variables
	// Note: const can be used in a similar way
	var (
		f = 5
		g = "don't sweat the technique"
		h = 10
	)
}

func f() {
	fmt.Println(global_var)
}
