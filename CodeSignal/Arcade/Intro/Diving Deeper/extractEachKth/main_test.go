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

func TestExtractEachKthSingel(t *testing.T) {
	ans := extractEachKth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	if !Equal(ans, []int{1, 2, 4, 5, 7, 8, 10}) {
		t.Errorf("extractEachKth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3) = %v; want [1, 2, 4, 5, 7, 8, 10]", ans)
	}
}

func TestExtractEachKthTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		b    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, []int{1, 2, 4, 5, 7, 8, 10}},
		{[]int{1, 1, 1, 1, 1}, 1, []int{}},
		{[]int{1, 2, 1, 2, 1, 2, 1, 2}, 2, []int{1, 1, 1, 1}},
		{[]int{1, 2, 1, 2, 1, 2, 1, 2}, 10, []int{1, 2, 1, 2, 1, 2, 1, 2}},
		{[]int{2, 4, 6, 8, 10}, 2, []int{2, 6, 10}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := extractEachKth(tt.a, tt.b)
			if !Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
