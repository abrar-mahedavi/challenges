package main

import "fmt"

func circleOfNumbers(n int, firstNumber int) int {
	if firstNumber < int(n/2) {
		return firstNumber + int(n/2)
	} else {
		return (firstNumber + int(n/2)) % n
	}
}

func main() {
	fmt.Printf("%v", circleOfNumbers(10, 2))
}
