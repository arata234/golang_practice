package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	/*
		sl := []int{100, 200}
		fmt.Println(sl)

		sl = append(sl, 300)
		fmt.Println(sl)

		sl = append(sl, 400, 500, 600)
		fmt.Println(sl)

		sl2 := make([]int, 5)
		fmt.Println(sl2)
		fmt.Println(len(sl2))
		// fmt.Println(cap(sl2))
		sl2 = append(sl2, 0)
		// 開発でメモリなどを気にするときは必要
		fmt.Println(cap(sl2))
	*/
	/*
		sl3 := sl2
		sl2[0] = 1000
		fmt.Println(sl2, sl3)

		sl := []int{1, 2, 3, 4, 5}
		sl2 := make([]int, 5, 10)
		n := copy(sl2, sl)
		fmt.Println(n, sl2)
	*/
	/*
		sl := []string{"A", "B", "C"}
		fmt.Println(sl)

		for _, v := range sl {
			fmt.Println(v)
		}
	*/

	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))

}
