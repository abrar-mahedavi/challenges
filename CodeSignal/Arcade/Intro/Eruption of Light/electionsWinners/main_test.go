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

func TestElectionsWinnersSingel(t *testing.T) {
	ans := electionsWinners([]int{2, 3, 5, 2}, 3)
	if ans != 2 {
		t.Errorf("electionsWinners([]int{2, 3, 5, 2}, 3) = %v; want 2", ans)
	}
}

func TestElectionsWinnersTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		b    int
		want int
	}{
		{[]int{2, 3, 5, 2}, 3, 2},
		{[]int{1, 3, 3, 1, 1}, 0, 0},
		{[]int{5, 1, 3, 4, 1}, 0, 1},
		{[]int{1, 1, 1, 1}, 1, 4},
		{[]int{1, 1, 1, 1}, 0, 0},
		{[]int{3, 1, 1, 3, 1}, 2, 2},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := electionsWinners(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
