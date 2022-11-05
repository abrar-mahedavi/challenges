package main

import (
	"fmt"
	"testing"
)

func TestCommonCharacterCountSingel(t *testing.T) {
	ans := commonCharacterCount("aabc", "aabc")
	if ans != 4 {
		t.Errorf("commonCharacterCount(\"aabc\",\"aabc\") = %v; want 4", ans)
	}
}

func TestAbsoluteValuesSumMinimizationTableDriven(t *testing.T) {
	var tests = []struct {
		a    string
		b    string
		want int
	}{
		{"aabcc", "adcaa", 3},
		{"zzzz", "zzzzzzz", 4},
		{"abca", "xyzbac", 3},
		{"a", "b", 0},
		{"a", "aaa", 1},
		{"1", "1", 1},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := commonCharacterCount(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
