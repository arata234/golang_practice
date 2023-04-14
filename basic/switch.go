package main

import "fmt"

func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	/*
		n := 5
		switch n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("I don't know.")
		}
	*/
	/*
		switch n := 2; n {
		case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("I don't know.")
		}
	*/

	anything("aaa")
	anything(1)
	/*
		var x interface{} = 3
		i, isInt := x.(int)
		fmt.Println(i+2, isInt)

		f, isFloat64 := x.(float64)
		fmt.Println(f, isFloat64)

		if x == nil {
			fmt.Println("None")
		} else if i, isInt := x.(int); isInt {
			fmt.Println(i, "x is Int")
		} else if s, isString := x.(string); isString {
			fmt.Println(s, "x is String")
		} else {
			fmt.Println("I don't know")
		}
	*/
	var x interface{} = 3
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know.")
	}
}
