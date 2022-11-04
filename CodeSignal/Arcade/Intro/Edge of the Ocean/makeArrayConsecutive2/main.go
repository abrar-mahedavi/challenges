package main

import (
	"fmt"
	"sort"
)

func makeArrayConsecutive2(statues []int) int {
	sort.Ints(statues)
	var MaxDifference = statues[len(statues)-1] - statues[0]
	var difference = MaxDifference - (len(statues) - 1)
	//fmt.Println(MaxDifference , difference)
	return difference
	/*
	   fmt.Println(statues)
	   var count = 0
	   for index := 0 ; index <len(statues) -2;index++ {
	       var difference = statues[index] - statues[index + 1]
	       fmt.Println(difference)
	       if (difference != 1) {
	           count = count +  1
	       }
	       fmt.Println(count)
	   }
	   return count
	*/
}

func main() {
	fmt.Printf("%v", makeArrayConsecutive2([]int{6, 2, 3, 8}))
}
