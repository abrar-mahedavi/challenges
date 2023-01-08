package main

import "strings"

func validTime(time string) bool {
    data:= strings.Split(time,":")
    hour, minutes := data[0] ,data[1]
    return (hour >="00" && hour <"24") && (minutes >="00" && minutes<"60")
}

func main() {
	fmt.Printf("%v", validTime("Ready, steady, go!"))
}
