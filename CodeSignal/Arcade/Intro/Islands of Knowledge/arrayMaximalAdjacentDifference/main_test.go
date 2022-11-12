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

func TestArrayMaximalAdjacentDifferenceSingel(t *testing.T) {
	ans := arrayMaximalAdjacentDifference([]int{10, 15, 15, 10})
	if ans != 5 {
		t.Errorf("arrayMaximalAdjacentDifference([]int{10,15,15,10}) = %v; want 5", ans)
	}
}

func TestArrayMaximalAdjacentDifferenceTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
	}{
		{[]int{2, 4, 1, 0}, 3},
		{[]int{1, 1, 1, 1}, 0},
		{[]int{-1, 4, 10, 3, -2}, 7},
		{[]int{10, 11, 13}, 2},
		{[]int{-2, -2, -2, -2, -2}, 0},
		{[]int{-1, 1, -3, -4}, 4},
		{[]int{-14, 15, -15}, 30},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := arrayMaximalAdjacentDifference(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
