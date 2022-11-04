package main

import (
	"fmt"
	"testing"
)

func TestMakeArrayConsecutive2Singel(t *testing.T) {
	ans := makeArrayConsecutive2([]int{6, 2, 3, 8})
	if ans != 3 {
		t.Errorf("makeArrayConsecutive2([]int{6, 2, 3, 8}) = %d; want 3", ans)
	}
}

func TestMakeArrayConsecutive2TableDriven(t *testing.T) {
	var tests = []struct {
		a    []int
		want int
	}{
		{[]int{6, 2, 3, 8}, 3},
		{[]int{0, 3}, 2},
		{[]int{5, 4, 6}, 0},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := makeArrayConsecutive2(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
