package main

import "fmt"

func main() {
Loop:
	for {
		for {
			for {
				fmt.Println("Start")
				break Loop
			}
			fmt.Println("表示しない")
		}
		fmt.Println("表示しない")
	}
	fmt.Println("End")

	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue
			}
			fmt.Println(i, j, i+j)
		}
		fmt.Println("表示しない")
	}
}
