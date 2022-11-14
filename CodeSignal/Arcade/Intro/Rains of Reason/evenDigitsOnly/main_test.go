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

func TestEvenDigitsOnlySingel(t *testing.T) {
	ans := evenDigitsOnly(12313)
	if ans != false {
		t.Errorf("evenDigitsOnly(12313) = %v; want false", ans)
	}
}

func TestEvenDigitsOnlyTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want bool
	}{
		{248622, true},
		{642386, false},
		{248842, true},
		{1, false},
		{8, true},
		{2462487, false},
		{468402800, true},
		{2468428, true},
		{5468428, false},
		{7468428, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := evenDigitsOnly(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
