package main

import (
	"fmt"
	"regexp"
)

func variableName(name string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z_$][a-zA-Z_$0-9]*$", name)
	// fmt.Println(match)
	return match
}

func main() {
	fmt.Printf("%v", variableName("var_1__Int"))
}
