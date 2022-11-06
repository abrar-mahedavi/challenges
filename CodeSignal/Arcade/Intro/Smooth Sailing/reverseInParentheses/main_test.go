package main

import (
	"fmt"
	"testing"
)

func TestReverseInParenthesesSingel(t *testing.T) {
	ans := reverseInParentheses("(bar)")
	if ans != "bar" {
		t.Errorf("reverseInParentheses(\"(bar)\") = %v; want bar", ans)
	}
}

func TestReverseInParenthesesTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"(bar)", "bar"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := isLucky(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
