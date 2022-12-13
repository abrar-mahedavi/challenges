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

func TestGrowingPlantSingel(t *testing.T) {
	ans := growingPlant(100, 10, 910)
	if ans != 10 {
		t.Errorf("growingPlant(100, 10, 910) = %v; want 10", ans)
	}
}

func TestGrowingPlantTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		c    int
		want int
	}{
		{100, 10, 910, 10},
		{10, 9, 4, 1},
		{5, 2, 7, 2},
		{7, 3, 443, 110},
		{6, 5, 10, 5},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v %v", tt.a, tt.b, tt.c)
		t.Run(testname, func(t *testing.T) {
			ans := growingPlant(tt.a, tt.b, tt.c)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
