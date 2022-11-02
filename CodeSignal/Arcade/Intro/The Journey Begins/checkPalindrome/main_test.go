package main

import (
	"fmt"
	"testing"
)

func TestCheckPalindromeSingel(t *testing.T) {
	ans := checkPalindrome("abba")
	if ans != true {
		t.Errorf("checkPalindrome(abba) = %v; want true", ans)
	}
}

func TestCheckPalindromeTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want bool
	}{
		{"abba", true},
		{"abcba", true},
		{"test", false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := checkPalindrome(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
