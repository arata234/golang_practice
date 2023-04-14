package main

import "fmt"

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know.")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	// 条件式内での変数の定義はその中でのみ有効
	s := "out"
	if s := "in"; true {
		fmt.Println(s)
	}
	fmt.Println(s)

}
