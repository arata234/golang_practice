package main

import "fmt"

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)
	fmt.Println(m["A"])

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)

	m4[1] = "Japan"
	m4[2] = "USA"
	fmt.Println(m4)
	fmt.Println(m4[2])

	m4[3] = "China"
	fmt.Println(m4)
	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))

	m5 := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range m5 {
		fmt.Println(k, v)
	}
}
