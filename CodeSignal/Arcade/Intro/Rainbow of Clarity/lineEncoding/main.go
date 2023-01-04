package main

import (
	"fmt"
	"strconv"
)

func lineEncoding(s string) (res string) {
	count := 1
	s += "#"
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			count++
		} else {
			if count > 1 {
				res = res + strconv.Itoa(count)
			}
			res = res + string(s[i-1])
			count = 1
		}
	}
	return
}

func lineEncoding1(s string) string {
	var mapOfCharacters = make(map[rune]int)
	// var temp rune
	var testString string = ""
	/*
		for key, value := range s {
			if mapOfCharacters[value] == 0 {
				mapOfCharacters[value] = mapOfCharacters[value] + 1
				// testString = testString + string(key) + "" + string(value)
				// temp = value
			} else {
				mapOfCharacters[value] = mapOfCharacters[value] + 1
			}
		}
	*/
	fmt.Println(mapOfCharacters)
	testString = ""
	for key, value := range mapOfCharacters {
		// fmt.Printf("%d %s", key, string(value))
		if value == 1 {
			testString = testString + string(key)
		} else {
			testString = testString + strconv.Itoa(value) + string(key)
		}

	}
	fmt.Println(testString)
	return "2a3b1c"
}

func main() {
	fmt.Printf("%v", lineEncoding("aabbbc"))
	fmt.Printf("%v", lineEncoding1("aabbbc"))
}
