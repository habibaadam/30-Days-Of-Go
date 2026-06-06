package main

import (
	"fmt"
	"project_layout_lessons/internal/greet"
)

func main() {
	message1 := greet.Hello("Habie")
	fmt.Println(message1)
}
