package main

import "fmt"

func differentSymbolsNaive(s string) int {
	charMap := make(map[rune]int)
	for _, value := range s {
		if 0 != charMap[value] {
			charMap[value] = 1
		} else {
			charMap[value] = 1
		}
	}
	return len(charMap)
}

func main() {
	fmt.Printf("%v", differentSymbolsNaive("adsda"))
}
