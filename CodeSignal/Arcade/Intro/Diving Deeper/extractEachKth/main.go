package main

import "fmt"

func extractEachKth(inputArray []int, k int) []int {
	output := make([]int, 0)
	for index, value := range inputArray {
		// fmt.Println(index == k-1)
		if (index+1)%k == 0 {
			continue
		}
		output = append(output, value)
	}
	return output
}

func main() {
	fmt.Printf("%v", extractEachKth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
}
