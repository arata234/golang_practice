package main

import "fmt"

// function

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(x int) int {
	d := 2 * x
	return d
}

func Noreturn() {
	fmt.Println("No return")
	return
}

func main() {
	i := Plus(10, 20)
	fmt.Println(i)
	q, r := Div(17, 5)
	fmt.Println(q, r)

	d := Double(4)
	fmt.Println(d)

	Noreturn()
}
