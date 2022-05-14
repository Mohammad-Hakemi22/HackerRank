package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
    var sum int32 = 0
    for i := 0; i < len(keyboards); i++ {
        for j := 0; j < len(drives); j++ {
            price := keyboards[i] + drives[j]
            if sum < price && price <= b {
                sum = price
            }
        }
    }
    if sum == 0 {
        return -1
    } else {
        return sum
    }
}

func main() {
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 2))
}
