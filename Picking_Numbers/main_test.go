package main

import (
	"reflect"
	"testing"
)

type testcase struct {
	arr []int32
	ans int32
}

func TestPickingNumbers(t *testing.T) {
	tests := []testcase{
		{arr: []int32{1, 2, 3, 4, 5}, ans: 2},
		{arr: []int32{1, 2, 1, 4, 90, 2, 6, 5, 6, 12, 2, 51, 1}, ans: 6},
		{arr: []int32{12, 14, 25, 13, 13, 11, 65}, ans: 3},
		{arr: []int32{9, 9, 9, 9, 9, 9}, ans: 6},
		{arr: []int32{4, 6, 5, 3, 3, 1}, ans: 3},
		{arr: []int32{1, 10, 20, 30, 40, 50, 60}, ans: 1},
	}

	for _, test := range tests {
		t.Run(string(test.arr), func(t *testing.T) {
			res := PickingNumbers(test.arr)
			if !reflect.DeepEqual(res, test.ans) {
				t.Errorf("Result: %d, want %d", res, test.ans)
			}
		})
	}
}


