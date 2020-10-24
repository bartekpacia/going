package main

import "fmt"

func main() {

	fmt.Printf("Enter x and y: ")

	var x, y int
	fmt.Scanf("%d %d", &x, &y)

	fmt.Println(x, y)

	q, r := IntDiv(x, y)

	fmt.Println(q, r)
}

// IntDiv divides an integer. Returns the quotitient and the remainder.
// dane początkowe: x, y
// warunek początkowy: x >= 0 i y >= 0
func IntDiv(x int, y int) (q int, r int) {
	q = x / y
	r = x - (q * y)

	return
}
