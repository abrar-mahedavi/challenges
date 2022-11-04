package main

import (
	"fmt"
	"testing"
)

func TestAlmostIncreasingSequenceSingel(t *testing.T) {
	ans := almostIncreasingSequence([]int{1, 3, 2, 1})
	if ans != false {
		t.Errorf("almostIncreasingSequence([]int{1, 3, 2, 1}) = %v; want false", ans)
	}
}

func TestAlmostIncreasingSequenceTableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		want bool
	}{
		{[]int{1, 3, 2, 1}, false},
		{[]int{1, 3, 2}, true},
		{[]int{1, 2, 1, 2}, false},
		{[]int{3, 6, 5, 8, 10, 20, 15}, false},
		{[]int{1, 1, 2, 3, 4, 4}, false},
		{[]int{1, 4, 10, 4, 2}, false},
		{[]int{10, 1, 2, 3, 4, 5}, true},
		{[]int{1, 1, 1, 2, 3}, false},
		{[]int{0, -2, 5, 6}, true},
		{[]int{1, 2, 3, 4, 5, 3, 5, 6}, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := almostIncreasingSequence(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
