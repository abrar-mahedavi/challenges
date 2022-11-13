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

func TestAvoidObstaclesSingel(t *testing.T) {
	ans := avoidObstacles([]int{5, 3, 6, 7, 9})
	if ans != 4 {
		t.Errorf("avoidObstacles([]int{5, 3, 6, 7, 9}) = %v; want 4", ans)
	}
}

func TestAvoidObstaclesTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
	}{
		{[]int{5, 3, 6, 7, 9}, 4},
		{[]int{2, 3}, 4},
		{[]int{1, 4, 10, 6, 2}, 7},
		{[]int{1000, 999}, 6},
		{[]int{19, 32, 11, 23}, 3},
		{[]int{5, 8, 9, 13, 14}, 6},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := avoidObstacles(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
