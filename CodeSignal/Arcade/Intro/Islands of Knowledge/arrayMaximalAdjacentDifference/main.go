package main

import "fmt"

func arrayMaximalAdjacentDifference(inputArray []int) int {
	diff := make([]int, 0)
	for index := 0; index < len(inputArray)-1; index++ {
		diff = append(diff, Abs(inputArray[index]-inputArray[index+1]))
	}
	var max int = diff[0]
	for _, value := range diff {
		if value > max {
			max = value
		}
	}
	return max
}

func Abs(number int) int {
	if number < 0 {
		return number * -1
	}
	return number
}

func main() {
	fmt.Printf("%v", arrayMaximalAdjacentDifference([]int{10, 15, 20, 5}))
}
