package main

import (
	"fmt"
	"math"
)

/*
	Square root estination by Newton's method
	Newton's Method is the way to estimation square root value by beginning from z = 1 then each iteration we will adjust z value by formular
	z -= (z*z - x) / (2*z)
	x is the one that we want to get sqaure root value and z is result
	By each iteration the z value will closer to the real one but how many of iteration is need? It's can be 10 billion but the value of each iteration will just different a little bit
	So we can define some "diff" value if diff is not more than this value we can stop the iteration
*/
func Sqrt(x float64) float64 {
	z := 1.0
	diff := 1.0
	for diff > 0.000000000000001 {
		var prevZ = z
		z -= (z*z - x) / (2 * z)
		diff = math.Abs(z - prevZ)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(64))
}
