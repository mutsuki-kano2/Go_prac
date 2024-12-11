package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		// 先頭と末尾を交換
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	str := "hello"
	fmt.Println(reverseString(str))
}
