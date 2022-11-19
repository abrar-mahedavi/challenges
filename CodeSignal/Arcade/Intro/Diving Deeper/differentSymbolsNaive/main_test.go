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

func TestDifferentSymbolsNaiveSingel(t *testing.T) {
	ans := differentSymbolsNaive("adsda")
	if ans != 3 {
		t.Errorf("differentSymbolsNaive(\"adsda\") = %v; want 3", ans)
	}
}

func TestDifferentSymbolsNaiveTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want int
	}{
		{"asd", 3},
		{"aba", 2},
		{"ccccccccccc", 1},
		{"bcaba", 3},
		{"codesignal", 10},
		{"", 0},
		{" ", 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := differentSymbolsNaive(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
