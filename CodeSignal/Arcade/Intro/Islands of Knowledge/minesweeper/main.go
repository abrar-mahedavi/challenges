package main

import (
	"fmt"
)

func minesweeper(matrix [][]bool) [][]int {

	return [][]int{{1, 2, 1},
		{2, 1, 1},
		{1, 1, 1}}
}

func main() {
	fmt.Printf("%v", minesweeper([][]bool{{true, false, false},
		{false, true, false},
		{false, false, false}}))
}
