package main

import "fmt"
import "regexp"

func differentSquares(matrix [][]int) int {
    m := make(map[string]int)
    for y:=0; y<len(matrix) - 1; y++ {
        for x:=0; x < len(matrix[y]) - 1; x++ {
            k:=fmt.Sprintf("%v.%v.%v.%v", matrix[y][x], matrix[y+1][x], matrix[y][x+1], matrix[y+1][x+1])
            m[k] ++
        }
    }
    return len(m)
}


func main() {
	fmt.Printf("%v", differentSquares([][]int{{1,2,1},
                                            {2,2,2},
                                            {2,2,2},
                                            {1,2,3},
                                            {2,2,1}})
}
