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

func TestKnapsackLightSingel(t *testing.T) {
	ans := knapsackLight(10, 5, 4, 6, 8)
	if ans != 10 {
		t.Errorf("knapsackLight(10, 5, 4, 6, 8) = %v; want 10", ans)
	}
}

func TestKnapsackLightTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		c    int
		d    int
		e    int
		want int
	}{
		{10, 5, 4, 6, 8, 10},
		{10, 5, 6, 4, 9, 16},
		{5, 3, 7, 4, 6, 7},
		{10, 2, 11, 3, 1, 0},
		{15, 2, 20, 3, 2, 15},
		{2, 5, 3, 4, 5, 3},
		{4, 3, 3, 4, 4, 4},
		{3, 5, 3, 8, 10, 3},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v %v %v %v", tt.a, tt.b, tt.c, tt.d, tt.e)
		t.Run(testname, func(t *testing.T) {
			ans := knapsackLight(tt.a, tt.b, tt.c, tt.d, tt.e)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
