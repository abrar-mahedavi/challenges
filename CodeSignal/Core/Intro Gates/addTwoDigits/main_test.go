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

func TestAddTwoDigitsSingle(t *testing.T) {
	ans := addTwoDigits(29)
	if ans != 11 {
		t.Errorf("addTwoDigits(29) = %v; want 11", ans)
	}
}

func TestAddTwoDigitsTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{29, 11},
		{48, 12},
		{10, 1},
		{25, 7},
		{52, 7},
		{99, 18},
		{44, 8},
		{50, 5},
		{39, 12},
		{26, 8},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := addTwoDigits(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
