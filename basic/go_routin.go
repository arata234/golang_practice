package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func init() {
	fmt.Println("init")
}

func main() {
	/*
		go sub()
		for {
			fmt.Println("Main Loop")
			time.Sleep(200 * time.Millisecond)
		}
	*/
	fmt.Println("main")
}
