package main

import "fmt"

// 関数を返り値に返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function in function")
	}
}

// 関数を引数に取る関数
func CallFunction(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {

	// 無名関数
	// f := func(x, y int) int {
	// 	return x + y
	// }

	// i := f(1, 2)
	// fmt.Println(i)

	// i2 := func(x, y int) int {
	// 	return x + y
	// }(1, 2)

	// fmt.Println(i2)

	// // 関数を返り値に返す関数
	// f2 := ReturnFunc()
	// f2()

	// // 関数を引数に取る関数
	// CallFunction(func() {
	// 	fmt.Println("I'm a CallFunction.")
	// })

	f3 := Later()
	fmt.Println(f3("Hello"))
	fmt.Println(f3("My"))
	fmt.Println(f3("name"))
	fmt.Println(f3("is"))
	fmt.Println(f3("Golang"))
}
