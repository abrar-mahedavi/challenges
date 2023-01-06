package main

import (
	"fmt"
)

func chessKnight(cell string) int {
	let := cell[0]
	num := cell[1] - '0'

	ret := 8
	if let == 'a' || let == 'h' {
		ret = ret / 2
		// fmt.Print(1, ret)
	}
	if let == 'b' || let == 'g' {
		ret = ret - 2
		// fmt.Print(2, ret)
	}
	if num == 1 || num == 8 {
		ret = ret / 2
		// fmt.Print(3, ret)
	}
	if num == 2 || num == 7 {
		ret = ret - 2
		// fmt.Print(4, ret)
	}
	return ret
}

func main() {
	fmt.Printf("%v", chessKnight("b1"))
}
