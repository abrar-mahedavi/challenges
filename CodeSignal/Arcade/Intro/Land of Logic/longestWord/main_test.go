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

func TestLongestWord(t *testing.T) {
	ans := longestWord("Ready, steady, go!")
	if ans != "steady" {
		t.Errorf("longestWord(\"abcdc\") = %v; want \"steady\"", ans)
	}
}

func TestLongestWordTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"Ready[[[, steady, go!", "steady"},
		{"ABCd", "ABCd"},
		{"To be or not to be", "not"},
		{"abaa", "abaaba"},
		{"You are the best!!!!!!!!!!!! CodeFighter ever!", "CodeFighter"},
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
