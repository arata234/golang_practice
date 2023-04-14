package main

import "fmt"

type character struct {
	Name string
	Age  int
}

type actor struct {
	character character
	birthday  string
}

func (c character) SayName() {
	fmt.Println(c.Name)
}

func main() {

	c := character{
		Name: "Gen Hoshino",
		Age:  42,
	}

	a := actor{
		character: c,
		birthday:  "1/28",
	}

	a.character.SayName()
}
