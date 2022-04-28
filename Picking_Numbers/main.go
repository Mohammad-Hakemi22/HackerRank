package main 


import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type realNumbers interface {
	constraints.Integer | constraints.Float
}


func maximum[T realNumbers](numbers []T) (T, T) {
	var max T = 0
	var max_ind T = 0

	for i, v := range numbers {
		if v > max {
			max = v
			max_ind = T(i)
		}
	}
	return max, max_ind
}


func pickingNumbers(a []int32) int32 {
	Frequency_data := make([]int32, 100)
	max_Frequency := make([]int32, 100)
	for i := 0; i < 100; i++ {
		for _, v := range a {
			if v == int32(i) {
				Frequency_data[i]++
			}
		}
	}
	for i := 0; i < len(Frequency_data)-1; i++ {
		max_Frequency[i] = Frequency_data[i] + Frequency_data[i+1]
	}
	max, _ := maximum(max_Frequency)
	return max
}

func main() {
	arr := []int32{4, 6, 5, 3, 3, 1}
	fmt.Println(pickingNumbers(arr))
}