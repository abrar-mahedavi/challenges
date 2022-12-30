package main

import "fmt"

func isPalindrome(st string) bool {
	if len(st) == 0 || len(st) == 1 {
		return true
	}
	if st[0] != st[len(st)-1] {
		return false
	}
	return isPalindrome(st[1 : len(st)-1])
}

func buildPalindrome(st string) string {
	tailst := ""
	for _, elem := range st {
		if isPalindrome(st + tailst) {
			return st + tailst
		}
		tailst = string(elem) + tailst
	}
	if isPalindrome(st + tailst) {
		return st + tailst
	}
	return ""
}

func main() {
	fmt.Printf("%v", buildPalindrome("bbbaacdafe"))
}
