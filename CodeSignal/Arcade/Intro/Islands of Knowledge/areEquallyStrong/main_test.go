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

func TestAreEquallyStrongSingel(t *testing.T) {
	ans := areEquallyStrong(10, 15, 15, 10)
	if ans != true {
		t.Errorf("areEquallyStrong(10,15,15,10) = %v; want true", ans)
	}
}

func TestAreEquallyStrongTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		c    int
		d    int
		want bool
	}{
		{10, 15, 15, 10, true},
		{15, 10, 15, 10, true},
		{15, 10, 15, 9, false},
		{10, 5, 5, 10, true},
		{10, 15, 5, 20, false},
		{10, 20, 10, 20, true},
		{5, 20, 20, 5, true},
		{20, 15, 5, 20, false},
		{5, 10, 5, 10, true},
		{1, 10, 10, 0, false},
		{5, 5, 10, 10, false},
		{10, 5, 10, 6, false},
		{1, 1, 1, 1, true},
		{0, 10, 0, 10, true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v %v %v", tt.a, tt.b, tt.c, tt.d)
		t.Run(testname, func(t *testing.T) {
			ans := areEquallyStrong(tt.a, tt.b, tt.c, tt.d)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
