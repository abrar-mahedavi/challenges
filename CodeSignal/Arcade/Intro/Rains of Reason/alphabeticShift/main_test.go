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

func TestAlphabeticShiftSingel(t *testing.T) {
	ans := alphabeticShift("crazy")
	if ans != "dsbaz" {
		t.Errorf("alphabeticShift(\"crazy\") = %v; want dsbaz", ans)
	}
}

func TestAlphabeticShiftTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"crazy", "dsbaz"},
		{"z", "a"},
		{"aaaabbbccd", "bbbbcccdde"},
		{"fuzzy", "gvaaz"},
		{"codesignal", "dpeftjhobm"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := alphabeticShift(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
