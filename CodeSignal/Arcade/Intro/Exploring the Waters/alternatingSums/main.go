package main

import "fmt"

func alternatingSums(a []int) []int {
	sumOfTeams := make([]int, 0)
	var teamOneSum, teamTwoSum int
	for index, value := range a {
		if index%2 == 0 {
			teamOneSum = teamOneSum + value
		} else {
			teamTwoSum = teamTwoSum + value
		}
	}
	sumOfTeams = append(sumOfTeams, teamOneSum)
	sumOfTeams = append(sumOfTeams, teamTwoSum)
	return sumOfTeams
}

func main() {
	fmt.Printf("%v", alternatingSums([]int{1, 2, 3}))
}
