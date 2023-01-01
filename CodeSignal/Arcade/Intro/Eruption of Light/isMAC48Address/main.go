package main

import (
	"fmt"
	"strings"
)

func isMAC48Address(inputString string) bool {
	data := strings.Split(inputString, "-")
	// fmt.Println(len(data))
	if len(data) != 6 {
		return false
	}
	// return true
	return isValid(data)
}

func isValid(data []string) bool {
	for _, p := range data {
		if len(p) != 2 {
			return false
		}
		for _, b := range p {
			if (b >= '0' && b <= '9') || (b >= 'A' && b <= 'F') {
				continue
			}
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%v", isMAC48Address("bbbaacdafe"))
}
