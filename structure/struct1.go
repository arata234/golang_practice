package main

import "fmt"

type User struct {
	Name string // init [ ]
	Age  int    // init [0]
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	//fmt.Println(user1)
	user1.Name = "Tom"
	user1.Age = 18
	fmt.Println(user1)

	user2 := User{}
	user2.Name = "Liz"
	user2.Age = 14
	fmt.Println(user2)

	user3 := User{Name: "Bob", Age: 25}
	fmt.Println(user3)

	user4 := User{"Monica", 29}
	fmt.Println(user4)

	user5 := new(User)
	fmt.Println(user5)

	user6 := &User{}
	fmt.Println(user6)

	UpdateUser(user1)
	UpdateUser2(user6)

	fmt.Println(user1)
	fmt.Println(user6)

}
