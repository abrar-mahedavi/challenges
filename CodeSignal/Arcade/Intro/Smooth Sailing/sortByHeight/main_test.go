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

func TestSortByHeightSingel(t *testing.T) {
	ans := sortByHeight([]int{-1, 150, 190, 170, -1, -1, 160, 180})
	if !Equal(ans, []int{-1, 150, 160, 170, -1, -1, 180, 190}) {
		t.Errorf("sortByHeight([]int{-1, 150, 190, 170, -1, -1, 160, 180}) = %v; want [-1, 150, 160, 170, -1, -1, 180, 190]", ans)
	}
}

func TestSortByHeightTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		want []int
	}{
		{[]int{-1, 150, 190, 170, -1, -1, 160, 180}, []int{-1, 150, 160, 170, -1, -1, 180, 190}},
		{[]int{-1, -1, -1, -1, -1}, []int{-1, -1, -1, -1, -1}},
		{[]int{-1}, []int{-1}},
		{[]int{4, 2, 9, 11, 2, 16}, []int{2, 2, 4, 9, 11, 16}},
		{[]int{2, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1}, []int{1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 2}},
		{[]int{23, 54, -1, 43, 1, -1, -1, 77, -1, -1, -1, 3}, []int{1, 3, -1, 23, 43, -1, -1, 54, -1, -1, -1, 77}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := sortByHeight(tt.a)
			if !Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
