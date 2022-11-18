package main

import "fmt"

func firstDigit(inputString string) string {
	for index, value := range inputString {
		if value > 47 && value < 58 {
			return string(inputString[index])
		}
	}
	return string(" ")
}

func main() {
	fmt.Printf("%v", firstDigit("var_1__Int"))
}
