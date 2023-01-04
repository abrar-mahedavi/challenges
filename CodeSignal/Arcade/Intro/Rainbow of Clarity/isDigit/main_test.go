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

func TestIsDigitSingel(t *testing.T) {
	ans := isDigit("0")
	if ans != true {
		t.Errorf("isDigit(\"0\") = %v; want true", ans)
	}
}

func TestIsDigitTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"0", true},
		{"Z", false},
		{"-", false},
		{"O", false},
		{"1", true},
		{"2", true},
		{"@", false},
		{"+", false},
		{"6", true},
		{"(", false},
		{")", false},
		{"!", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := isDigit(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
