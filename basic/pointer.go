package main

import "fmt"

func Double(i int) {
	i = i * 2
}

// 参照渡し 引数をポインタ型とすることでアドレス内での計算を行う
func Double_pointer(i *int) {
	*i = *i * 2
}

func Double_slice(s []int) {
	for i := range s {
		s[i] = s[i] * 2
	}
}

func main() {
	var n int = 100
	fmt.Println("value is ", n)
	fmt.Println("address is ", &n)

	Double(n)
	fmt.Println("Double number n is", n)

	// pにnのアドレスを渡す
	// *で型宣言することでポインタ型にすることができる
	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	// pとnのアドレスが同じだから
	*p = 300
	fmt.Println(n)

	n = 200
	fmt.Println(*p)

	Double_pointer(&n)
	fmt.Println(n)

	// sliceは参照型だからそのままつかってok
	var sl []int = []int{1, 2, 3}
	Double_slice(sl)
	fmt.Println(sl)

}
