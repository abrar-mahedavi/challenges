package main

import "fmt"

func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	if maxW >= weight1+weight2 {
		return (value1 + value2)
	} else if value2 > value1 && maxW >= weight2 {
		return value2
	} else if value1 > value2 && maxW >= weight1 {
		return value1
	} else if maxW >= weight1 {
		return value1
	} else if maxW >= weight2 {
		return value2
	}
	return 0
}

func main() {
	fmt.Printf("%v", knapsackLight(10, 5, 6, 4, 8))
}
