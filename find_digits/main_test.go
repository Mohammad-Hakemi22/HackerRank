package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	number int32
	divisor int32
}

func TestFindDigits(t *testing.T) {
	tests := []TestCase{
		{number: 111, divisor: 3},
		{number: 124, divisor: 3},
		{number: 1012, divisor: 3},
		{number: 12, divisor: 2},
	}

		for _, tc := range tests {
			t.Run(string(tc.number), func(t *testing.T) {
				res := FindDigits(tc.number)
				if !reflect.DeepEqual(res, tc.divisor) {
					t.Errorf("Result: %d, want %d", res, tc.divisor)
				}
			})
		}
	
}