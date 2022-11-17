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

func TestCircleOfNumbersSingel(t *testing.T) {
	ans := circleOfNumbers(10, 2)
	if ans != 7 {
		t.Errorf("circleOfNumbers(10,2) = %v; want 7", ans)
	}
}

func TestCircleOfNumbersTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{10, 2, 7},
		{10, 7, 2},
		{4, 1, 3},
		{6, 3, 0},
		{18, 6, 15},
		{12, 10, 4},
		{18, 5, 14},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := circleOfNumbers(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
