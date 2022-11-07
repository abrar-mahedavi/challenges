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

func TestAlternatingSumsSingel(t *testing.T) {
	ans := alternatingSums([]int{1, 2, 3})
	if !Equal(ans, []int{4, 2}) {
		t.Errorf("alternatingSums([]int{1, 2, 3}) = %v; want [4,2]", ans)
	}
}

func TestAlternatingSumsTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{4, 2}},
		{[]int{50, 60, 60, 45, 70}, []int{180, 105}},
		{[]int{100, 50}, []int{100, 50}},
		{[]int{80}, []int{80, 0}},
		{[]int{100, 50, 50, 100}, []int{150, 150}},
		{[]int{100, 51, 50, 100}, []int{150, 151}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := alternatingSums(tt.a)
			if !Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
