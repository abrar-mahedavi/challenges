package main

import "fmt"

func evenDigitsOnly(n int) bool {
	for n > 0 {
		rem := n % 10
		if rem%2 != 0 {
			return false
		}
		n = n / 10
	}
	return true
}

func main() {
	fmt.Printf("%v", evenDigitsOnly(12313))
}
