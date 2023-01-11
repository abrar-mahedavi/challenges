package main

import (
	"fmt"
	"testing"
)

func TestDifferentSquaresSingel(t *testing.T) {
	ans := differentSquares([][]int{{1,2,1},
                                  {2,2,2},
                                  {2,2,2},
                                  {1,2,3},
                                  {2,2,1}})
	if ans != 9 {
		t.Errorf("differentSquares([][]int{{1,2,1},{2,2,2},{2,2,2},{1,2,3},{2,2,1}})) = %v; want 6", ans)
	}
}

func TestDifferentSquaresTableDriven(t *testing.T) {
	var tests = []struct {
		a    [][]int
		want int
	}{
		{[][]int{{9,9,9,9,9},
             {9,9,9,9,9},
             {9,9,9,9,9},
             {9,9,9,9,9},
             {9,9,9,9,9},
             {9,9,9,9,9}}, 1},
		{[][]int{{3}}, 0},
		{[][]int{{3,4,5,6,7,8,9}}, 0},
		{[][]int{{3},
             {4},
             {5},
             {6},
             {7}}, 0},
		{[][]int{{2,5,3,4,3,1,3,2},
             {4,5,4,1,2,4,1,3},
             {1,1,2,1,4,1,1,5},
             {1,3,4,2,3,4,2,4},
             {1,5,5,2,1,3,1,1},
             {1,2,3,3,5,1,2,4},
             {3,1,4,4,4,1,5,5},
             {5,1,3,3,1,5,3,5},
             {5,4,4,3,5,4,4,4}}, 54},
		{[][]int{{1,1,1,1,1,1,2,2,2,3,3,3,9,9,9,2,3,9}}, 0},

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
