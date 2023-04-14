package main

import "fmt"

func Double(i int) {
	i = i * 2
}

func main() {
	var n int = 300
	fmt.Println("value is ", n)
	fmt.Print("address is ", &n)

	Double(n)

	fmt.Println(n)

	// var p *int = &n
}
