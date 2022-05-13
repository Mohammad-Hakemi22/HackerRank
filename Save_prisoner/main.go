package main

import (
	"fmt"
)

func saveThePrisoner(n int32, m int32, s int32) int32 {
	return int32(1 + (s - 1 + m - 1) % n)

}


func main() {
	fmt.Println(saveThePrisoner(654809340, 204894365, 472730208))
}
