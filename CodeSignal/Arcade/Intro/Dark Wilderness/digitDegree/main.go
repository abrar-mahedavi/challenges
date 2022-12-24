package main

import "fmt"

func sumDigit(n int) int {
	s := 0
	for n > 0 {
		s += n % 10
		n = n / 10
	}
	return s
}

func digitDegree(n int) int {
	i := 0
	for n > 9 {
		n = sumDigit(n)
		i++
	}
	return i
}

func main() {
	fmt.Printf("%v", digitDegree(100))
}
