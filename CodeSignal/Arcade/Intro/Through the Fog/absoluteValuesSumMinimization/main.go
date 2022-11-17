package main

import (
	"fmt"
	"sort"
)

// TODO :- debug why one of the test is failing
func absoluteValuesSumMinimization(a []int) int {
	mapSummation := make(map[int]int)
	for index := range a {
		var sum int = 0
		for _, value := range a {
			sum = sum + Abs(a[index]-value)
		}
		mapSummation[a[index]] = sum

	}
	minSum := mapSummation[a[0]]
	var key int = a[0]
	for key1, value := range mapSummation {
		if value < minSum {
			minSum = value
			key = key1
		}
	}
	return key

}

func absoluteValuesSumMinimization1(a []int) int {
	sort.Ints(a)
	return a[(len(a)-1)/2]
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Printf("%v", absoluteValuesSumMinimization([]int{2, 4, 7}))
}
