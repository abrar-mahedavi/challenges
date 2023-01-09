package main

import "strings"
import "regexp"

func sumUpNumbers(inputString string) int {
    re := regexp.MustCompile("[0-9]+")
    res := re.FindAllString(inputString, -1)
    var sum = 0;
    for i := 0; i < len(res); i++ {
        intVar, _ := strconv.Atoi(res[i])
        sum += intVar

    }
    return sum;
}

func main() {
	fmt.Printf("%v", sumUpNumbers("2 apples, 12 oranges"))
}
