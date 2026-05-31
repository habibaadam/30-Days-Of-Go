package main

import "fmt"

func main() {
    /* Using The spread operator */
	todos := []string{"learn golang", "watch anime", "solve puzzles"}
	moreTodo := []string{"eat cake"}

	spreadOperator(todos, moreTodo)

	// looping through a range of slices syntax
	points := []int{10, 24, 35, 40, 60}
	loopRange(points)

}

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