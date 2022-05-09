package main

import (
	"fmt"
)

func countingValleys(steps int32, path string) int32 {
	var count, v_count int32 = 0, 0
	for i, v := range path {
		if v == rune('U') {
			count++
		} else if v == rune('D') {
			count--
		}
		if count == 0 && rune(path[i]) == 'U'{
			v_count++
		}
	}
	return v_count
}


func main() {
	fmt.Println(countingValleys(8, "DDUUUUDD"))
}

// if v == rune('U') && count != 0 {
// 	count++
// } else if v == rune('D') {
// 	count--
// }
// if count == 0 {
// 	v_count++
// }
