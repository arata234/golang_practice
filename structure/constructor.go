package main

import "fmt"

type Character struct {
	Name string
	Age  int
}

func NewCharacter(name string, age int) *Character {
	return &Character{Name: name, Age: age}
}

func main() {
	c1 := Character{"Gen Hoshino", 42}
	c2 := NewCharacter("Osamu Mukai", 41)
	fmt.Println(c1)
	fmt.Println(*c2)
}
