package main

import "fmt"

func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for index, value := range inputArray {
		if value == elemToReplace {
			inputArray[index] = substitutionElem
		}
	}
	return inputArray
}

func main() {
	fmt.Printf("%v", arrayReplace([]int{1, 2, 3}, 1, 3))
}
