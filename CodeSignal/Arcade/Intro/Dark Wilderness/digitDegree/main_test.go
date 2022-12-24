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

func TestDigitDegreeSingel(t *testing.T) {
	ans := digitDegree(100)
	if ans != 1 {
		t.Errorf("digitDegree(100) = %v; want 1", ans)
	}
}

func TestDigitDegreeTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{5, 0},
		{100, 1},
		{91, 2},
		{99, 2},
		{1000000000, 1},
		{9, 0},
		{73, 2},
		{877, 2},
		{777773, 3},
		{304, 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := digitDegree(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
