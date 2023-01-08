package main

import "fmt"
import "regexp"

func longestWord(text string) string {
    longest := ""
    re := regexp.MustCompile("[A-Za-z]+")
    for _, elem := range re.FindAllString(text, -1) {
        if len(elem) > len(longest) {
            longest = elem
        }
    }
    return longest
}


func main() {
	fmt.Printf("%v", longestWord("Ready, steady, go!"))
}
