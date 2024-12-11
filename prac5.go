package main

import (
	"fmt"
)

func mergeArrays(arr1, arr2 []int) []int {
	return append(arr1, arr2...)
}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	merged := mergeArrays(arr1, arr2)
	fmt.Println("マージされた配列:", merged)
}
