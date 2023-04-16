package main

import (
	"fmt"
	"scope/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.20"
	return AppName + " " + Version
}

func main() {
	fmt.Println(foo.Max)
	fmt.Println(foo.ReturnMin())
	fmt.Println(appName())
	// fmt.Println(AppName)
}
