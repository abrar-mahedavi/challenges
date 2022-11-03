package main

//ShapeArea is a function to calculate the shape of the area
func ShapeArea(n int) int {
	if n == 1 {
		return 1
	} else {
		var size = 1
		for index := 1; index < n+1; index++ {
			size = size + index*4 - 4
		}
		return size
	}
}


func main() {
	fmt.Printf("%v", ShapeArea(3))
}

