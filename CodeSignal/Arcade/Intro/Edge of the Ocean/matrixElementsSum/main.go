package main

import "fmt"

func matrixElementsSum(matrix [][]int) int {
	var sum int
	for row := 0; row < len(matrix[0]); row++ {
		for col := 0; col < len(matrix); col++ {
			if matrix[col][row] == 0 {
				break
			}
			sum = sum + matrix[col][row]
			// fmt.Printf("a[%d][%d] => %d ",row,col,value)
		}
	}
	return sum
}

func main() {
	fmt.Printf("%v", matrixElementsSum([][]int{{0, 1, 1, 2},
		{0, 5, 0, 0},
		{2, 0, 3, 3}}))
}
