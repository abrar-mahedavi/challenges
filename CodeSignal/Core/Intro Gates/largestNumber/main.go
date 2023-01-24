package main

import "fmt"

func largestNumber(n int) (int){
    var sum = 9
    for index:=1 ;index<n ;index++ {
        sum = (sum*10) + 9
    }
    return sum
}


func main() {
	fmt.Printf("%v", largestNumber(2))
}
