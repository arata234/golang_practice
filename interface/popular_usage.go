package main

import "fmt"

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

// 同じ名前の関数を定義する
func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

// 同じ名前の関数を定義する
func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Kohei", Age: 33},
		&Car{Number: "12-34", Model: "MINI Cooper"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
}
