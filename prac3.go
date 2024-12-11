package main

import (
	"fmt"
)

func findMin(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("配列が空です")
	}
	min := arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
	}
	return min, nil
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	min, err := findMin(arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("最小値は", min, "です")
}
