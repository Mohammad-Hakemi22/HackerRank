package main

import (
	"fmt"
)

func birthday(s []int32, d int32, m int32) int32 {
	var sum,count, i int32 = 0, 0, 0
	sumArr := make([]int32, 1)
	for i = 0; i < m; i++ {
		sum += s[i]
	}
	sumArr[0] = sum
	for i := m; int(i) < len(s); i++ {
		sum = (sum - s[i - m]) + s[i]
		sumArr = append(sumArr, sum)
	}
	for _, v := range sumArr {
		if v == d {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(birthday([]int32{2, 2, 1, 3, 2}, 4, 2))
}
