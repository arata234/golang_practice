package main

import "fmt"

//channel (queue)
// 複数のgo routine間でのdataの受け渡しをするために設計されたdata structure

func main() {
	var ch1 chan int

	//recipient
	// var ch2 <-chan int

	//sender
	// var ch3 chan<- int

	ch1 = make(chan int)
	ch2 := make(chan int, 5)
	fmt.Println(cap(ch1))
	// fmt.Println(cap(ch2))

	ch2 <- 1
	// fmt.Println("len(ch2)", len(ch2))

	ch2 <- 2
	ch2 <- 3
	fmt.Println("len(ch2)", len(ch2))

	i := <-ch2
	fmt.Println(i)
	fmt.Println("len(ch2)", len(ch2))

	i2 := <-ch2
	fmt.Println(i2)
	fmt.Println("len(ch2)", len(ch2))
	fmt.Println(<-ch2)
	fmt.Println("len(ch2)", len(ch2))

	ch2 <- 2
	ch2 <- 3
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6

}
