package main

import (
	"fmt"
	"testing"
)

func TestShapeAreaSingel(t *testing.T) {
	ans := ShapeArea(2)
	if ans != 5 {
		t.Errorf("IntMin(2) = %d; want 5", ans)
	}
}

func TestShapeAreaTableDriven(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{1, 1},
		{2, 5},
		{3, 13},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := ShapeArea(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
