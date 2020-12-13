package main

import (
	"fmt"

	"github.com/clnbs/karateChop/internal/app/binarychop"
)

func main() {
	// Canonical example to use recursive binary chop
	tab := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chop := binarychop.NewBinaryChop(&binarychop.Recursive{})
	valueToFind := 8
	index := chop.Execute(valueToFind, tab)
	if index == -1 {
		fmt.Println("no value found")
	}
	fmt.Println("value", valueToFind, "found at index", index)
}
