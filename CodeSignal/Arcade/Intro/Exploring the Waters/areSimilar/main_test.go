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

func TestAreSimilarSingel(t *testing.T) {
	ans := areSimilar([]int{1, 2, 3}, []int{1, 2, 3})
	if ans != true {
		t.Errorf("areSimilar([]int{1, 2, 3}, []int{1, 2, 3}) = %v; want true", ans)
	}
}

func TestAddBorderTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		b    []int
		want bool
	}{
		{[]int{1, 2, 3},
			[]int{1, 2, 3},
			true},
		{[]int{1, 2, 3},
			[]int{2, 1, 3},
			true},
		{[]int{1, 2, 2},
			[]int{2, 1, 1},
			false},
		{[]int{1, 1, 4},
			[]int{1, 2, 3},
			false},
		{[]int{1, 2, 3},
			[]int{1, 10, 2},
			false},
		{[]int{2, 3, 9},
			[]int{10, 3, 2},
			false},
		{[]int{4, 6, 3},
			[]int{3, 4, 6},
			false},
		{[]int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279,
			[]int{832, 998, 148, 570, 533, 561, 455, 147, 894, 279},
			true},
		{[]int{832, 998, 148, 570, 533, 561, 894, 147, 455, 279},
			[]int{832, 570, 148, 998, 533, 561, 455, 147, 894, 279,
			false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := areSimilar(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
