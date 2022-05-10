package main

import (
	"fmt"
)

func designerPdfViewer(h []int32, word string) int32 {
    words := make([]int32, len(word))
    words_id := make([]int32, len(word))
    for i, _ := range words {
        words[i] = rune(word[i]) - 97
    }
	for i, _ := range words {
        words_id[i] = h[words[i]]
    }
    max := max(words_id)
    return max * int32(len(words_id))

}

func max(arr []int32) int32 {
    max := arr[0]
    for _, v := range arr {
        if v > max {
            max = v
        }
    }
    return max
}

func main() {
	fmt.Println(designerPdfViewer([]int32{1 ,3 ,1 ,3 ,1 ,4 ,1 ,3 ,2 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5 ,5}, "abc"))
}