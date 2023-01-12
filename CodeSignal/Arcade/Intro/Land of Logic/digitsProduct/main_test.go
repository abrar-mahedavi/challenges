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

func TestDigitsProduct(t *testing.T) {
	ans := digitsProduct(12)
	if ans != 26 {
		t.Errorf("digitsProduct(\"12\") = %v; want \"26\"", ans)
	}
}

func TestDigitsProductTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{12, 26},
		{19, -1},
		{450, 2559},
		{0, 10},
		{13, -1},
		{1, 1},
		{243, 399},
		{576, 889},
		{360, 589},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := longestWord(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
