package main

import (
	"fmt"
	"strconv"
)

func deleteDigit(n int) int {
	s := strconv.Itoa(n)
	max := 0
	for i := 0; i < len(s); i++ {
		ds := string(s[0:i]) + string(s[i+1:])
		// fmt.Println(ds)
		dn, _ := strconv.Atoi(ds)
		//fmt.Println(dn)
		if dn > max {
			max = dn
		}
	}

	return max
}

func main() {
	fmt.Printf("%v", deleteDigit(1001))
}
