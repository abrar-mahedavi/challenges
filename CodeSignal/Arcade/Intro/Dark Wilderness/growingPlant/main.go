package main

import "fmt"

func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
	var count int
	var height int
	for height < desiredHeight {
		height = height + upSpeed
		count = count + 1
		if height >= desiredHeight {
			break
		}
		height = height - downSpeed
	}
	return count
}

func main() {
	fmt.Print("😃")
	fmt.Printf("%v", growingPlant(100, 10, 910))
}
