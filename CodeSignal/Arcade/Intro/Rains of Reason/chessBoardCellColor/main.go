package main

import "fmt"

func chessBoardCellColor(cell1 string, cell2 string) bool {
	var A1 rune = rune(cell1[0])
	var B1 rune = rune(cell2[0])

	var A2 rune = rune(cell1[1])
	var B2 rune = rune(cell2[1])
	return Abs(A1-B1)%2 == Abs(A2-B2)%2
}

func Abs(n rune) int {
	n1 := int(n)
	if n1 < 0 {
		return n1 * -1
	}
	return n1
}

func main() {
	fmt.Printf("%v", chessBoardCellColor("A1", "C3"))
}
