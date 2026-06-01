package main

import (
	"fmt"
	"strconv"
	"errors"
)



func spreadOperator(firstArr []string, secondArr []string) {
	firstArr = []string{"learn golang", "watch anime", "solve puzzles"}
	secondArr = []string{"eat cake"}
	firstArr = append(firstArr, secondArr...) //spread operator

	fmt.Println(firstArr)
}

func loopRange(slice []int) {
	total := 0

	for i, v := range slice {
		fmt.Printf("Day is %v and Point is %v\n", i, v)
		total = total + v
	}
	fmt.Printf("Total Points accumulated is %v\n", total)
}

func checkExistence(newMap map[string]int, key string) {
	val, ok := newMap[key]
	fmt.Println(val, ok)
}

func sumVariadic(nums ...int) (total int ) {
	total = 0
	for _, n := range nums {
		total = total + n
	}
	return
}

// error handling in go
// go treats errors as values rather than exceptions
// functions that can fail return a value and an error

func parseInt(s string) (int, error) {

	result, error := strconv.Atoi(s)
	if error != nil {
		return 0, fmt.Errorf("value passed must be a number")
	}
	if result < 1 || result > 5 {
		return 0, fmt.Errorf("result led to a num less than 0 and more than 5")
	}
	return result, nil
}

func run() {
	input := "300"

	res, err := parseInt(input)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

// defer simulation
func doSomething(success bool) error {

	fmt.Println("resource acquired")

	defer fmt.Println("cleanup: resource released") // runs at the end of function

	if !success {
		return errors.New("something went wrong")
	}
	fmt.Println("doing some important work")
	fmt.Println("work completed")
	return nil
}




func main() {
	// running example for error handling
	run()

	// defer example run
	fmt.Println("Case 1: success")
	if err := doSomething(true); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Case 1: failure")
	if err := doSomething(false); err != nil {
		fmt.Println("error:", err)
	}


    /* Using The spread operator */
	todos := []string{"learn golang", "watch anime", "solve puzzles"}
	moreTodo := []string{"eat cake"}

	spreadOperator(todos, moreTodo)

	// looping through a range of slices syntax
	points := []int{10, 24, 35, 40, 60}
	loopRange(points)

	// comma ok syntax to check existence of a key in a map
	newMap := map[string]int{
		"day1": 20,
		"day2": 30,
	}
	checkExistence(newMap, "day3")

	// variadic function + named return value
	total := sumVariadic(20, 40, 50, 40)
	fmt.Println(total)

	// anonymous or IIFE(immediately invoked function expression)
	result := func(a int , b int) int {
		return a + b
	}(5, 20)
	fmt.Printf("Result of IIFE is %v\n", result)

}