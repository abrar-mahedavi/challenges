package main

import "fmt"

func depositProfit(deposit int, rate int, threshold int) int {
	var sum float64 = float64(deposit)
	index := 0
	for ; sum <= float64(threshold); index++ {
		sum = sum + (sum * float64(rate) / 100.0)
		// fmt.Println(sum)
		if sum >= float64(threshold) {
			break
		}
	}
	return index + 1
}

func main() {
	fmt.Printf("%v", depositProfit(100, 20, 170))
}
