package main

import (
	"fmt"
)

func average(arr []int) (float64, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("配列が空です")
	}
	total := 0
	for _, num := range arr {
		total += num
	}
	return float64(total) / float64(len(arr)), nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	avg, err := average(arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("平均値は", avg, "です")
}
