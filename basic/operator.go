package main

import "fmt"

func main() {

	//算術演算子
	// fmt.Println(1 + 2)
	// fmt.Println("ABC" + "DE")

	// fmt.Println(2 - 1)
	// fmt.Println(3 * 2)
	// fmt.Println(80 / 5)
	// fmt.Println(9 % 4)

	// n := 0
	// n += 4
	// fmt.Println(n)
	// n++
	// fmt.Println(n)
	// n--
	// fmt.Println(n)

	// fmt.Println(1 == 1)
	// fmt.Println(1 == 2)
	// fmt.Println(1 < 5)
	// fmt.Println(1 > 5)
	// fmt.Println(true == false)
	// fmt.Println(true != false)

	fmt.Println(true && false == true) // 1 and 0 == 1 -> 0
	fmt.Println(true && true == true)  // 1 and 1 == 1 -> 1
	fmt.Println(true && false == false)// 1 and 0 == 0 -> 1

	fmt.Println(true || false == true) // 1 or 0 == 1 -> 1
	fmt.Println(true || true == true) // 1 or 1 == 1 -> 1
	fmt.Println(false || false == true) // 0 or 0 == 1-> 0

	fmt.Println(!true)
	fmt.Println(!false)
}
