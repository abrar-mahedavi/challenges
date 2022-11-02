package main

import "fmt"

//Add is to add the 2 numbers
func Add(param1, param2 int) int {
	return param1 + param2
}

func main() {
	fmt.Printf("%v", Add(3, 3))
}
