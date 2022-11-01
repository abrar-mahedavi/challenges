package main

import (
	"fmt"
	"testing"
)

func TestAddSingel(t *testing.T) {
	ans := Add(2, 2)
	if ans != 4 {
		t.Errorf("IntMin(2) = %d; want 5", ans)
	}
}

func TestAddTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{-3, 3, 0},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Add(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
