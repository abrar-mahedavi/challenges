package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isIPv4Address(inputString string) bool {
	addresses := strings.Split(inputString, ".")
	// fmt.Println(addresses)
	if len(addresses) != 4 {
		return false
	}
	for _, value := range addresses {
		if value == "00" || value == "01" {
			return false
		}
		i, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if !(0 <= i && i <= 255) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%v", isIPv4Address("10.15.20.5"))
}
