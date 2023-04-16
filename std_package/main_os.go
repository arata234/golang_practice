package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		fmt.Println("Start")

		defer func() {
			fmt.Println("defer")
		}()

		_, err := os.Open("A.txt")
		if err != nil {
			log.Fatalln(err)
		}
	*/
	/*
		// Args
		fmt.Println(os.Args[0])
		fmt.Println(os.Args[1])
		fmt.Println(os.Args[2])
		fmt.Println(os.Args[3])

		fmt.Printf("length = %d\n", len(os.Args))
		for i, v := range os.Args {
			fmt.Println(i, v)
		}
	*/
	/*
		//Create
		f, _ := os.Create("foo.txt")
		f.Write([]byte("Hello\n"))
		f.WriteAt([]byte("Golang\n"), 6)
		f.Seek(0, os.SEEK_END)
		f.WriteString("Yeah")

		bs := make([]byte, 128)

		n, err := f.Read(bs)
		if err != nil {

		}
		fmt.Println(n)
		fmt.Println(string(bs))

		bs2 := make([]byte, 128)

		nn, err := f.ReadAt(bs2, 10)
		fmt.Println(nn)
		fmt.Println(string(bs2))
	*/

	f, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

}
