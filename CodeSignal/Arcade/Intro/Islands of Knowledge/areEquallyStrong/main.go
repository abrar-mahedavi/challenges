package main

import "fmt"

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	/*
		if ((yourLeft+ yourRight) == (friendsLeft + friendsRight)) {
				return true
		}
		return false
	*/
	return (yourLeft == friendsRight && yourRight == friendsLeft) || (yourLeft == friendsLeft && yourRight == friendsRight)
}

func main() {
	fmt.Printf("%v", areEquallyStrong(10, 15, 20, 5))
}
