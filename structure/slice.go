package main

import "fmt"

// 構造体のリスト

type Character struct {
	Name string
	Age  int
}

type Characters []*Character

func main() {
	c1 := Character{Name: "Gen Hoshino", Age: 42}
	c2 := Character{Name: "Osamu Mukai", Age: 41}
	c3 := Character{Name: "Ryo Yoshizawa", Age: 29}

	characters := Characters{}
	characters = append(characters, &c1, &c2, &c3)

	for _, v := range characters {
		fmt.Println(v.Age)
	}

	characters2 := make([]*Character, 0)
	characters2 = append(characters2, &c1, &c2, &c3)

	for _, v := range characters2 {
		fmt.Println(*v)
	}
}
