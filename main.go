package main

import (
	"algorithm/algosort"
	"fmt"
)

func main() {
	sequenceSlice := []int{3, 4, 5, 6, 1, 10, 2, 9, 8}
	// bubble sort
	bubble := algosort.Bubble(sequenceSlice)
	fmt.Println(bubble)
}
