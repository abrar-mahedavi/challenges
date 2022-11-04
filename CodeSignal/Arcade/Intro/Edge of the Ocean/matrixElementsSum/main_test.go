package main

import (
	"fmt"
	"testing"
)

func TestMatrixElementsSumSingel(t *testing.T) {
	ans := matrixElementsSum([][]int{{0, 1, 1, 2},
		{0, 5, 0, 0},
		{2, 0, 3, 3}})
	if ans != 9 {
		t.Errorf("matrixElementsSum([][]int{{0, 1, 1, 2},{0, 5, 0, 0},{2, 0, 3, 3}})) = %v; want 9", ans)
	}
}

func TestMatrixElementsSumTableDriven(t *testing.T) {
	var tests = []struct {
		a    [][]int
		want int
	}{
		{[][]int{{0, 1, 1, 2},
			{0, 5, 0, 0},
			{2, 0, 3, 3}}, 9},
		{[][]int{{1, 1, 1, 0},
			{0, 5, 0, 1},
			{2, 1, 3, 10}}, 9},
		{[][]int{{1, 1, 1},
			{2, 2, 2},
			{3, 3, 3}}, 18},
		{[][]int{{0}}, 0},
		{[][]int{{1, 0, 3},
			{0, 2, 1},
			{1, 2, 0}}, 5},
		{[][]int{{1},
			{5},
			{0},
			{3}}, 6},
		{[][]int{{1, 2, 3, 4, 5}}, 15},
		{[][]int{{2},
			{5},
			{10}}, 17},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%v ", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := matrixElementsSum(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
