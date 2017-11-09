package main

import "fmt"

// func add(x int, y int) int {
func add(x, y int) int { // same as above
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(18))
}
