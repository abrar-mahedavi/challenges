package main

import "fmt"

func commonCharacterCount(s1 string, s2 string) int {
	mapOfCharcterAndCountOfS1 := make(map[rune]int)
	for _, value := range s1 {
		if mapOfCharcterAndCountOfS1[value] != 0 {
			mapOfCharcterAndCountOfS1[value] = mapOfCharcterAndCountOfS1[value] + 1
		} else {
			mapOfCharcterAndCountOfS1[value] = 1
		}
	}
	mapOfCharcterAndCountOfS2 := make(map[rune]int)
	for _, value := range s2 {
		if mapOfCharcterAndCountOfS2[value] != 0 {
			mapOfCharcterAndCountOfS2[value] = mapOfCharcterAndCountOfS2[value] + 1
		} else {
			mapOfCharcterAndCountOfS2[value] = 1
		}
	}
	// fmt.Printf("%+v",mapOfCharcterAndCountOfS1)
	// fmt.Printf("%+v",mapOfCharcterAndCountOfS2)
	var commonElements int
	for key, value := range mapOfCharcterAndCountOfS1 {
		if value != 0 {
			if mapOfCharcterAndCountOfS1[key] <= mapOfCharcterAndCountOfS2[key] {
				commonElements = commonElements + mapOfCharcterAndCountOfS1[key]
			} else {
				commonElements = commonElements + mapOfCharcterAndCountOfS2[key]
			}
		}
	}
	return commonElements
}

func main() {
	fmt.Printf("%v ", commonCharacterCount("aabc", "aabc"))
}
