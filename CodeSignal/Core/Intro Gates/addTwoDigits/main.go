package main

import "fmt"

func addTwoDigits(n int) (int) {
    var sum int
    for n > 0 {
        var rem = n % 10
        sum = sum + rem
        n = n / 10
    }
    return sum
}


func main() {
	fmt.Printf("%v", addTwoDigits(29))
}
