package main

func checkPalindrome(inputString string) bool {
	var lastIndex = len(inputString) - 1
	for index := range inputString {
		if inputString[index] != inputString[lastIndex] {
			return false
		}
		lastIndex--
	}
	return true
}

/*
func main() {
	fmt.Printf("%v", checkPalindrome("abba"))
}
*/
