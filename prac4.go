package main

import (
	"fmt"
	"math"
)

func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(math.Abs(float64(n)))) + 1
}

func main() {
	num := -12345
	fmt.Printf("%d の桁数は %d 桁です\n", num, countDigits(num))
}
