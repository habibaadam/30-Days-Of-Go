package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// using generics on structs
type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID int
	Name string
	Data T
}

// using generics on maps
type CustomMap[T comparable, v int | string] map[T]v


func Add[T constraints.Ordered](a T, b T) T {
	return a + b
}

// input - array of integers eg [1, 2, 3] => (func)
// output after func - array of integers eg [2, 4, 6]
// initially works with only integers => func MapValues(values []int, mapFunc func(int) int) []int
// works with multiple types and not just integers after defining a generic
func MapValues[T constraints.Ordered] (values []T, mapFunc func(T) T) []T {
	var newVals []T

	for _, val := range values {
		newVal := mapFunc(val)
		newVals = append(newVals, newVal)
	}
	return newVals
}

func main() {

	u := User[int]{ //flexibility for data to be anything.
		ID: 0,
		Name: "Habie",
		Data: 4,
	}
	fmt.Println(u)

	m := make(CustomMap[int, string])
	m[3] = "3"

	result := MapValues([]float64{1.3, 2.4, 3.4}, func(n float64 ) float64 {
		return n * 3
	})
	fmt.Printf("result: %v\n", result)

	//res := Add(1.2, 2.2)
	//fmt.Printf("result: %v\n", res)
}