package main

import "fmt"

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	mi = 100
	fmt.Println(mi)
	// i := 10
	// fmt.Println(mi + i)
	mi.Print()

}
