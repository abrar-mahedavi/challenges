package main

import (
	"fmt"
	"testing"
)

func Equal(a, b []int) bool {
	// fmt.Printf("%T %T", a, b)
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		// fmt.Printf("v => %v b => %v\n", v, b[i])
		//fmt.Printf("v(Type) => %T b(Type) => %T\n", v, b[i])
		// fmt.Printf("v != b[i] => %v \n", v != b[i])
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestDeleteDigitSingel(t *testing.T) {
	ans := deleteDigit(1001)
	if ans != 101 {
		t.Errorf("deleteDigit(\"0\") = %v; want true", ans)
	}
}

func TestDeleteDigitTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{152, 52},
		{1001, 101},
		{10, 1},
		{222219, 22229},
		{109, 19},
		{222250, 22250},
		{44435, 4445},
		{12, 2},
		{218616, 28616},
		{861452, 86452},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := deleteDigit(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
