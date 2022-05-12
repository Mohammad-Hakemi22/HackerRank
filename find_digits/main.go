package main

import "fmt"

func FindDigits(n int32) int32 {
	var c int32 = 0
	digits := []int32{}
	val := n
	for val > 0 {
		digits = append(digits, val%10)
		val /= 10
	}
	for _, v := range digits {
		if v == 0 {
			continue
		}
		if n % v == 0 {
			c++
		}
	}
	return c
}

func main() {
	fmt.Println(FindDigits(123))
}