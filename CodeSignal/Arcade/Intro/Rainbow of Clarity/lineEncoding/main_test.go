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

func TestLineEncodingSingel(t *testing.T) {
	ans := lineEncoding("0")
	if ans != "0" {
		t.Errorf("lineEncoding(\"0\") = %v; want true", ans)
	}
}

func TestLineEncodingTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"aabbbc", "2a3bc"},
		{"abbcabb", "a2bca2b"},
		{"abcd", "abcd"},
		{"zzzz", "4z"},
		{"wwwwwwwawwwwwww", "7wa7w"},
		{"ccccccccccccccc", "15c"},
		{"qwertyuioplkjhg", "qwertyuioplkjhg"},
		{"ssiiggkooo", "2s2i2gk3o"},
		{"adfaaa", "adf3a"},
		{"bbjaadlkjdl", "2bj2adlkjdl"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := lineEncoding(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
