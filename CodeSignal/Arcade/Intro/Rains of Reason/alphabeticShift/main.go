package main

import "fmt"

func alphabeticShift(inputString string) string {
	finalString := make([]rune, 0)
	for _, value := range inputString {
		if value == 'z' {
			finalString = append(finalString, 'a')
		} else {
			finalString = append(finalString, value+1)
		}
	}
	return string(finalString)
}

func main() {
	fmt.Printf("%v", alphabeticShift("crazy"))
}
