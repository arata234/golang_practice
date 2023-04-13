package main

import "fmt"


var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}




func main() {
	//var name type = value
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	//同時に変数宣言する
	var t, f bool = true, false
	fmt.Println(t, f)

	//同時に変数宣言する
	var (
		i2 int = 200
		s2 string = "GoLang"
	)
	fmt.Println(i2, s2)

	//変数宣言だけ先に行う
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	//すでに入っている変数を変更する
	i = 150
	fmt.Println(i)

	//暗黙的な定義
	//name := value

	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 暗黙的な定義では型の違う値に変更することはできない
	// i4 = "Go"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()
	
	//他の関数で定義した変数は使用できない
	//fmt.Println(s4)

	// 変数が使われていないとerrorが出る
	// var s5 string = "Not Use"
	

}