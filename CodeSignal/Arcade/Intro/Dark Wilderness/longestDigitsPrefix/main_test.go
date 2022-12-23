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

func TestLongestDigitsPrefixSingel(t *testing.T) {
	ans := longestDigitsPrefix("123aa1")
	if ans != "123" {
		t.Errorf("longestDigitsPrefix(\"123aa1\") = %v; want \"123\" ", ans)
	}
}

func TestLongestDigitsPrefixTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"123aa1", "123"},
		{"0123456789", "0123456789"},
		{"  3) always check for whitespaces", ""},
		{"12abc34", "12"},
		{"the output is 42", ""},
		{"aaaaaaa", ""},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := longestDigitsPrefix(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestLongestDigitsPrefix1TableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"123aa1", "123"},
		{"0123456789", "0123456789"},
		{"  3) always check for whitespaces", ""},
		{"12abc34", "12"},
		{"the output is 42", ""},
		{"aaaaaaa", ""},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := longestDigitsPrefix1(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
