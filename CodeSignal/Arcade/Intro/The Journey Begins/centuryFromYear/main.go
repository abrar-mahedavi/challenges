package main

import (
	"math"
)

func centuryFromYear(year int) int {
	return int(math.Floor(float64(year-1)/100) + 1)
}

/*
func main() {
	fmt.Printf("%v", centuryFromYear(2001))
}
*/
