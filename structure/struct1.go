package main

import "fmt"

type User struct {
	Name string
	Age  int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "updateuser"
}

func UpdateUser_pointer(user *User) {
	user.Name = "updateuser_pointer"
}

func main() {
	var user1 User = User{Name: "user1", Age: 1}
	fmt.Println(user1)

	user2 := User{Name: "user2", Age: 2}
	fmt.Println(user2)

	user3 := User{"user3", 3}
	fmt.Println(user3)

	var user4 User = User{"user4", 4}
	fmt.Println(user4)

	user5 := User{}
	fmt.Println(user5)

	// address
	user6 := new(User)
	fmt.Println(user6)
	user7 := &User{}
	fmt.Println(user7)

	UpdateUser(user1)
	fmt.Println(user1)

	UpdateUser_pointer(user6)
	fmt.Println(*user6)

}
