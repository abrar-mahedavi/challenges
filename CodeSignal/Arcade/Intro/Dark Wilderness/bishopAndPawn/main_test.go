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

func TestBishopAndPawnSingel(t *testing.T) {
	ans := bishopAndPawn("a1", "c3")
	if ans != true {
		t.Errorf("bishopAndPawn(\"a1\",\"c3\") = %v; want true", ans)
	}
}

func TestBishopAndPawnTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		b    string
		want bool
	}{
		{"a1", "c3", true},
		{"h1", "h3", false},
		{"a5", "c3", true},
		{"g1", "f3", false},
		{"e7", "d6", true},
		{"e7", "a3", true},
		{"e3", "a7", true},
		{"a1", "h8", true},
		{"a1", "h7", false},
		{"h1", "a8", true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := bishopAndPawn(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
