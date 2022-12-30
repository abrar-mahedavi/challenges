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

func TestBuildPalindromeSingel(t *testing.T) {
	ans := buildPalindrome("abcdc")
	if ans != "abcdcba" {
		t.Errorf("buildPalindrome(\"abcdc\") = %v; want \"abcdcba\"", ans)
	}
}

func TestBuildPalindromeTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"abcdc", "abcdcba"},
		{"ababab", "abababa"},
		{"abba", "abba"},
		{"abaa", "abaaba"},
		{"aaaaba", "aaaabaaaa"},
		{"abc", "abcba"},
		{"kebab", "kebabek"},
		{"abccba", "abccba"},
		{"abcabc", "abcabcbacba"},
		{"cbdbedffcg", "cbdbedffcgcffdebdbc"},
		{"euotmn", "euotmnmtoue"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := buildPalindrome(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
