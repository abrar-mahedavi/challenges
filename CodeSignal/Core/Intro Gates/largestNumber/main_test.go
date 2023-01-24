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

func TestLargestNumberSingle(t *testing.T) {
	ans := largestNumber(2)
	if ans != 99 {
		t.Errorf("largestNumber(2) = %v; want 99", ans)
	}
}

func TestLargestNumberTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{2, 99},
		{1, 9},
		{7, 9999999},
		{4, 9999},
		{3, 999},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := largestNumber(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
