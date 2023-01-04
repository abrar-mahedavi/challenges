package main

import (
	"fmt"
	"unicode"
)

func isDigit(symbol string) bool {
	return unicode.IsDigit(rune(symbol[0]))
}

func main() {
	fmt.Printf("%v", isDigit("0"))
}
