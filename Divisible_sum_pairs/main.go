package main

import (
	"fmt"
)

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var i, count int32 = 0, 0
	for i = 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}
