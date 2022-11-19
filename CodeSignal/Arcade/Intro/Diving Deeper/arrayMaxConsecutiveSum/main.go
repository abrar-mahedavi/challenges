package main

import "fmt"

func arrayMaxConsecutiveSum(inputArray []int, k int) int {
	var sum int = 0
	for index := 0; index <= len(inputArray)-k; index++ {
		var sumTemp int
		for inner := index; inner < (index + k); inner++ {
			sumTemp = sumTemp + inputArray[inner]
		}
		if sumTemp > sum {
			sum = sumTemp
		}
		// fmt.Println(sumTemp)
	}
	return sum
}

func main() {
	fmt.Printf("%v", arrayMaxConsecutiveSum([]int{2, 3, 5, 1, 6}, 2))
}
