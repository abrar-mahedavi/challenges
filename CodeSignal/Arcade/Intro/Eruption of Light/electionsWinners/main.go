package main

import "fmt"

func electionsWinners(votes []int, k int) int {
	return votes[0]
}

func main() {
	fmt.Printf("%v", electionsWinners([]int{2, 3, 5, 2}, 3))
}
