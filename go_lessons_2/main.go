package main

import "fmt"



type User struct {
	ID int
	Name string
	Email string
	Age int
}

// a method associated with a specific type
// value receiver
func (u User) Introduce() string {
	return fmt.Sprintf("Hi, I am %s", u.Name)
}

// struct methods recieving pointers and updating age
// pointer receiver. used if a struct method needs to change struct's fields
// best for perfomance if you have large structs, as they use memory address
func (u *User) updatebirthday() {
	u.Age++
}


func main() {

	user1 := User{
		ID: 1,
		Name: "Habie",
		Email: "habie@gmail.com",
		Age: 14,
	}

	fmt.Println(user1)
	user1.Email = "habs@gmail.com"
	fmt.Println(user1)

	fmt.Println(user1.Introduce())

	user1.updatebirthday()

	fmt.Println(user1.Age)

}
