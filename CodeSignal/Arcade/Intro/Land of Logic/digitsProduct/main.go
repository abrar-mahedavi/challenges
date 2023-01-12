package main

import "fmt"
import "regexp"

func digitsProduct(product int) int {
    for i := 1; i < 255900; i++{
      if product == summation(i){
          return i
      }
    }
    return -1
}

func summation(number int) int{
    remainder:= 0
    sumResult:= 1
    for number != 0 {
        remainder = number % 10
        sumResult *= remainder
        number = number / 10
    }
    return sumResult
}


func main() {
	fmt.Printf("%v", digitsProduct(12))
}
