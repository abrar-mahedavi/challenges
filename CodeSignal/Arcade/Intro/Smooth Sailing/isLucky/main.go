package main

import (
	"fmt"
	"strconv"
)

func isLucky(n int) bool {
	var numberToString string = strconv.Itoa(n)
	// fmt.Println(numberToString)
	firstHalf := numberToString[:len(numberToString)/2]
	secondHalf := numberToString[len(numberToString)/2 : len(numberToString)]

	// fmt.Println(string(firstHalf))
	// fmt.Println(string(secondHalf))

	firstNumber, _ := strconv.Atoi(firstHalf)
	secondNumber, _ := strconv.Atoi(secondHalf)
	// fmt.Println(sumOfDigits(firstNumber))
	// fmt.Println(sumOfDigits(secondNumber))
	if sumOfDigits(firstNumber) == sumOfDigits(secondNumber) {
		return true
	}
	return false
}

func sumOfDigits(n int) int {
	var sum int
	for n > 0 {
		var rem = n % 10
		n = n / 10
		sum = sum + rem
	}
	return sum
}

func main() {
	fmt.Printf("%v", isLucky(1230))
}
