package main

import "fmt"
import "strconv"

func sudoku(grid [][]int) bool {
    if grid[0][0] == grid[0][1] {
        return false
    }
    for y:=0; y<9;y++{
        sy:=0
        sx:=0
        s:=0
        for x:=0;x<9;x++{
            sy+=grid[y][x]
            sx+=grid[x][y]
            s+=grid[y-y%3+x/3][y%3*3+x%3]
        }
        if sy!= 45 || sx != 45 || s!= 45 {
            return false
        }
    }
    return true
}

func main() {
	fmt.Printf("%v", sudoku(3))
}
