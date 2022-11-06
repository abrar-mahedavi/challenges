package main

import (
	"sort"
)

func sortByHeight(a []int) []int {
	slc2 := make([]int, len(a))
	copy(slc2, a)
	sort.Ints(slc2)
	// fmt.Println(slc2)
	slc2 = RemoveAnElement(slc2, -1)
	// fmt.Println(slc2)
	var inc int
	for index, value := range a {
		if value != -1 {
			a[index] = slc2[inc]
			inc++
		}
	}
	return a
}

//RemoveAnElement by the key needs sorted output
func RemoveAnElement(s []int, element int) []int {
	var counter int
	for index, value := range s {
		if value != element {
			counter = index
			break
		}
	}
	return s[counter:]
}

/*
func main() {
	fmt.Printf("%v", sortByHeight([]int{-1, 150, 160, 170, -1, -1, 180, 190}))
}
*/
