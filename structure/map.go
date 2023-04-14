package main

import "fmt"

type Character struct {
	Name string
	Age  int
}

func main() {
	m := map[int]Character{
		1: {Name: "Gen Hoshino", Age: 42},
		2: {Name: "Osamu Mukai", Age: 41},
	}
	fmt.Println(m)

	m2 := map[Character]string{
		{Name: "Gen Hoshino", Age: 42}: "1/28",
		{Name: "Osamu Mukai", Age: 41}: "2/7",
	}
	fmt.Println(m2)

}
