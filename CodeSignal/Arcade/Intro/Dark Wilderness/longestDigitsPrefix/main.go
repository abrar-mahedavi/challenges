package main

import "fmt"

func longestDigitsPrefix(inputString string) string {
	var i int
	for ; i < len(inputString); i++ {
		if inputString[i] < '0' || inputString[i] > '9' {
			break
		}
	}
	return inputString[:i]
}

// TODO :- Please check the logic and update it
func longestDigitsPrefix1(inputString string) string {
	var i int
	var temp string = ""
	var tempOfTemp string = ""
	for ; i < len(inputString); i++ {

		if inputString[i] < '0' || inputString[i] > '9' {
			tempOfTemp = temp
			continue
		}
		if len(tempOfTemp) != len(temp) {
			tempOfTemp = temp
		}
		temp += string(inputString[i])
	}
	return tempOfTemp
}

func main() {
	fmt.Printf("%v\n", longestDigitsPrefix("12323ab34413433434232"))
	fmt.Printf("%v\n", longestDigitsPrefix1("12323ab34413433434232"))
}
