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

func TestArrayReplaceSingel(t *testing.T) {
	ans := arrayReplace([]int{1, 2, 3}, 1, 3)
	if !Equal(ans, []int{3, 2, 3}) {
		t.Errorf("areSimilar([]int{1, 2, 3}, 1, 3) = %v; want true", ans)
	}
}

func TestArrayReplaceTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		b    int
		c    int
		want []int
	}{
		{[]int{1, 2, 3}, 1, 3, []int{3, 2, 3}},
		{[]int{1, 2, 1}, 1, 3, []int{3, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 3, 0, []int{1, 2, 0, 4, 5}},
		{[]int{1, 1, 1}, 1, 10, []int{10, 10, 10}},
		{[]int{1, 2, 1, 2, 1}, 2, 1, []int{1, 1, 1, 1, 1}},
		{[]int{1, 2, 1, 2, 1}, 2, 2, []int{1, 2, 1, 2, 1}},

		{[]int{3, 1}, 3, 9, []int{9, 1}},
		{[]int{10, 10}, 0, 9, []int{10, 10}},
		{[]int{2, 1}, 3, 9, []int{2, 1}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v %v", tt.a, tt.b, tt.c)
		t.Run(testname, func(t *testing.T) {
			ans := arrayReplace(tt.a, tt.b, tt.c)
			if !Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
