package main

import "fmt"

// var c, python, java bool // Declare at package level
var i, j int = 1, 2

const Pi = 3.14

func main() {
	var c, python, java = true, false, "no!" // Declare at function level
	fmt.Println(i, j, c, python, java)

	// Zero values
	var a int
	var empty string
	var falsy bool
	fmt.Println(a, empty, falsy)

	// Type conversion
	num := 42
	floatingNum := float64(num)
	unsignNum := uint(floatingNum)
	fmt.Println(num, floatingNum, unsignNum)

	fmt.Println("Approximation of Pi is:", Pi)
}
