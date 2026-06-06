package main

import (
	"fmt"
	"math"
)



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

// interfaces - a way to group struct methods associated with type
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// sqaure methods
func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circum() float64 {
	return 2 * math.Pi * c.radius
}

func printShape(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
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
