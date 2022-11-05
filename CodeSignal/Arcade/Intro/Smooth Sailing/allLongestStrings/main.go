package main

import (
	"fmt"
	"math"
)

func allLongestStrings(inputArray []string) []string {
	mapOfLongestStrings := make(map[int][]string)
	for _, value := range inputArray {
		if _, ok := mapOfLongestStrings[len(value)]; ok {
			getData := mapOfLongestStrings[len(value)]
			getData = append(getData, value)
			mapOfLongestStrings[len(value)] = getData
		} else {
			newData := []string{}
			newData = append(newData, value)
			mapOfLongestStrings[len(value)] = newData
		}
	}
	maxNumber := math.MinInt32
	for n := range mapOfLongestStrings {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return mapOfLongestStrings[maxNumber]
}

func main() {
	fmt.Printf("%v", allLongestStrings([]string{"aba",
		"aa",
		"ad",
		"vcd",
		"aba"}))
}
